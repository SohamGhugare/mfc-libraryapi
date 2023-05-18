package initializers

import (
	"libraryapi/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Conecting to database
func ConnectDatabase(uri string){
	log.Println("Connecting to database...")
	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
	if err!=nil{
		log.Fatal("Failed to connect to database...")
	}
	DB = db
	log.Println("Database connection established!")
}

// Syncing all the schemas
func SyncDatabase(){
	log.Println("Syncing database...")
	err := DB.AutoMigrate(&models.Book{})
	if err!=nil{
		log.Fatal("Failed to sync database...")
	}
	log.Println("Successfully synced all models!")
}