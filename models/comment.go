package models

type Comment struct {
	Id int
	Content string
	CreatedAt string
	UserId int
	PostId int
}