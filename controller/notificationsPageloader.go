package controller

import (
	"html/template"
	"net/http"
)

func LoadNotificationsPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path != "/notifications" {
			template.Must(template.ParseFiles("web/static/templates/404.html")).Execute(w, nil)
			return
		}

		template.Must(template.ParseFiles("web/static/templates/notifications.html")).Execute(w, nil)
	}
}