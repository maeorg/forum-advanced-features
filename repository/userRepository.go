package repository

import (
	"database/sql"
	"forum/models"
)

func GetUserById(userId int) *sql.Row {
	foundUser := Database.QueryRow("SELECT * FROM users WHERE id = ?", userId)
	return foundUser
}

func GetUserByUsername(username string) *sql.Row {
	foundUser := Database.QueryRow("SELECT * FROM users WHERE username = ?", username)
	return foundUser
}

func SaveUser(user models.User) error {
	query := `INSERT INTO users (username, email, password) values (?, ?, ?)`
	_, err := Database.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfUserExists(username, email string) *sql.Row {
	var foundUser *sql.Row
	if username != "" {
		foundUser = Database.QueryRow("SELECT username, email FROM users WHERE username = ?", username)
	} else if email != "" {
		foundUser = Database.QueryRow("SELECT username, email FROM users WHERE email = ?", email)
	}
	return foundUser
}
