package models

type Author struct {
	ID 			int		`json:"id"`
	Name 		string	`json:"name"`
	Biography 	string 	`json:"biography"`
	DateOfBirth string 	`json:"dateOfBirth"`
}