package repository

import "forum/models"

func SaveNotification(notification models.Notification) error {
	query := `INSERT INTO notifications (type, created_at, post_id, post_creator_id, user_id) values (?, ?, ?, ?, ?)`
	_, err := Database.Exec(query, notification.Type, notification.CreatedAt, notification.PostId, notification.PostCreatorId, notification.UserId)
	if err != nil {
		return err
	}
	return nil
}