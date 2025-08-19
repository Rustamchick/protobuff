package main

import (
	"grpc-pet/cmd/test"
	"grpc-pet/pkg/config"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// logrus.New()

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetLevel(logrus.DebugLevel)

	if err := godotenv.Load(); err != nil {
		logrus.Errorf("error loading env vars: %s", err)
	}

	cfg := config.InitConfig()

	test.TestAll(cfg)
}
