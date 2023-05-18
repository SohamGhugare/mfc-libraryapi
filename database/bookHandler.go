package database

import (
	"libraryapi/initializers"
	"libraryapi/models"
)

// Creating a new book in database
func CreateBook(book *models.Book) error {
	tx := initializers.DB.Create(book)
	return tx.Error
}