package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/magomed066/auth-go-app/internal/env"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
		os.Exit(1)
	}

	cfg := config{
		addr: env.GetPort(),
	}

	api := application{
		config:	cfg,
	}

	//Error log
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	
	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}