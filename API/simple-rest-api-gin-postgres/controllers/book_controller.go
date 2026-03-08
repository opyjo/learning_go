package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opyjo/simple-rest-api-gin-postgres/models"
	"github.com/opyjo/simple-rest-api-gin-postgres/services"
)

func GetBooks(c *gin.Context) {
    var books []models.Book
    err := services.GetAllBooks(&books)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    err := services.GetBookByID(&book, id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := services.CreateBook(&book)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := services.UpdateBook(&book, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    err := services.DeleteBook(&book, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}