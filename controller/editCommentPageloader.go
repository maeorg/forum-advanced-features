package controller

import (
	"forum/models"
	"forum/services"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type EditCommentPage struct {
	Comment                     models.Comment
	User                     models.User
	NumberOfNewNotifications int
}

func LoadEditCommentPage(w http.ResponseWriter, r *http.Request) {
	commentId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/editComment/"))
	foundComment := services.GetCommentById(commentId)
	if foundComment.Id == 0 {
		template.Must(template.ParseFiles("web/static/templates/404.html")).Execute(w, nil)
		return
	}

	if r.Method == "GET" {
		user := GetCurrentUser(w, r)
		numberOfNewNotifications := services.GetNumberOfNewNotificationsByUserId(user.Id)

		editCommentPage := EditCommentPage{
			Comment:                     foundComment,
			User:                     user,
			NumberOfNewNotifications: numberOfNewNotifications,
		}

		template.Must(template.ParseFiles("web/static/templates/editComment.html")).Execute(w, editCommentPage)
	}

	if r.Method == "POST" {
		content := r.FormValue("content")

		// make new updatedComment object
		updatedComment := models.Comment{
			Id:        foundComment.Id,
			Content:   content,
			CreatedAt: foundComment.CreatedAt,
			UserId:    foundComment.UserId,
			PostId: foundComment.PostId,
		}

		// save updated comment to database
		services.UpdateComment(updatedComment)

		http.Redirect(w, r, "/posts/"+strconv.Itoa(updatedComment.PostId), http.StatusSeeOther)
	}
}
