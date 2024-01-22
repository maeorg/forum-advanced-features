package repository

import (
	"database/sql"
	"forum/models"
)

func SaveNotification(notification models.Notification) error {
	query := `INSERT INTO notifications (type, created_at, read, post_id, post_creator_id, user_id) values (?, ?, ?, ?, ?, ?)`
	_, err := Database.Exec(query, notification.Type, notification.CreatedAt, notification.Read, notification.PostId, notification.PostCreatorId, notification.UserId)
	return err
}

func GetNotificationsByUserId(userId int) (*sql.Rows, error) {
	foundNotifications, err := Database.Query("SELECT * FROM notifications WHERE post_creator_id = ? ORDER BY created_at DESC", userId)
	if err != nil {
		return nil, err
	}
	return foundNotifications, nil
}

func MarkAllNotificationsToRead() (sql.Result, error) {
	result, err := Database.Exec("UPDATE notifications SET read  = ?", 1)
	return result, err
}
