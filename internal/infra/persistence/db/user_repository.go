package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/elaurentium/vaultui/internal/domain/entities"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.Users) error
	GetByEmail(ctx context.Context, email string) (*entities.Users, error)
	GetById(ctx context.Context, id uuid.UUID) (*entities.Users, error)
	Update(ctx context.Context, user *entities.Users) error
	CheckEmailExist(ctx context.Context, email string) (bool, error) 
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *entities.Users) error {
	query := `
		INSERT INTO SYS_USR (SYS_ID, SYS_NAME, SYS_EMAIL, SYS_HPASS, SYS_SALT, SYS_BDAY, SYS_ROLE, SYS_ACTIVE, SYS_VEMAIL, SYS_CREATE_AT, SYS_UPDATE_AT)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Name, user.Email, user.HashedPassword, user.Salt, user.Birthday,
		user.Role, user.IsActive, user.EmailVerified, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *userRepository) Update(ctx context.Context, user *entities.Users) error {
	query := `
		UPDATE SYS_USR SET SYS_NAME = $2, SYS_EMAIL = $3, SYS_HPASS = $4, SYS_SALT = $5, SYS_BDAY = $6, SYS_ACTIVE = $7,
			SYS_VEMAIL = $8, SYS_ROLE = $9, SYS_CREATE_AT = $10, SYS_UPDATE_AT = $11
		WHERE SYS_ID = $1 AND SYS_DELETED_AT IS NULL`
	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Name, user.Email, user.HashedPassword, user.Salt,
		user.Role, user.IsActive, user.EmailVerified, user.UpdatedAt)
	return err
}


func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entities.Users, error) {
	user := &entities.Users{}

	query := `
		SELECT * FROM SYS_USR WHERE SYS_EMAIL = $1 AND SYS_DELETED_AT IS NULL 
	`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.HashedPassword, &user.Salt, &user.Role, &user.IsActive, &user.EmailVerified, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) CheckEmailExist(ctx context.Context, email string) (bool, error) {
	var exist bool

	err := r.db.QueryRowContext(ctx, "SELECT EXIST(SELECT 1 FROM SYS_USR WHERE SYS_EMAIL = $1 AND SYS_DELETED_AT IS NULL)", email).Scan(&exist)

	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}

	return exist, nil
}

func (r *userRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.Users, error) {
	user := &entities.Users{}

	query := `
		SELECT * FROM SYS_USR WHERE SYS_ID = $1
	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.HashedPassword, &user.Salt, &user.Role, &user.IsActive, &user.EmailVerified, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}