package handlers

import (
	"github.com/go-redis/redis/v8"
	kafka "github.com/mirjalilova/auth-service/pkg/kafka/producer"
	// "github.com/mirjalilova/auth-service/pkg/logger"
	"github.com/mirjalilova/auth-service/service"
)

type Handlers struct {
	Auth     *service.AuthService
	User     *service.UserService
	RDB      *redis.Client
	Producer kafka.KafkaProducer
	// Logger   *logger.Logger
}

func NewHandler(auth *service.AuthService, user *service.UserService, rdb *redis.Client, pr *kafka.KafkaProducer, /* l *logger.Logger */) *Handlers {
	return &Handlers{
		Auth:     auth,
		User:     user,
		RDB:      rdb,
		Producer: *pr,
		// Logger:   l,
	}
}