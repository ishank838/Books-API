package controller

import("database/sql"
	"net/http"
	"encoding/json"
	"github.com/ishank838/Books-API/models"
	"github.com/gorilla/mux"
	"log"
	"github.com/ishank838/Books-API/repository"
	"strconv")

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Controller struct {}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter,r *http.Request) {

	books = []models.Book{}

	repo := repository.BookRepository{}

	books = repo.GetAllBooks(db)

	json.NewEncoder(w).Encode(books)
	
	}

}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter,r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	repo := repository.BookRepository{}

	ID, err := strconv.ParseUint(params["id"],10,32)
	logFatal(err)

	book = repo.GetBook(db, uint32(ID))

	json.NewEncoder(w).Encode(book)
	
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc{
	return func (w http.ResponseWriter,r *http.Request) {
	var book models.Book
	var bookID uint32
	
	json.NewDecoder(r.Body).Decode(&book)

	repo := repository.BookRepository{}

	bookID = repo.AddBook(db,book)

	json.NewEncoder(w).Encode(bookID)
	
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter,r *http.Request) {
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)

	log.Println(book)

	repo := repository.BookRepository{}

	rowsUpdated := repo.UpdateBook(db,book)

	json.NewEncoder(w).Encode(rowsUpdated)
	
	}
} 

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter,r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"],10,32)
	logFatal(err)

	repo := repository.BookRepository{}

	resultUpdated := repo.RemoveBook(db, uint32(id))

	json.NewEncoder(w).Encode(resultUpdated)
	
	}
}
