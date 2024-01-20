package services

import (
	"database/sql"
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

func GetNotificationsByUserId(userId int) []models.Notification {
	foundNotifications, err := repository.GetNotificationsByUserId(userId)
	if err != nil {
		fmt.Println("Error getting posts from database. " + err.Error())
		return nil
	}
	return formatNotifications(foundNotifications)
}

func formatNotifications(foundNotifications *sql.Rows) []models.Notification {
	var id, postId, postCreatorId, userId int
	var notificationType, createdAt string

	notifications := []models.Notification{}
	
	for foundNotifications.Next() {
		foundNotifications.Scan(&id, &notificationType, &createdAt, &postId, &postCreatorId, &userId)
		notification := models.Notification{
			Id:         id,
			Type:      notificationType,
			CreatedAt: createdAt,
			PostId: postId,
			PostCreatorId: postCreatorId,
			UserId: userId,
		}
		notifications = append(notifications, notification)
	}

	return notifications
}
