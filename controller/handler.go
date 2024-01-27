package controller

import (
	"forum/controller/authentication"
	"net/http"
)

func Handler() {
	http.Handle("/database/images/", http.StripPrefix("/database/images/", http.FileServer(http.Dir("./database/images"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/", LoadIndex)
	http.HandleFunc("/signup", authentication.RegisterUser)
	http.HandleFunc("/login", authentication.Login)
	http.HandleFunc("/logout", authentication.Logout)
	http.HandleFunc("/createPost", AddPost)
	http.HandleFunc("/addComment/", AddComment)
	http.HandleFunc("/category", AddCategory)
	http.HandleFunc("/likePost/", Like)
	http.HandleFunc("/dislikePost/", Dislike)
	http.HandleFunc("/likeComment/", Like)
	http.HandleFunc("/dislikeComment/", Dislike)
	http.HandleFunc("/posts/", LoadPostAndCommentsByPostId)
	http.HandleFunc("/filter", LoadIndex)
	http.HandleFunc("/activity", LoadActivityPage)
	http.HandleFunc("/notifications", LoadNotificationsPage)
	http.HandleFunc("/deletePost/", DeletePost)
}
