package controller

import (
	"forum/models"
	"forum/services"
	"html/template"
	"net/http"
)

type ActivityPage struct {
	PostsCreatedByCurrentUser []models.Post
	PostsLikedByCurrentUser []models.Post
	PostsDislikedByCurrentUser []models.Post
	User models.User
}

func LoadActivityPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path != "/activity" {
			template.Must(template.ParseFiles("web/static/templates/404.html")).Execute(w, nil)
			return
		}

		user := GetCurrentUser(w, r)
		posts := services.GetPostsByUserId(user.Id)
		likedPosts := services.GetAllLikedPostsByUserId(user.Id)
		dislikedPosts := services.GetAllDislikedPostsByUserId(user.Id)
		
		activityPage := ActivityPage {
			PostsCreatedByCurrentUser: posts,
			PostsLikedByCurrentUser: likedPosts,
			PostsDislikedByCurrentUser: dislikedPosts,
			User: user,
		}

		template.Must(template.ParseFiles("web/static/templates/activity.html")).Execute(w, activityPage)
	}
}