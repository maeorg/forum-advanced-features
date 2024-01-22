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

func MarkAllNotificationsToRead() {
	_, err := repository.MarkAllNotificationsToRead()
	if err != nil {
		fmt.Println("Error marking notification to read. " + err.Error())
	} else {
		fmt.Println("Marked all notifications to read")
	}
}

func formatNotifications(foundNotifications *sql.Rows) []models.Notification {
	var id, postId, postCreatorId, userId, read int
	var notificationType, createdAt string

	notifications := []models.Notification{}
	
	for foundNotifications.Next() {
		foundNotifications.Scan(&id, &notificationType, &createdAt, &read, &postId, &postCreatorId, &userId)
		notification := models.Notification{
			Id:         id,
			Type:      notificationType,
			CreatedAt: createdAt,
			Read: read,
			PostId: postId,
			PostCreatorId: postCreatorId,
			UserId: userId,
		}
		notifications = append(notifications, notification)
	}

	return notifications
}
