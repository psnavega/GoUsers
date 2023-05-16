package repositories

import (
	"api/src/models"
	"database/sql"
)

// Defines repositorie users
type Users struct {
	db *sql.DB
}

// Generate new repositorie
func GenerateRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create a new user at database
func (repository Users) Create(user models.Users) (uint64, error) {
	stmt, err := repository.db.Prepare(
		"INSERT INTO users (name, username, email, password) VALUES(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}
