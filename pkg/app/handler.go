package app

import (
	"golang.org/x/exp/slog"

	kafka "github.com/mirjalilova/auth-service/pkg/kafka/consumer"
	// "github.com/mirjalilova/auth-service/pkg/logger"
	"github.com/mirjalilova/auth-service/service"
)

func Reader(brokers []string, kcm *kafka.KafkaConsumerManager, authService *service.AuthService, userService *service.UserService) {

	if err := kcm.RegisterConsumer(brokers, "reg-user", "auth", kafka.UserRegisterHandler(authService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			slog.Warn("Consumer for topic 'reg-user' already exists")
		} else {
			slog.Error("Error registering consumer: %v", "err", err)
		}
	}

	if err := kcm.RegisterConsumer(brokers, "upd-user", "auth", kafka.UserEditProfileHandler(userService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			slog.Warn("Consumer for topic 'upd-user' already exists")
		} else {
			slog.Error("Error registering consumer: %v", "err", err)
		}
	}

	if err := kcm.RegisterConsumer(brokers, "upd-pass", "auth", kafka.UserEditPasswordHandler(userService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			slog.Warn("Consumer for topic 'upd-pass' already exists")
		} else {
			slog.Error("Error registering consumer: %v", "err", err)
		}
	}

	if err := kcm.RegisterConsumer(brokers, "upd-setting", "auth", kafka.UserEditSettingHandler(userService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			slog.Warn("Consumer for topic 'upd-setting' already exists")
		} else {
			slog.Error("Error registering consumer: %v", "err", err)
		}
	}
}
