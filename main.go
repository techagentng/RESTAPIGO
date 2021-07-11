package main

import (
	//"encoding/json"
    "log"
	//"math/rand"
	//"strconv"
	"net/http"
	"github.com/gorilla/mux"
)

// Book Struct(model) like the es6 class in javascript
type Book struct {
	ID		string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author	`json:"author"`
}

// Author contruct
type Author struct {
	FirstName string `json:firstname`
	LastName string `json:lastname`
}

// Get all books
func getBook(w http.ResponseWriter, r *http.Request){
	
}
// Create books
func createBook(w http.ResponseWriter, r *http.Request){
	
}
// Update books
func updateBook(w http.ResponseWriter, r *http.Request){
	
}
// Delete books
func deleteBook(w http.ResponseWriter, r *http.Request){
	
}

func main(){
	// init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBook).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}