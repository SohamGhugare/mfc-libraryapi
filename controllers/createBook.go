package controllers

import (
	"fmt"
	"libraryapi/database"
	"libraryapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context){
	// Creating a struct to bind the request body to
	var body struct {
		Title string
		Author string
		Type string
	}

	// Error handling
	if err:=c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create book, check the request body",
		})
		return 
	}

	book := models.Book{
		Title: body.Title,
		Author: body.Author,
		Type: body.Type,
	}

	if err:=database.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to create book: %v", err),
		})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{
		"book": book,
	})
}