package users

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Register(ctx context.Context, data NewUser) (User, error)
	Login(ctx context.Context, data LoginUserParams) (User, error)
}

type svc struct {
	repository userRepository
}

type userRepository interface {
	Register(context.Context, NewUser) (User, error)
	GetByEmail(context.Context, string) (User, error)
}

func NewService(repo userRepository) Service {
	return &svc{
		repository: repo,
	}
}

// Methods
func (s *svc) Register(ctx context.Context, data NewUser) (User, error) {
	user, err := s.repository.Register(ctx, data)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return User{}, ErrUserAlreadyExists
		}

		return User{}, err
	}

	return user, nil
}

func (s *svc) Login(ctx context.Context, data LoginUserParams) (User, error) {
	if len(data.Password) > 72 {
		return User{}, ErrPasswordTooLong
	}
	user, err := s.repository.GetByEmail(ctx, data.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, ErrInvalidCredentials
		}

		return User{}, err
	}
	
	return user, nil
}
