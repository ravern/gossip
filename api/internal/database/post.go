package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Title     string         `gorm:"uniqueIndex" json:"title"`
	Body      string         `json:"body"`
	Tags      pq.StringArray `gorm:"type:text[]" json:"tags"`
	AuthorID  uuid.UUID      `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
