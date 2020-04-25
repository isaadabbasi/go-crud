package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var books []Book = getBooksMocks()

// GetBooks - returns list of books
func GetBooks(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(books)
}

// GetBook - returns a single book
func GetBook(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(res).Encode(book)
			return
		}
	}
	json.NewEncoder(res).Encode(Book{})
}

// CreateBook - [POST] expecting type Book in request body
func CreateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(res).Encode(books)
}

// UpdateBook - expects bookId in params and type Book in request body
func UpdateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, book := range books {
		if params["id"] == book.ID {
			books = append(books[:index], books[index+1:]...)
			var book Book
			json.NewDecoder(req.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(res).Encode(&book)
			break
		}
	}
	json.NewEncoder(res).Encode(books)
}

// DeleteBook - To Delete the book, obv
func DeleteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, book := range books {
		if params["id"] == book.ID {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(books)
}
