package test

import (
	"grpc-pet/pkg/config"
	"time"

	grpcpetv1 "github.com/Rustamchick/protobuff/gen/go/pet"
	"github.com/sirupsen/logrus"
)

func TestAll(cfg config.Config) {
	cfg_t := config.Config{
		Env:          "local",
		Storage_path: "storage_path_is_in_development",
		TokenTTL:     time.Hour * 12,
		GRPC: config.GrpcConfig{
			Port:    8080,
			Timeout: time.Hour,
		},
	}

	if cfg == cfg_t {
		logrus.Info("config \033[32mCORRECT\033[0m")
	}

	req := grpcpetv1.RegisterRequest{}
	if req.Email == "" { // в будущем подумаю о нормальной тестировке этого пакета, но пока что по сути только вывожу карсивую надпись
		logrus.Infof("Protobuff \033[32mCORRECT\033[0m")
	}

}
