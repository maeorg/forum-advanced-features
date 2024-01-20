package authentication

import (
	"fmt"
	"forum/models"
	"forum/services"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		template.Must(template.ParseFiles("web/static/templates/signup.html")).Execute(w, r)
	}

	if r.Method == "POST" {
		var err error

		// check if email or username is already taken
		username := r.FormValue("username")
		email := r.FormValue("email")
		exists := services.CheckIfExists(username, "")
		if exists {
			fmt.Println("Username already registered")
			http.Redirect(w, r, "/signup", http.StatusFound)
			return
		}
		exists = services.CheckIfExists("", email)
		if exists {
			fmt.Println("Email already registered")
			http.Redirect(w, r, "/signup", http.StatusFound)
			return
		}

		// encrypt password
		password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Failed to encrypt password")
			return
		}

		// make new user object
		user := models.User {
			Username: username,
			Email: email,
			Password: password,
		}

		// save user to database
		services.SaveUser(user)
		
		fmt.Println("Succesfully created user")
		SetCookieAndLogIn(w, user, r)
	}
}