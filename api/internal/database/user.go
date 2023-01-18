package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Handle       string         `gorm:"uniqueIndex" json:"handle"`
	Email        string         `gorm:"uniqueIndex" json:"email"`
	PasswordHash *string        `json:"-"`
	AvatarURL    *string        `json:"avatar_url"`
	IsVerified   bool           `json:"-"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
