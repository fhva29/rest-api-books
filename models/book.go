package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Genre       string
	Year        int
	Description *string `gorm:"type:text"`

	AuthorID uint
	Author   Author
}
