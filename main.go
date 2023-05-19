package main

import (
	"libraryapi/controllers"
	"libraryapi/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.ConnectDatabase("data/books.db")
	initializers.SyncDatabase()
}

func main(){
	// Initializing the server
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong!",
		})
	})

	r.GET("/book/all", controllers.GetAllBooks)

	r.POST("/create", controllers.CreateBook)

	r.Run()
}