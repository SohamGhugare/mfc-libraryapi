package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context){
	// Creating a struct to bind the request body to
	var book struct {
		Title string
		Author string
		Type string
	}

	// Error handling
	if err:=c.Bind(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create book, check the request body",
		})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{
		"book": book,
	})
}