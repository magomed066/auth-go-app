package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error ) {
	// dsn := "postgres://postgres:root@localhost:5432/price_list"
	// dsn := "host=localhost user=root password=root dbname=price_list port=5432 sslmode=disable"

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
}
