package controller

import (
	"fmt"
	"forum/models"
	"forum/services"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Like(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userId := GetCurrentUser(w, r).Id
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		likeType := "like"
		postId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/likePost/"))
		commentId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/likeComment/"))

		// make new like object
		oppositeLike := models.Like{
			Type:      "dislike",
			UserId:    userId,
			PostId:    postId,
			CommentId: commentId,
		}

		ProcessLike(oppositeLike, likeType, userId, postId, commentId, w, r)
	}
}

func Dislike(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userId := GetCurrentUser(w, r).Id
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		likeType := "dislike"
		postId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/dislikePost/"))
		commentId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/dislikeComment/"))

		// make new like object
		oppositeLike := models.Like{
			Type:      "like",
			UserId:    userId,
			PostId:    postId,
			CommentId: commentId,
		}

		ProcessLike(oppositeLike, likeType, userId, postId, commentId, w, r)
	}
}

func AddNotification(notificationType string, postId, postCreatorUserId, userId int) {
	notification := models.Notification{
		Type:          notificationType,
		CreatedAt:     time.Now().Format(time.RFC3339),
		Read:          0,
		PostId:        postId,
		PostCreatorId: postCreatorUserId,
		UserId:        userId,
	}
	services.SaveNotification(notification)
}

func ProcessLike(oppositeLike models.Like, likeType string, userId, postId, commentId int, w http.ResponseWriter, r *http.Request) {

	// make new like object
	like := models.Like{
		Type:      likeType,
		UserId:    userId,
		PostId:    postId,
		CommentId: commentId,
	}

	var foundLike models.Like
	postCreatorUserId := services.GetPostById(postId).UserId

	// check if the same post/comment is already liked/disliked by the same user
	foundOppositeLike := services.GetLike(oppositeLike.UserId, oppositeLike.PostId, oppositeLike.CommentId, oppositeLike.Type)
	if foundOppositeLike.Id != 0 {
		services.RemoveLike(foundOppositeLike)
		services.SaveLike(like)
		if commentId <= 0 {
			AddNotification(likeType, postId, postCreatorUserId, userId)
		}
		fmt.Println("Post/comment is already liked/disliked by the same user. Reversing likes")
		goto REDIRECT
	}

	// check if like already exists
	foundLike = services.GetLike(like.UserId, like.PostId, like.CommentId, like.Type)
	if foundLike.Id == 0 {
		// save like to database
		services.SaveLike(like)
		if commentId <= 0 {
			AddNotification(likeType, postId, postCreatorUserId, userId)
		}
	} else {
		services.RemoveLike(foundLike)
	}

	REDIRECT:
	isOnPostsPage := strings.HasPrefix(r.Referer(), "https://"+r.Host+"/posts/") || strings.HasPrefix(r.Referer(), "http://"+r.Host+"/posts/")
	if commentId != 0 {
		foundComment := services.GetCommentById(commentId)
		http.Redirect(w, r, "/posts/"+strconv.Itoa(foundComment.PostId), http.StatusSeeOther)
	} else if isOnPostsPage {
		http.Redirect(w, r, "/posts/"+strconv.Itoa(postId), http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
