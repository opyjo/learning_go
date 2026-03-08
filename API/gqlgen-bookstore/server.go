package main

import (
	"log"
	"net/http"

	"gqlgen-bookstore/database"
	"gqlgen-bookstore/graph"
	"gqlgen-bookstore/graph/generated"
	"gqlgen-bookstore/models"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
    database.Connect()
    database.DB.AutoMigrate(&models.Book{})

    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

    http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
    http.Handle("/query", srv)

    log.Println("Server is running on http://localhost:8080/")
    log.Fatal(http.ListenAndServe(":8080", nil))
}