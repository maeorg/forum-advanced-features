package models

type Like struct {
	Id int
	Type string
	UserId int
	PostId int
	CommentId int
}