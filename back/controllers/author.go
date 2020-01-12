package controllers

import (
	"../models"
	"../repository/author"
	"../utils"
	"database/sql"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var authors []models.Author

type AuthorController struct{}

func (c AuthorController) GetAuthors(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author models.Author
		var error models.Error

		authors = []models.Author{}
		authorRepo := authorRepository.AuthorRepository{}
		authors, err := authorRepo.GetAuthors(db, author, authors)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, authors)
	}
}

func (c AuthorController) GetAuthorsByTitle(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author models.Author
		var error models.Error

		params := mux.Vars(r)
		name := params["name"]

		authors = []models.Author{}
		authorRepo := authorRepository.AuthorRepository{}
		authors, err := authorRepo.GetAuthorsByTitle(db, author, authors, name)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, authors)
	}
}

func (c AuthorController) GetAuthor(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author models.Author
		var error models.Error

		params := mux.Vars(r)

		authors = []models.Author{}
		authorRepo := authorRepository.AuthorRepository{}

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			error.Message = "Incorrect id."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		author, err = authorRepo.GetAuthor(db, author, id)

		if err != nil {
			error.Message = "No values for current id"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, author)
	}
}

func (c AuthorController) AddAuthor(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author models.Author
		var authorID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&author)

		if author.Name == "" || author.Biography == "" || author.DateOfBirth == "" {
			error.Message = "Enter missing fields."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		authorRepo := authorRepository.AuthorRepository{}
		authorID, err := authorRepo.AddAuthor(db, author)

		if err != nil {
			error.Message = "Could not insert row"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, authorID)
	}
}

func (c AuthorController) UpdateAuthor(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author models.Author
		var error models.Error

		json.NewDecoder(r.Body).Decode(&author)

		if author.Name == "" || author.Biography == "" || author.DateOfBirth == "" {
			error.Message = "Enter all fields."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		authorRepo := authorRepository.AuthorRepository{}
		rowsUpdated, err := authorRepo.UpdateAuthor(db, author)

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

func (c AuthorController) RemoveAuthor(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)

		authorRepo := authorRepository.AuthorRepository{}

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			error.Message = "Incorrect id."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		rowsDeleted, err := authorRepo.RemoveAuthor(db, id)

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}