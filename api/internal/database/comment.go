package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID       uuid.UUID
	Body     string
	AuthorID uuid.UUID
	PostID   uuid.UUID
}
