package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opyjo/simple-rest-api-gin-postgres/controllers"
	"github.com/opyjo/simple-rest-api-gin-postgres/database"
)

func main() {
    router := gin.Default()

    database.Connect()

    router.GET("/books", controllers.GetBooks)
    router.GET("/books/:id", controllers.GetBook)
    router.POST("/books", controllers.CreateBook)
    router.PUT("/books/:id", controllers.UpdateBook)
    router.DELETE("/books/:id", controllers.DeleteBook)

    router.Run(":8080")
}