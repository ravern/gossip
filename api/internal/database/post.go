package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID       uuid.UUID
	Title    string
	Body     string
	Tags     []string
	AuthorID uuid.UUID
}
