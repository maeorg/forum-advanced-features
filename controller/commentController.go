package controller

import (
	"forum/models"
	"forum/services"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CommentAndLikes struct {
	Comment models.Comment
	NumberOfLikes int
	NumberOfDislikes int
	Author models.User
}

type PostAndComments struct {
	PostAndLikes PostAndLikes
	CommentsAndLikes []CommentAndLikes
	User         models.User
}

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

		http.Redirect(w, r, "/posts/"+strconv.Itoa(comment.PostId), http.StatusSeeOther)
	}
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
			Comment: comment, 
			NumberOfLikes: numberOfLikes, 
			NumberOfDislikes: numberOfDislikes,
			Author: author,
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
	postAndLikes := PostAndLikes {
		Post: foundPost,
		NumberOfLikes: numberOfLikes,
		NumberOfDislikes: numberOfDislikes,
		Author: author,
	}
	
	user := GetCurrentUser(w, r)

	postAndComments := PostAndComments {
		PostAndLikes: postAndLikes,
		CommentsAndLikes: commentsAndLikes,
		User: user,
	}

	template.Must(template.ParseFiles("web/static/templates/posts.html")).Execute(w, postAndComments)
}
