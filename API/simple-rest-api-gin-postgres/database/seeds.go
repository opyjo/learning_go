package database

import (
	"github.com/opyjo/simple-rest-api-gin-postgres/models"
	"gorm.io/gorm"
)

// SeedBooks seeds the database with initial book data
func SeedBooks(db *gorm.DB) error {
    books := []models.Book{
        {Title: "1984", Author: "George Orwell", Description: "A dystopian social science fiction novel and cautionary tale."},
        {Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Description: "A novel about the Roaring Twenties and the pursuit of the American Dream."},
        {Title: "To Kill a Mockingbird", Author: "Harper Lee", Description: "A novel centered on themes of racial injustice and moral growth."},
				{Title: "1984", Author: "George Orwell", Description: "A dystopian social science fiction novel and cautionary tale."},
				  {
   Title: "Pride and Prejudice",
    Author: "Jane Austen",
    Description: "A romantic comedy novel set in Regency England."},
  {
   Title: "The Lord of the Rings",
    Author: "J.R.R. Tolkien",
    Description: "A high-fantasy epic set in Middle-earth."},
  {
   Title: "The Hitchhiker's Guide to the Galaxy",
    Author: "Douglas Adams",
    Description: "A comedic science fiction series."},
  {
   Title: "Frankenstein",
    Author: "Mary Shelley",
    Description: "A gothic horror novel about a scientist who creates a monster."},
  {
   Title: "One Hundred Years of Solitude",
    Author: "Gabriel García Márquez",
    Description: "A magical realism novel chronicling the history of the Buendía family."},
  {
   Title: "The Count of Monte Cristo",
    Author: "Alexandre Dumas",
    Description: "An adventure novel about a man seeking revenge."},
  {
   Title: "Brave New World",
    Author: "Aldous Huxley",
    Description: "A dystopian novel exploring a future society based on scientific advancement."},
  {
   Title: "The Silmarillion",
    Author: "J.R.R. Tolkien",
    Description: "A collection of myths and legends that form the backstory of Middle-earth."},
  {
   Title: "Dune",
    Author: "Frank Herbert",
    Description: "A science fiction novel set in the distant future."},
  {
   Title: "Neuromancer",
    Author: "William Gibson",
    Description: "A cyberpunk novel considered one of the genre's founding works."},
	}


				
    

    // Check and create books if they do not already exist
    for _, book := range books {
        err := db.FirstOrCreate(&book, models.Book{Title: book.Title}).Error
        if err != nil {
            return err
        }
    }
    return nil
}