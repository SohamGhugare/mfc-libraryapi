package models

import (
	"errors"

	"gorm.io/gorm"
)

// Book model
type Book struct {
	gorm.Model
	Title string `gorm:"not null"`
	Author string `gorm:"not null"`
	Type string `gorm:"not null"`
}

// Checking if name and email are passed as "" which may lead to non-null values
func (b *Book) BeforeSave(tx *gorm.DB) (err error){
	if b.Title == "" {
		return errors.New("title cannot be empty")
	}
	if b.Author == "" {
		return errors.New("author cannot be empty")
	}
	if b.Type == "" {
		return errors.New("type cannot be empty")
	}
	return nil
}