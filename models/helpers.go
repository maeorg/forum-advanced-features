package models

type PostAndLikes struct {
	Post             Post
	NumberOfLikes    int
	NumberOfDislikes int
	Categories       []Category
	Author           User
}
