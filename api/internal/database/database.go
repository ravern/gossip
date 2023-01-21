package database

import "gorm.io/gorm"

func Configure(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&PostLike{})
	db.AutoMigrate(&CommentLike{})
}
