package postgres

import (
	"database/sql"
	"fmt"

	"github.com/mirjalilova/auth-service/pkg/storage"

	"golang.org/x/exp/slog"

	_ "github.com/lib/pq"
	"github.com/mirjalilova/auth-service/config"
)

type Storage struct {
	Db      *sql.DB
	OrderDb *sql.DB
	AuthS   storage.AuthI
	UserS   storage.UserI
}

func NewPostgresStorage(config *config.Config) (*Storage, error) {
	conn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", 
		config.DB_USER, 
		config.DB_PASSWORD, 
		config.DB_HOST, 
		config.DB_PORT, 
		config.DB_NAME,
	)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	slog.Info("connected to db")

	return &Storage{
		Db:      db,
		AuthS:   NewAuthRepo(db),
		UserS:   NewUserRepo(db),
	}, nil
}
