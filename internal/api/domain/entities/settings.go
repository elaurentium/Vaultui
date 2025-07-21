package entities

import "github.com/google/uuid"

type Settings struct {
	ID      uuid.UUID `json:"id"`
	Email   string    `json:"email"`
	Domains []string  `json:"domains"`
}
