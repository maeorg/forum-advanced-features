package controller

import (
	"fmt"
	"forum/models"
	"forum/repository"
	"forum/services"
	"html/template"
	"io"
	"net/http"
	"os"
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
	postId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/editPost/"))
	foundPost := services.GetPostById(postId)
	if foundPost.Id == 0 {
		template.Must(template.ParseFiles("web/static/templates/404.html")).Execute(w, nil)
		return
	}

	if r.Method == "GET" {
		user := GetCurrentUser(w, r)
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

	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")

		// set image maximum size 20 MB
		maxImgSize := int64(20 * 1024 * 1024)
		imageUrl := foundPost.ImageUrl

		file, handler, err := r.FormFile("img")
		if err == nil {
			defer file.Close()

			filenameSplit := strings.Split(handler.Filename, ".")
			fileExtension := strings.ToLower(filenameSplit[len(filenameSplit)-1])
			if !(fileExtension == "jpg" || fileExtension == "jpeg" || fileExtension == "gif" || fileExtension == "png" || fileExtension == "svg") {
				http.Error(w, "File format not allowed", http.StatusBadRequest)
				return
			}

			if handler.Size > maxImgSize {
				http.Error(w, "File size exceeds 20 MB limit", http.StatusBadRequest)
				return
			}

			f, err := os.Create("./database/images/" + handler.Filename)
			if err != nil {
				http.Error(w, "Cannot post image", http.StatusBadRequest)
				return
			}
			defer f.Close()
			io.Copy(f, file)

			err = os.Remove(strings.TrimPrefix(imageUrl, "."))
			if err != nil {
				fmt.Println("Error removing post image. " + err.Error())
			} else {
				fmt.Println("Removed image for post with id", postId)
			}

			imageUrl = "../database/images/" + handler.Filename
		}

		// make new updatedPost object
		updatedPost := models.Post{
			Id:        foundPost.Id,
			Title:     title,
			Content:   content,
			CreatedAt: foundPost.CreatedAt,
			UserId:    foundPost.UserId,
			ImageUrl:  imageUrl,
		}

		repository.DeleteCategoryPostInfoByPostId(foundPost.Id)

		// save updated post to database
		services.UpdatePost(updatedPost)

		r.ParseForm()
		categories := r.Form["category"]

		for _, category := range categories {
			categoryId, _ := strconv.Atoi(category)
			categoryPost := models.CategoryPost{
				PostId:     postId,
				CategoryId: categoryId,
			}
			services.SaveCategoryPost(categoryPost)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
