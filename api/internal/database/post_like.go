package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostLike struct {
	gorm.Model
	ID        uuid.UUID
	UserID    uuid.UUID
	CommentID uuid.UUID
}
