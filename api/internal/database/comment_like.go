package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentLike struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID      `gorm:"uniqueIndex:idx_comment_likes_user_id_comment_id" json:"user_id"`
	User      *User          `json:"-"`
	CommentID uuid.UUID      `gorm:"uniqueIndex:idx_comment_likes_user_id_comment_id" json:"-"`
	Comment   *Comment       `json:"-"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
