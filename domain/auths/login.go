package domain

import (
	"time"

	"github.com/google/uuid"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	ContactIfo string    `json:"contact_info"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ID         uuid.UUID `json:"id"`
}

type LoginResp struct {
	User
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	IssuedAt  time.Time `json:"issued_at"`
}
