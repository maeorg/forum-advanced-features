package controller

import (
	"forum/models"
	"forum/services"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content := r.FormValue("content")
		createdAt := time.Now().Format(time.RFC3339)
		userId := GetCurrentUser(w, r).Id
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		postId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/addComment/"))

		// make new comment object
		comment := models.Comment{
			Content:   content,
			CreatedAt: createdAt,
			UserId:    userId,
			PostId:    postId,
		}

		// save comment to database
		services.SaveComment(comment)

		postCreatorUserId := services.GetPostById(postId).UserId
		AddNotification("comment", postId, postCreatorUserId, userId)

		http.Redirect(w, r, "/posts/"+strconv.Itoa(comment.PostId), http.StatusSeeOther)
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		commentId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/deleteComment/"))
		services.DeleteCommentById(commentId)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
