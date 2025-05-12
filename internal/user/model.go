package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password,omitempty"`
	PhoneNumber  string    `json:"phone_number,omitempty"`
	PhoneVerfied bool      `json:"phone_verified"`
	Address      string    `json:"address"`
	SocialID     string    `json:"social_id,omitempty"`
	Provider     string    `json:"provider"` // email, phone, social login
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
