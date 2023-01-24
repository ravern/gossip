package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Body         string         `gorm:"not null" json:"body"`
	AuthorID     uuid.UUID      `gorm:"not null" json:"-"`
	Author       *User          `json:"author,omitempty"`
	PostID       uuid.UUID      `gorm:"not null" json:"-"`
	Post         *Post          `json:"-"`
	CommentLikes []CommentLike  `json:"likes"`
	CreatedAt    time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
