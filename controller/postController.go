package controller

import (
	"fmt"
	"forum/models"
	"forum/services"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")
		createdAt := time.Now().Format(time.RFC3339)
		userId := GetCurrentUser(w, r).Id
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// set image maximum size 20 MB
		maxImgSize := int64(20 * 1024 * 1024)
		var imageURL string

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

			imageURL = "../database/images/" + handler.Filename
		}

		// make new post object
		post := models.Post{
			Title:     title,
			Content:   content,
			CreatedAt: createdAt,
			ImageUrl:  imageURL,
			UserId:    userId,
		}

		// save post to database
		savedPost := services.SavePost(post)
		postId := savedPost.Id

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

func DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/deletePost/"))

		// delete image of the post
		foundPost := services.GetPostById(postId)
		imageUrl := foundPost.ImageUrl
		if imageUrl != "" {
			err := os.Remove(strings.TrimPrefix(imageUrl, "."))
			if err != nil {
				fmt.Println("Error removing post image. " + err.Error())
			} else {
				fmt.Println("Removed image for post with id", postId)
			}
		}

		services.DeletePostById(postId)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// func EditPost(w http.ResponseWriter, r *http.Request) {
// 	// if r.Method == "POST" {
// 		postId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/editPost/"))
// 		foundPost := services.GetPostById(postId)
// 		template.Must(template.ParseFiles("web/static/templates/editPost.html")).Execute(w, foundPost)
// 	// }
// }