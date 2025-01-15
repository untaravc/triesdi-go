package auth_repository

import "github.com/google/uuid"

type Auth struct {
	ID 	 uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email string
	Password string
}