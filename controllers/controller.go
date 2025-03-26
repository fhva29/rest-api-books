package controllers

import (
	"rest_api_books/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeControllers() {
	db = config.GetDatabase()
}
