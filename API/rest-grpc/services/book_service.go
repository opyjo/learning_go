package services

import (
	"context"
	"errors"

	"github.com/opyjo/grpc-bookstore/database"
	"github.com/opyjo/grpc-bookstore/models"
	"github.com/opyjo/grpc-bookstore/proto"
	"gorm.io/gorm"
)

type BookServiceServer struct {
    proto.UnimplementedBookServiceServer
}

func (s *BookServiceServer) CreateBook(ctx context.Context, req *proto.Book) (*proto.Book, error) {
    book := models.Book{
        Title:       req.Title,
        Author:      req.Author,
        Description: req.Description,
    }
    if err := database.DB.Create(&book).Error; err != nil {
        return nil, err
    }
    res := &proto.Book{
        Id:          int32(book.ID),
        Title:       book.Title,
        Author:      book.Author,
        Description: book.Description,
    }
    return res, nil
}

func (s *BookServiceServer) GetBook(ctx context.Context, req *proto.BookID) (*proto.Book, error) {
    var book models.Book
    if err := database.DB.First(&book, req.Id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("book not found")
        }
        return nil, err
    }
    res := &proto.Book{
        Id:          int32(book.ID),
        Title:       book.Title,
        Author:      book.Author,
        Description: book.Description,
    }
    return res, nil
}

func (s *BookServiceServer) UpdateBook(ctx context.Context, req *proto.Book) (*proto.Book, error) {
    var book models.Book
    if err := database.DB.First(&book, req.Id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("book not found")
        }
        return nil, err
    }
    book.Title = req.Title
    book.Author = req.Author
    book.Description = req.Description
    if err := database.DB.Save(&book).Error; err != nil {
        return nil, err
    }
    res := &proto.Book{
        Id:          int32(book.ID),
        Title:       book.Title,
        Author:      book.Author,
        Description: book.Description,
    }
    return res, nil
}

func (s *BookServiceServer) DeleteBook(ctx context.Context, req *proto.BookID) (*proto.Empty, error) {
    if err := database.DB.Delete(&models.Book{}, req.Id).Error; err != nil {
        return nil, err
    }
    return &proto.Empty{}, nil
}

func (s *BookServiceServer) ListBooks(ctx context.Context, req *proto.Empty) (*proto.BookList, error) {
    var books []models.Book
    if err := database.DB.Find(&books).Error; err != nil {
        return nil, err
    }
    var res proto.BookList
    for _, book := range books {
        res.Books = append(res.Books, &proto.Book{
            Id:          int32(book.ID),
            Title:       book.Title,
            Author:      book.Author,
            Description: book.Description,
        })
    }
    return &res, nil
}