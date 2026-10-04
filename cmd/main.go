package main

import (
	"charm.land/log/v2"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	db "github.com/magomed066/auth-go-app"
	"github.com/magomed066/auth-go-app/cmd/migrations"
	"github.com/magomed066/auth-go-app/internal/env"
)

func initConfig() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file", "err", err)
	}

	db, err := db.Connect(env.GetEnvs().DB_URL)
	if err != nil {
		log.Fatal("Could not connect DB", "err", err)
	}

	migrations.MigreateDB(db)

	log.Info("DB is migrated")
	log.Info("DB is connected")

	return db
}

func main() {
	db := initConfig()

	cfg := config{
		addr: env.GetEnvs().PORT,
		db: db,
	}

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		log.Fatal("Server failed to start", "err", err)
	}
}
