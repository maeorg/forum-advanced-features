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
	CommentsByCurrentUserAndPosts []CommentAndPost 
	User models.User
}

type CommentAndPost struct {
	Comment models.Comment
	Post models.Post
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
		
		commentsByCurrentUser := services.GetCommentsByUserId(user.Id)
		var commentsByCurrentUserAndPosts []CommentAndPost
		for _, comment := range commentsByCurrentUser {
			post := services.GetPostById(comment.PostId)
			commentsByCurrentUserAndPosts = append(commentsByCurrentUserAndPosts, 
				CommentAndPost{
					Comment: comment,
					Post: post,
			})
		}
		
		activityPage := ActivityPage {
			PostsCreatedByCurrentUser: posts,
			PostsLikedByCurrentUser: likedPosts,
			PostsDislikedByCurrentUser: dislikedPosts,
			CommentsByCurrentUserAndPosts: commentsByCurrentUserAndPosts,
			User: user,
		}

		template.Must(template.ParseFiles("web/static/templates/activity.html")).Execute(w, activityPage)
	}
}