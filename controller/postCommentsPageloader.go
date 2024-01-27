package controller

import (
	"forum/models"
	"forum/services"
	"html/template"
	"net/http"
	"strconv"
)

type CommentAndLikes struct {
	Comment          models.Comment
	NumberOfLikes    int
	NumberOfDislikes int
	Author           models.User
}

type PostAndCommentsPage struct {
	PostAndLikes             models.PostAndLikes
	CommentsAndLikes         []CommentAndLikes
	User                     models.User
	NumberOfNewNotifications int
}

func LoadPostAndCommentsByPostId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/posts/"):]
	postId, _ := strconv.Atoi(id)
	postComments := services.GetAllCommentsByPostId(postId)

	var commentsAndLikes []CommentAndLikes
	for _, comment := range postComments {
		numberOfLikes := services.GetNumberOfLikesByCommentId(comment.Id, "like")
		numberOfDislikes := services.GetNumberOfLikesByCommentId(comment.Id, "dislike")
		author, _ := services.GetUserById(comment.UserId)
		commentsAndLikes = append(commentsAndLikes, CommentAndLikes{
			Comment:          comment,
			NumberOfLikes:    numberOfLikes,
			NumberOfDislikes: numberOfDislikes,
			Author:           author,
		})
	}

	foundPost := services.GetPostById(postId)
	if foundPost.Id == 0 {
		template.Must(template.ParseFiles("web/static/templates/404.html")).Execute(w, nil)
		return
	}
	numberOfLikes := services.GetNumberOfLikesByPostId(postId, "like")
	numberOfDislikes := services.GetNumberOfLikesByPostId(postId, "dislike")
	author, _ := services.GetUserById(foundPost.UserId)
	postAndLikes := models.PostAndLikes{
		Post:             foundPost,
		NumberOfLikes:    numberOfLikes,
		NumberOfDislikes: numberOfDislikes,
		Author:           author,
	}

	user := GetCurrentUser(w, r)

	numberOfNewNotifications := services.GetNumberOfNewNotificationsByUserId(user.Id)
	postAndComments := PostAndCommentsPage{
		PostAndLikes:             postAndLikes,
		CommentsAndLikes:         commentsAndLikes,
		User:                     user,
		NumberOfNewNotifications: numberOfNewNotifications,
	}

	template.Must(template.ParseFiles("web/static/templates/posts.html")).Execute(w, postAndComments)
}
