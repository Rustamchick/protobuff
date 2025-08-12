package main

import (
	"fmt"
	"grpc-pet/pkg/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error loading env vars: ", err)
	}

	cfg := config.MustInitConfig()

	fmt.Printf("config: %+v", cfg)
}
