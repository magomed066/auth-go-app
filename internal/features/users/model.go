package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  []byte `gorm:"not null" json:"-"`
	Phone     string `json:"phone"`
}

type CreateUserParams struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Phone     string `json:"phone"`
}

type NewUser struct {
	CreateUserParams
	PasswordHash []byte
}

type LoginUserParams struct {
	Email 		string	`json:"email" validate:"required,email"`
	Password 	string	`json:"password" validate:"required"`
}
