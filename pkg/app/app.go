package app

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/mirjalilova/auth-service/api"
	"github.com/mirjalilova/auth-service/api/handlers"
	"github.com/mirjalilova/auth-service/config"
	kafka "github.com/mirjalilova/auth-service/pkg/kafka/consumer"
	prd "github.com/mirjalilova/auth-service/pkg/kafka/producer"
	"golang.org/x/exp/slog"

	// "github.com/mirjalilova/auth-service/pkg/logger"
	"github.com/mirjalilova/auth-service/pkg/storage/postgres"
	"github.com/mirjalilova/auth-service/service"
)

func Run(cfg *config.Config) {
	// Define the base path where the logs are stored
	// basepath := "/home/feruza/GIT/mirjalilova/MedMon/auth-service"
	// Initialize the logger
	// l := logger.NewLogger(basepath, cfg.LOG_PATH)

	// Postgres Connection
	db, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		slog.Error("can't connect to db: %v", err)
		return
	}
	defer db.Db.Close()
	slog.Info("Connected to Postgres")

	// Redis Connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		slog.Error("Failed to connect to Redis: %v", err)
		return
	}
	slog.Info("Connected to Redis")

	// Initialize services
	authService := service.NewAuthService(db)
	userService := service.NewUserService(db)

	// Kafka setup
	brokers := []string{"kafka:9092"}
	cm := kafka.NewKafkaConsumerManager()
	pr, err := prd.NewKafkaProducer(brokers)
	if err != nil {
		slog.Error("Failed to create Kafka producer:", err)
		return
	}

	// Kafka message reader
	Reader(brokers, cm, authService, userService)

	// HTTP Server setup
	h := handlers.NewHandler(authService, userService, rdb, &pr)
	router := api.Engine(h)
	router.SetTrustedProxies(nil)

	if err := router.Run(cfg.AUTH_PORT); err != nil {
		slog.Error("can't start server: %v", err)
	}

	slog.Info("REST server started on port %s", cfg.AUTH_PORT)
}