package services

import (
	"github.com/opyjo/simple-rest-api-gin-postgres/database"
	"github.com/opyjo/simple-rest-api-gin-postgres/models"
)

func GetAllBooks(books *[]models.Book) error {
    if err := database.DB.Find(books).Error; err != nil {
        return err
    }
    return nil
}

func GetBookByID(book *models.Book, id string) error {
    if err := database.DB.Where("id = ?", id).First(book).Error; err != nil {
        return err
    }
    return nil
}

func CreateBook(book *models.Book) error {
    if err := database.DB.Create(book).Error; err != nil {
        return err
    }
    return nil
}

func UpdateBook(book *models.Book, id string) error {
    if err := database.DB.Model(book).Where("id = ?", id).Updates(book).Error; err != nil {
        return err
    }
    return nil
}

func DeleteBook(book *models.Book, id string) error {
    if err := database.DB.Where("id = ?", id).Delete(book).Error; err != nil {
        return err
    }
    return nil
}