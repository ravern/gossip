package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentLike struct {
	gorm.Model
	ID        uuid.UUID
	UserID    uuid.UUID
	CommentID uuid.UUID
}
