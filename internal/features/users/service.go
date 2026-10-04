package users

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Register(ctx context.Context, data NewUser) (User, error)
}

type svc struct {
	repository Repository
}

func NewService(repo *Repository) Service {
	return &svc{
		repository: *repo,
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
