package app

import (
	grpcapp "grpc-pet/pkg/app/grpc"
	"time"

	"github.com/sirupsen/logrus"
)

type App struct {
	GRPCApp *grpcapp.App
}

func New(log *logrus.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: init storage

	// TODO: init auth service

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCApp: grpcApp,
	}
}
