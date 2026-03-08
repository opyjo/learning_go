package main

import (
	"github.com/gin-gonic/gin"
)

type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {ID: "1", Title: "1984", Author: "George Orwell"},
    {ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"},
    {ID: "3", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
}

func main() {
    router := gin.Default()
    router.GET("/", homePage)
    router.GET("/books", getBooks)
    router.GET("/books/:id", getBookByID)
    router.POST("/books", createBook)
    router.PUT("/books/:id", updateBook)
    router.DELETE("/books/:id", deleteBook)
    router.Run(":8080")
}

func homePage(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Welcome to the Gin REST API!",
    })
}

func getBooks(c *gin.Context) {
    c.IndentedJSON(200, books)
}

func getBookByID(c *gin.Context) {
    id := c.Param("id")

    for _, book := range books {
        if book.ID == id {
            c.IndentedJSON(200, book)
            return
        }
    }
    c.JSON(404, gin.H{"message": "Book not found"})
}

func createBook(c *gin.Context) {
    var newBook Book

    if err := c.BindJSON(&newBook); err != nil {
        return
    }

    books = append(books, newBook)
    c.IndentedJSON(201, newBook)
}

func updateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBook Book

    if err := c.BindJSON(&updatedBook); err != nil {
        return
    }

    for i, book := range books {
        if book.ID == id {
            books[i] = updatedBook
            c.IndentedJSON(200, updatedBook)
            return
        }
    }
    c.JSON(404, gin.H{"message": "Book not found"})
}

func deleteBook(c *gin.Context) {
    id := c.Param("id")

    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            c.JSON(200, gin.H{"message": "Book deleted"})
            return
        }
    }
    c.JSON(404, gin.H{"message": "Book not found"})
}