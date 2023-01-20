package database

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const NormalRole UserRole = "user"
const ModeratorRole UserRole = "moderator"
const AdminRole UserRole = "admin"

func (r *UserRole) Scan(value interface{}) error {
	*r = UserRole(value.(string))
	return nil
}

func (r UserRole) Value() (driver.Value, error) {
	return string(r), nil
}

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Handle       string         `gorm:"uniqueIndex" json:"handle"`
	Email        string         `gorm:"uniqueIndex" json:"email"`
	PasswordHash *string        `json:"-"`
	AvatarURL    *string        `json:"avatar_url"`
	IsVerified   bool           `json:"-"`
	Role         UserRole       `gorm:"default:user" json:"-"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
