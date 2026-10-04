package users

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (repo *Repository) Register(ctx context.Context, data NewUser) (User, error) {
	user := User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.PasswordHash,
		Phone:     data.Phone,
	}

	err := repo.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (repo *Repository) Login(ctx context.Context, data LoginUserParams) (User, error) {
	var user User

	err := repo.db.WithContext(ctx).Where("email = ?", data.Email).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}
