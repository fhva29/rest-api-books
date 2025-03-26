package config

import (
	"fmt"
	"log"
	"os"
	"rest_api_books/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	db_folder := "database/"
	err := os.MkdirAll(db_folder, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	database, err := gorm.Open(sqlite.Open(db_folder+"books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	err = database.AutoMigrate(&models.Author{}, &models.Book{})
	if err != nil {
		log.Fatal(err)
	}

	db = database
}

func GetDatabase() *gorm.DB {
	return db
}
