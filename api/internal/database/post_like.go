package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostLike struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID      `gorm:"uniqueIndex:idx_post_likes_user_id_post_id" json:"user_id"`
	User      *User          `json:"-"`
	PostID    uuid.UUID      `gorm:"uniqueIndex:idx_post_likes_user_id_post_id" json:"-"`
	Post      *Post          `json:"-"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
