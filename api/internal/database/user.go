package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Handle       string    `gorm:"uniqueIndex"`
	Email        string    `gorm:"uniqueIndex"`
	PasswordHash string
	AvatarURL    string
	IsVerified   bool
}
