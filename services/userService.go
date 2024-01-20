package services

import (
	"fmt"
	"forum/models"
	"forum/repository"

	"golang.org/x/crypto/bcrypt"
)

func GetUserById(id int) (models.User, error) {
	foundUser := repository.GetUserById(id)

	var username string
	var email string
	var password []byte

	err := foundUser.Scan(&id, &username, &email, &password)

	if err != nil {
		fmt.Println("Error getting user from database. " + err.Error())
		return models.User{}, err
	}

	user := models.User{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
	}

	return user, err
}

func GetUserByUsername(username string) (models.User, error) {
	foundUser := repository.GetUserByUsername(username)

	var id int
	var email string
	var password []byte

	err := foundUser.Scan(&id, &username, &email, &password)

	if err != nil {
		fmt.Println("Error getting user from database. " + err.Error())
		return models.User{}, err
	}

	user := models.User{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
	}

	return user, err
}

func SaveUser(user models.User) {
	err := repository.SaveUser(user)
	if err != nil {
		fmt.Println("Error saving user to database. " + err.Error())
	} else {
		fmt.Println("Saved user to database")
	}
}

func CheckIfExists(username, email string) bool {
	foundUser := repository.CheckIfUserExists(username, email)
	err := foundUser.Scan(&username, &email)
	if err != nil {
		return false
	}
	return true
}

func CheckIfCorrectPassword(user models.User, password string) bool {
	compare := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if compare == nil {
		return true
	}
	return false
}
