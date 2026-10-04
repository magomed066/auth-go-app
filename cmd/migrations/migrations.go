package migrations

import (
	"charm.land/log/v2"
	"github.com/magomed066/auth-go-app/internal/features/users"
	"gorm.io/gorm"
)

func MigreateDB(db *gorm.DB) {
	err := db.AutoMigrate(&users.User{})

	if err != nil {
		log.Fatal("Could not migrate DB", "err", err)
	}
}
