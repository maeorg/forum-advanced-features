package controller

import (
	"forum/controller/authentication"
	"forum/models"
	"forum/services"
	"net/http"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request) models.User {
	cookie, err := r.Cookie("session")
	user := models.User{}
	if err == nil {
		username := ""
		for _, userSessionInfo := range authentication.LoggedInUsers {
			if userSessionInfo.SessionIdentifier == cookie.Value {
				username = userSessionInfo.User.Username
			}
		}
		foundUser, _ := services.GetUserByUsername(username)
		user = foundUser
	}
	return user
}
