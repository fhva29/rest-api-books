package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name           string    `json:"name"`
	Nacionality    string    `json:"nacionality"`
	DataNascimento time.Time `json:"data_nascimento"`
	Books          []Book    `gorm:"foreignkey:AuthorID" json:"books,omitempty"`
}
