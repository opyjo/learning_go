package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{
	{ID: "1", Title: "Learn Go basics", Done: false},
	{ID: "2", Title: "Build first endpoint", Done: true},
	{ID: "3", Title: "Test the API", Done: false},
	{ID: "4", Title: "Deploy the API", Done: false},
	{ID: "5", Title: "Write documentation", Done: false},
	{ID: "6", Title: "Refactor code", Done: false},
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if newTodo.ID == "" {
		newTodo.ID = fmt.Sprintf("%d", len(todos)+1)
	}

	todos = append(todos, newTodo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newTodo); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		addTodo(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/todos", todosHandler)
	log.Println("server starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}