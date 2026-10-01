package main

import (
	"charm.land/log/v2"

	"github.com/joho/godotenv"
	db "github.com/magomed066/auth-go-app"
	"github.com/magomed066/auth-go-app/internal/env"
)

func initConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file", "err", err)
	}

	db.Connect(env.GetEnvs().DB_URL)
	log.Info("DB is connected")
}

func main() {
	initConfig()
	
	cfg := config{
		addr: env.GetEnvs().PORT,
	}

	api := application{
		config:	cfg,
	}
	
	if err := api.run(api.mount()); err != nil {
		log.Fatal("Server failed to start", "err", err)
	}
}