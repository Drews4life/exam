package main

import (
	"./controllers"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var db *sql.DB
var err error

func init() {

	db, err = sql.Open("mysql", "root:Drews123*@tcp(localhost:3306)/books")

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println(" = = = = = = = = = = = = = Success! = = = = = = = = = = = = =\n")
}

func main() {
	router := mux.NewRouter()
	BookController := controllers.BookController{}
	AuthorController := controllers.AuthorController{}

	router.HandleFunc("/books", BookController.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/search/{title}", BookController.GetBooksByTitle(db)).Methods("GET")
	router.HandleFunc("/books/{id}", BookController.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", BookController.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", BookController.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", BookController.RemoveBook(db)).Methods("DELETE")

	router.HandleFunc("/authors", AuthorController.GetAuthors(db)).Methods("GET")
	router.HandleFunc("/authors/search/{name}", AuthorController.GetAuthorsByTitle(db)).Methods("GET")
	router.HandleFunc("/authors/{id}", AuthorController.GetAuthor(db)).Methods("GET")
	router.HandleFunc("/authors", AuthorController.AddAuthor(db)).Methods("POST")
	router.HandleFunc("/authors", AuthorController.UpdateAuthor(db)).Methods("PUT")
	router.HandleFunc("/authors/{id}", AuthorController.RemoveAuthor(db)).Methods("DELETE")

	http.ListenAndServe(":9090", router)
}
