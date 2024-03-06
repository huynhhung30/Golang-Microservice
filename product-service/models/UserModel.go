package models

import (
	"time"
)

type UserModel struct {
	Id            int       `json:"id"`
	UserType      string    `json:"user_type"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Avatar        string    `json:"avatar"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Status        string    `json:"status"`
	EmailVerified bool      `json:"email_verified"`
	Address       string    `json:"address"`
	PhoneNumber   string    `json:"phone_number"`
	LoginMethod   string    `json:"login_method"`
	SocialId      string    `json:"social_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
