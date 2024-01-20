package controller

import (
	"forum/models"
	"forum/services"
	"html/template"
	"net/http"
	"strconv"
)

type IndexPage struct {
	PostAndLikes []models.PostAndLikes
	Categories   []models.Category
	User         models.User
}

func LoadIndex(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			template.Must(template.ParseFiles("web/static/templates/404.html")).Execute(w, nil)
			return
		}
		
		postsAndLikes := GetAllPosts(w, r)
		categories := services.GetAllCategories()
		user := GetCurrentUser(w, r)

		indexPage := IndexPage{
			PostAndLikes: postsAndLikes,
			Categories:   categories,
			User:         user,
		}

		template.Must(template.ParseFiles("web/static/templates/index.html")).Execute(w, indexPage)
	}

	if r.Method == "POST" {
		chosenFilter := r.FormValue("filter")
		user := GetCurrentUser(w, r)

		var foundPosts []models.Post

		if chosenFilter == "PostsCreatedByUser" {
			foundPosts = services.GetPostsByUserId(user.Id)
		} else if chosenFilter == "PostsLikedByUser" {
			foundPosts = services.GetAllLikedPostsByUserId(user.Id)
		} else {
			category, _ := strconv.Atoi(chosenFilter)

			if category == 0 {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}

			foundPosts = services.GetPostsBycategoryId(category)
		}

		var postsAndLikes []models.PostAndLikes
		for _, post := range foundPosts {
			numberOfLikes := services.GetNumberOfLikesByPostId(post.Id, "like")
			numberOfDislikes := services.GetNumberOfLikesByPostId(post.Id, "dislike")
			categories := services.GetPostCategories(post)
			author, _ := services.GetUserById(post.UserId)
			postsAndLikes = append(postsAndLikes, models.PostAndLikes{
				Post:             post,
				NumberOfLikes:    numberOfLikes,
				NumberOfDislikes: numberOfDislikes,
				Categories:       categories,
				Author: author,
			})
		}

		categories := services.GetAllCategories()

		indexPage := IndexPage{
			PostAndLikes: postsAndLikes,
			Categories:   categories,
			User:         user,
		}
		template.Must(template.ParseFiles("web/static/templates/index.html")).Execute(w, indexPage)
	}
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) []models.PostAndLikes {
	foundPosts := services.GetAllPosts()
	var postsAndLikes []models.PostAndLikes
	for _, post := range foundPosts {
		numberOfLikes := services.GetNumberOfLikesByPostId(post.Id, "like")
		numberOfDislikes := services.GetNumberOfLikesByPostId(post.Id, "dislike")
		categories := services.GetPostCategories(post)
		author, _ := services.GetUserById(post.UserId)
		postsAndLikes = append(postsAndLikes, models.PostAndLikes{
			Post:             post,
			NumberOfLikes:    numberOfLikes,
			NumberOfDislikes: numberOfDislikes,
			Categories:       categories,
			Author:           author,
		})
	}
	return postsAndLikes
}
