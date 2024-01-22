package models

type Notification struct {
	Id int
	Type string
	CreatedAt string
	Read int
	PostId int
	PostCreatorId int
	UserId int
}