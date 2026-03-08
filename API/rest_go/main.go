package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// main is the entry point of the program.
// It sets up a basic HTTP server that listens on port 8080.
// The "/" route is handled by the homePage function.
// The server runs indefinitely until an error occurs.
func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/book", getBook)
	http.HandleFunc("/book/create", createBook)
	http.HandleFunc("/book/update", updateBook)
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// homePage is a handler function for the home page of the Simple REST API in Go.
// It writes a welcome message to the http.ResponseWriter.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Simple REST API in Golang!")
}


type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books []Book

// init is a special Go function that runs before the main function.
// It initializes the books slice with some example data.
func init() {
    books = []Book{
        {ID: "1", Title: "1984", Author: "George Orwell"},
        {ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"},
        {ID: "3", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
    }
}

// getBooks is a handler function that returns a list of books in JSON format.
// It writes the list of books to the http.ResponseWriter.
func getBooks(w http.ResponseWriter, _ *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

// getBook is a handler function that returns a single book by its ID in JSON format.
// It writes the book to the http.ResponseWriter.
func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    id := r.URL.Query().Get("id")
		fmt.Println(id)
    for _, book := range books {
        if book.ID == id {
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}


func createBook(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    var newBook Book
    err := json.NewDecoder(r.Body).Decode(&newBook)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    books = append(books, newBook)
    json.NewEncoder(w).Encode(newBook)
}


func updateBook(w http.ResponseWriter, r *http.Request) {
    if r.Method != "PUT" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    id := r.URL.Query().Get("id")
    var updatedBook Book
    err := json.NewDecoder(r.Body).Decode(&updatedBook)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    for i, book := range books {
        if book.ID == id {
            books[i] = updatedBook
            json.NewEncoder(w).Encode(updatedBook)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}