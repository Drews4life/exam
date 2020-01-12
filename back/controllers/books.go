package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"strconv"

	"../models"
	"../repository/book"
	"../utils"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

var books []models.Book

type BookController struct{}

func (c BookController) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func (c BookController) GetBooksByTitle(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		params := mux.Vars(r)
		title := params["title"]

		log.Print("title: ", title)

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}
		books, err := bookRepo.GetBooksByTitle(db, book, books, title)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func (c BookController) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		params := mux.Vars(r)

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			error.Message = "Incorrect id."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		book, err = bookRepo.GetBook(db, book, id)

		if err != nil {
			error.Message = "No values for current id"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

func (c BookController) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter missing fields."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Could not insert row"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, bookID)
	}
}

func (c BookController) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter all fields."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		spew.Dump(book)
		bookRepo := bookRepository.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		spew.Dump(err)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c BookController) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)

		bookRepo := bookRepository.BookRepository{}

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			error.Message = "Incorrect id."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		rowsDeleted, err := bookRepo.RemoveBook(db, id)

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}