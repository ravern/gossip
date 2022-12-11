package database

import "gorm.io/gorm"

func Configure(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
