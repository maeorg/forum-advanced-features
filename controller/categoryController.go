package controller

import (
	"fmt"
	"forum/models"
	"forum/services"
	"net/http"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")

		exists := services.CheckIfCategoryExists(name)
		if exists {
			fmt.Println("Category already exists")
			return
		}

		// make new category object
		category := models.Category{
			Name: name,
		}

		// save category to database
		services.SaveCategory(category)
	}
}
