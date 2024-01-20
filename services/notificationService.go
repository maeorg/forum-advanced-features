package services

import (
	"fmt"
	"forum/models"
	"forum/repository"
)

func SaveNotification(notification models.Notification) {
	err := repository.SaveNotification(notification)
	if err != nil {
		fmt.Println("Error saving notification to database. " + err.Error())
	} else {
		fmt.Println("Saved notification to database")
	}
}