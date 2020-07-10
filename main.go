package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"github.com/ishank838/Books-API/driver"
	"github.com/ishank838/Books-API/controller"
)

func init() {
	//loads enviornment variables
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main(){

	db := driver.ConnectDB()

	controlObj := controller.Controller{}

	r:=mux.NewRouter()
	r.HandleFunc("/books", controlObj.GetBooks(db)).Methods("GET")
	r.HandleFunc("/books/{id}", controlObj.GetBook(db)).Methods("GET")
	r.HandleFunc("/books", controlObj.AddBook(db)).Methods("POST")
	r.HandleFunc("/books", controlObj.UpdateBook(db)).Methods("PUT")
	r.HandleFunc("/books/{id}", controlObj.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
