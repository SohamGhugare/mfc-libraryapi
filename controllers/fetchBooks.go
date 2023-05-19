package controllers

import (
	"libraryapi/database"
	"libraryapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route to get all the books stored
func GetAllBooks(c *gin.Context){
	var books []models.Book
	if err:=database.FetchAllBooks(&books); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get books, check server logs",
		})
		return 
	}
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}