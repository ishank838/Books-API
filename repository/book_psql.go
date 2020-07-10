package repository

import("database/sql"
	"log"
	"github.com/ishank838/Books-API/models")

type BookRepository struct {}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (b BookRepository) GetAllBooks(db *sql.DB) []models.Book{
	
	var book models.Book
	var books[] models.Book

	rows, err := db.Query("select * from books;")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID,&book.Title,&book.Author,&book.Year)
		logFatal(err)

		books = append(books,book)
	}

	return books
}

func (b BookRepository) GetBook(db *sql.DB, ID uint32) models.Book {
	var book models.Book

	row := db.QueryRow("Select * from books where id=$1",ID)

	err := row.Scan(&book.ID,&book.Title,&book.Author,&book.Year)
	logFatal(err)

	return book
}

func (b BookRepository) AddBook(db *sql.DB,book models.Book) uint32 {

	var bookID uint32

	err := db.QueryRow("insert into books (title,authur,year) values($1,$2,$3) RETURNING id;",
	book.Title,
	book.Author,
	book.Year).Scan(&bookID)

	logFatal(err)

	return bookID
}

func (b BookRepository) UpdateBook(db *sql.DB,book models.Book) int64{

	log.Println(book)

	res, err := db.Exec("update books set title=$1, authur=$2, year=$3 where id=$4 RETURNING id;",
		book.Title,
		book.Author,
		book.Year,
		book.ID)

	logFatal(err)

	rowsUpdated, err := res.RowsAffected()
	logFatal(err)

	return rowsUpdated
}

func (b BookRepository) RemoveBook(db *sql.DB,ID uint32) int64{

	res, err := db.Exec("delete from books where id=$1;",ID)
	logFatal(err)

	resultUpdated, err := res.RowsAffected()
	logFatal(err)

	return resultUpdated
}