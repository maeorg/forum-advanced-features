package controller

import (
	"forum/models"
	"forum/services"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type EditPostPage struct {
	Post                     models.Post
	ChosenCategories         []models.Category
	User                     models.User
	NumberOfNewNotifications int
	AllCategories            []models.Category
}

func LoadEditPostPage(w http.ResponseWriter, r *http.Request) {
	user := GetCurrentUser(w, r)
	postId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/editPost/"))
	foundPost := services.GetPostById(postId)
	chosenCategories := services.GetPostCategories(foundPost)
	numberOfNewNotifications := services.GetNumberOfNewNotificationsByUserId(user.Id)
	allCategories := services.GetAllCategories()

	editPostPage := EditPostPage{
		Post:                     foundPost,
		ChosenCategories:         chosenCategories,
		User:                     user,
		NumberOfNewNotifications: numberOfNewNotifications,
		AllCategories:            allCategories,
	}

	template.Must(template.ParseFiles("web/static/templates/editPost.html")).Execute(w, editPostPage)
}
