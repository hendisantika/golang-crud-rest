package db

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ContactID   uuid.UUID `json:"contact_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Street      string    `json:"street"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
