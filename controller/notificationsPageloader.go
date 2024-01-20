package controller

import (
	"forum/models"
	"forum/services"
	"html/template"
	"net/http"
)

type NotificationsPage struct {
	NotificationsWithUsername []NotificationWithUsername
	User models.User
}

type NotificationWithUsername struct {
	Notification models.Notification
	Username string
}

func LoadNotificationsPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path != "/notifications" {
			template.Must(template.ParseFiles("web/static/templates/404.html")).Execute(w, nil)
			return
		}

		user := GetCurrentUser(w, r)
		notifications := services.GetNotificationsByUserId(user.Id)
		var notificationsWithUsername []NotificationWithUsername
		for _, notification := range notifications {
			foundUser, _ := services.GetUserById(notification.UserId)
			username := foundUser.Username
			notificationWithUsername := NotificationWithUsername {
				Notification: notification,
				Username: username,
			}
			notificationsWithUsername = append(notificationsWithUsername, notificationWithUsername)
		}

		notificationsPage := NotificationsPage {
			NotificationsWithUsername: notificationsWithUsername,
			User: user,
		}

		template.Must(template.ParseFiles("web/static/templates/notifications.html")).Execute(w, notificationsPage)
	}
}