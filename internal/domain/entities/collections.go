package entities

import "github.com/google/uuid"

type Collection struct {
	ID             uuid.UUID `json:"id"`
	OrganizationId uuid.UUID `json:"organization_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description,omitempty"`
	Manage         bool      `json:"manage,omitempty"`
}