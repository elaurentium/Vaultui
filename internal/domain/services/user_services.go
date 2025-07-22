package services

import (
	"context"
	"errors"
	"time"

	"github.com/elaurentium/vaultui/internal/domain/entities"
	"github.com/elaurentium/vaultui/internal/domain/repositories"
	"github.com/elaurentium/vaultui/internal/infra/api/auth"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo repositories.UserRepository
	auth     auth.AuthService
}

func NewUserService(userRepo repositories.UserRepository, auth auth.AuthService) *UserService {
	return &UserService{
		userRepo: userRepo,
		auth:     auth,
	}
}

func (s *UserService) Register(ctx context.Context, name, email, password, birthday string) (*entities.Users, error) {
	email_exist, err := s.userRepo.CheckEmailExist(ctx, email)

	if err != nil {
		return nil, err
	}

	if email_exist {
		return nil, errors.New("email already exist")
	}

	hashedPassword, salt, err := s.auth.HashPassword(password)

	if err != nil {
		return nil, err
	}

	user := &entities.Users{
		ID: uuid.New(),
		Name: name,
		Email: email,
		Birthday: birthday,
		HashPassword: hashedPassword,
		Salt: salt,
		EmailVerified: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}
