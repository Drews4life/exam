package bookRepository

import (
	"../../models"
	"database/sql"
	"errors"
	"log"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {

	rows, err := db.Query("select * from `books`.`book`")

	if err != nil {
		return []models.Book{}, err
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func (b BookRepository) GetBooksByTitle(db *sql.DB, book models.Book, books []models.Book, title string) ([]models.Book, error) {
	rows, err := db.Query("select * from `books`.`book` where title like " + "'%" + title + "%'");

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) (models.Book, error) {
	rows := db.QueryRow("select * from `books`.`book` where id=?", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) (int, error) {
	_, err := db.Exec("insert into `books`.`book` (`title`, `author`, `year`) values (?, ?, ?)",
		book.Title, book.Author, book.Year)

	//log.Println(q + "\n")
	if err != nil {
		log.Println(err)
		return 0, errors.New("Could not insert")
	}

	return 1, nil
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) (int64, error) {
	result, err := db.Exec("update `books`.`book` set title=?, author=?, year=? where id=?",
		&book.Title, &book.Author, &book.Year, &book.ID)
		log.Println("book: ", book)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from `books`.`book` where id=?", id)

	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}
