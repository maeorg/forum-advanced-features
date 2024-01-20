package models

type Notification struct {
	Id int
	Type string
	CreatedAt string
	PostId int
	PostCreatorId int
	UserId int
}