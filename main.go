package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var books = []Book{
	{ID: 1, Title: "Book One", Author: "John Doe", Pages: 300, Description: "Test description"},
	{ID: 2, Title: "Book Two", Author: "Steve Smith", Pages: 300, Description: "Test description"},
	{ID: 3, Title: "Book Three", Author: "Mary Jane", Pages: 300, Description: "Test description"},
}
var lastId = 3

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Pages       int    `json:"pages"`
	Description string `json:"description"`
}

type Error struct {
	Message string `json:"message"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "Invalid book ID"})
		return
	}
	for _, item := range books {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Error{Message: "Book not found"})
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "Invalid book ID"})
		return
	}
	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Error{Message: "Book not found"})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "Invalid request payload"})
		return
	}
	book.ID = lastId + 1
	lastId = book.ID
	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "Invalid book ID"})
		return
	}
	for index, item := range books {
		if item.ID == id {
			var book Book
			if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(Error{Message: "Invalid request payload"})
				return
			}
			book.ID = id
			books[index] = book
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Error{Message: "Book not found"})
}

func main() {
	srv := &http.Server{
		Handler:      router(),
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	return r
}
