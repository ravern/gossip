package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Title     string         `gorm:"uniqueIndex;not null" json:"title"`
	Body      string         `gorm:"not null" json:"body"`
	Tags      pq.StringArray `gorm:"type:text[];not null" json:"tags"`
	AuthorID  uuid.UUID      `gorm:"not null" json:"-"`
	Author    *User          `json:"author,omitempty"`
	Comments  []Comment      `json:"comments"`
	PostLikes []PostLike     `json:"likes"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
