package authentication

import (
	"fmt"
	"forum/models"
	"forum/services"
	"html/template"
	"net/http"

	"github.com/gofrs/uuid/v5"
	_ "github.com/gofrs/uuid/v5"
)

var LoggedInUsers []UserSessionInfo

type UserSessionInfo struct {
	User              models.User
	SessionIdentifier string
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		template.Must(template.ParseFiles("web/static/templates/login.html")).Execute(w, r)
	}

	if r.Method == "POST" {

		// get username and password from form
		username := r.FormValue("username")
		password := r.FormValue("password")

		// check if username exists and password matches
		foundUser, err := services.GetUserByUsername(username)
		correctPassword := services.CheckIfCorrectPassword(foundUser, password)
		if err != nil || !correctPassword {
			fmt.Println("Wrong username or password")
			http.Error(w, "Wrong username or password", http.StatusBadRequest)
			return
		} else {
			// check if user already logged in
			for i, userSessionInfo := range LoggedInUsers {
				if userSessionInfo.User.Username == username {
					fmt.Println("User already logged in. Logging out from other places and logging in from here.")
					LoggedInUsers = append(LoggedInUsers[:i], LoggedInUsers[i+1:]...)
				}
			}
			SetCookieAndLogIn(w, foundUser, r)
		}
	}
}

func SetCookieAndLogIn(w http.ResponseWriter, foundUser models.User, r *http.Request) {
	generatedUuid, _ := uuid.NewV4()
	sessionIdentifier := generatedUuid.String()

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionIdentifier,
		Path:     "/",
		MaxAge:   60 * 60,
		HttpOnly: true,
	})

	userSessionInfo := UserSessionInfo{
		User:              foundUser,
		SessionIdentifier: sessionIdentifier,
	}

	LoggedInUsers = append(LoggedInUsers, userSessionInfo)
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
			sessionIdentifier := cookie.Value
			for i, userSessionInfo := range LoggedInUsers {
				if userSessionInfo.SessionIdentifier == sessionIdentifier {
					LoggedInUsers = append(LoggedInUsers[:i], LoggedInUsers[i+1:]...)
				}
			}
			fmt.Println("Logged out user with uuid", sessionIdentifier)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
