package controller

import (
	"forum/models"
	"forum/services"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type PostAndLikes struct {
	Post             models.Post
	NumberOfLikes    int
	NumberOfDislikes int
	Categories       []models.Category
	Author           models.User
}

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

func GetAllPosts(w http.ResponseWriter, r *http.Request) []PostAndLikes {
	foundPosts := services.GetAllPosts()
	var postsAndLikes []PostAndLikes
	for _, post := range foundPosts {
		numberOfLikes := services.GetNumberOfLikesByPostId(post.Id, "like")
		numberOfDislikes := services.GetNumberOfLikesByPostId(post.Id, "dislike")
		categories := services.GetPostCategories(post)
		author, _ := services.GetUserById(post.UserId)
		postsAndLikes = append(postsAndLikes, PostAndLikes{
			Post:             post,
			NumberOfLikes:    numberOfLikes,
			NumberOfDislikes: numberOfDislikes,
			Categories:       categories,
			Author:           author,
		})
	}
	return postsAndLikes
}
