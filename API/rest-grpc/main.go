package main

import (
	"fmt"
	"log"
	"net"

	"github.com/opyjo/grpc-bookstore/database"
	"github.com/opyjo/grpc-bookstore/models"
	"github.com/opyjo/grpc-bookstore/proto"
	"github.com/opyjo/grpc-bookstore/services"
	"google.golang.org/grpc"
)

func main() {
    database.Connect()
    database.DB.AutoMigrate(&models.Book{})

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterBookServiceServer(grpcServer, &services.BookServiceServer{})

    fmt.Println("Server is running on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}