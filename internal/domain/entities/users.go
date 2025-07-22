package entities

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Birthday      string    `json:"birthday"`
	HashPassword  string    `json:"-"`
	Salt          string    `json:"-"`
	EmailVerified bool      `json:"email_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at,omitempty"`
}
