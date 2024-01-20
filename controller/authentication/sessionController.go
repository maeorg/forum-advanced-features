package authentication

import (
	"fmt"
	"forum/models"
	"forum/services"
	"html/template"
	"net/http"
)

var LoggedInUsers []string

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		template.Must(template.ParseFiles("web/static/templates/login.html")).Execute(w, r)
	}

	if r.Method == "POST" {

		// get username and password from form
		username := r.FormValue("username")
		password := r.FormValue("password")

		// check if user already logged in
		for _, user := range LoggedInUsers {
			if user == username {
				fmt.Println("User already logged in")
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
		}

		// check if username exists and password matches
		foundUser, err := services.GetUserByUsername(username)
		correctPassword := services.CheckIfCorrectPassword(foundUser, password)
		if err != nil || !correctPassword {
			fmt.Println("Wrong username or password")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		} else {
			SetCookieAndLogIn(w, foundUser, r)
		}
	}
}

func SetCookieAndLogIn(w http.ResponseWriter, foundUser models.User, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    foundUser.Username,
		Path:     "/",
		MaxAge:   60 * 60,
		HttpOnly: true,
	})
	LoggedInUsers = append(LoggedInUsers, foundUser.Username)
	fmt.Println("Logged in succesfully. Welcome", foundUser.Username)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		})
		cookie, err := r.Cookie("session")
		if err == nil {
			username := cookie.Value
			for i, user := range LoggedInUsers {
				if user == username {
					LoggedInUsers = append(LoggedInUsers[:i], LoggedInUsers[i+1:]...)
				}
			}
			fmt.Println("Logged out user", username)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
