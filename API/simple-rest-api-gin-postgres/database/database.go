package database

import (
	"fmt"
	"log"

	"github.com/opyjo/simple-rest-api-gin-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    var err error
    dsn := "host=localhost user=postgres password= dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    // Connect to the database
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    fmt.Println("Database connection established")

    // Auto-migrate the Book model
    DB.AutoMigrate(&models.Book{})

     // Seed the database
    if err := SeedBooks(DB); err != nil {
        log.Fatal("Failed to seed books:", err)
    }
}