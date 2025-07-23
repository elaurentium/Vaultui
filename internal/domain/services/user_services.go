package services

import (
	"context"
	"errors"
	"time"

	"github.com/elaurentium/vaultui/internal/domain/entities"
	"github.com/elaurentium/vaultui/internal/domain/repositories"
	"github.com/elaurentium/vaultui/internal/infra/auth"
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
		ID:             uuid.New(),
		Name:           name,
		Email:          email,
		Birthday:       birthday,
		HashedPassword: hashedPassword,
		Salt:           salt,
		Role:           "user",
		IsActive:       true,
		EmailVerified:  true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = s.userRepo.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if !s.auth.VerifyPassword(password, user.HashedPassword, user.Salt) {
		return "", "", errors.New("invalid credentials")
	}

	if !user.IsActive {
		return "", "", errors.New("account is deactivated")
	}

	user.UpdatedAt = time.Now()

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return "", "", err
	}

	accessToken, err := s.auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.auth.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) Update(ctx context.Context, id uuid.UUID, name string, email string) (*entities.Users, error) {
	user, err := s.userRepo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Email = email
	user.UpdatedAt = time.Now()

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ChangePassword(ctx context.Context, id uuid.UUID, currentPassword, newPassword string) error {
	user, err := s.userRepo.GetById(ctx, id)

	if err != nil {
		return err
	}

	if !s.auth.VerifyPassword(currentPassword, user.HashedPassword, user.Salt) {
		return errors.New("current password is incorrect")
	}

	hashedPassword, salt, err := s.auth.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.HashedPassword = hashedPassword
	user.Salt = salt
	user.UpdatedAt = time.Now()

	return s.userRepo.Update(ctx, user)
}
