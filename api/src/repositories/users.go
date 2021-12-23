package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represents a user repository
type Users struct {
	db *sql.DB
}

// NewUserRepository creates a repository to handle users
func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create adds a new user to the database
func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}
