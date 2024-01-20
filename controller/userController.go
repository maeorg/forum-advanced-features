package controller

import (
	"forum/models"
	"forum/services"
	"net/http"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request) models.User {
	cookie, err := r.Cookie("session")
	user := models.User{}
	if err == nil {
		foundUser, _ := services.GetUserByUsername(cookie.Value)
		user = foundUser
	}
	return user
}
