package repositories

import (
	"context"

	"github.com/elaurentium/vaultui/internal/domain/entities"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.Users) error
	GetById(ctx context.Context, id uuid.UUID) (*entities.Users, error)
	GetByEmail(ctx context.Context, email string) (*entities.Users, error)
	Update(ctx context.Context, user *entities.Users) error
	CheckEmailExist(ctx context.Context, email string) (bool, error) 
}