package authorRepository

import (
	"database/sql"
	"../../models"
	"errors"
	"log"
)

type AuthorRepository struct{}

func (a AuthorRepository) GetAuthors(db *sql.DB, author models.Author, authors []models.Author) ([]models.Author, error) {

	rows, err := db.Query("select * from `books`.`author`")

	if err != nil {
		return []models.Author{}, err
	}

	for rows.Next() {
		err = rows.Scan(&author.ID, &author.Name, &author.Biography, &author.DateOfBirth)
		authors = append(authors, author)
	}

	if err != nil {
		return []models.Author{}, err
	}

	return authors, nil
}

func (a AuthorRepository) GetAuthorsByTitle(db *sql.DB, author models.Author, authors []models.Author, name string) ([]models.Author, error) {
	rows, err := db.Query("select * from `books`.`author` where name like " + "'%" + name + "%'");

	for rows.Next() {
		err = rows.Scan(&author.ID, &author.Name, &author.Biography, &author.DateOfBirth)
		authors = append(authors, author)
	}

	if err != nil {
		return []models.Author{}, err
	}

	return authors, nil
}

func (a AuthorRepository) GetAuthor(db *sql.DB, author models.Author, id int) (models.Author, error) {
	rows := db.QueryRow("select * from `books`.`author` where id=?", id)
	err := rows.Scan(&author.ID, &author.Name, &author.Biography, &author.DateOfBirth)

	if err != nil {
		return models.Author{}, err
	}

	return author, nil
}

func (a AuthorRepository) AddAuthor(db *sql.DB, author models.Author) (int, error) {
	_, err := db.Exec("insert into `books`.`author` (`name`, `biography`, `dateOfBirth`) values (?, ?, ?)",
		author.Name, author.Biography, author.DateOfBirth)

	if err != nil {
		log.Println(err)
		return 0, errors.New("Could not insert")
	}

	return 1, nil
}

func (a AuthorRepository) UpdateAuthor(db *sql.DB, author models.Author) (int64, error) {
	result, err := db.Exec("update `books`.`author` set name=?, biography=?, dateOfBirth=? where id=?",
		&author.Name, &author.Biography, &author.DateOfBirth, &author.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (a AuthorRepository) RemoveAuthor(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from `books`.`author` where id=?", id)

	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}