package services

import (
	"fmt"
	"forum/models"
	"forum/repository"
)

func SaveCategory(category models.Category) {
	err := repository.SaveCategory(category)
	if err != nil {
		fmt.Println("Error saving category to database. " + err.Error())
	} else {
		fmt.Println("Saved category to database")
	}
}

func CheckIfCategoryExists(category string) bool {
	foundCategory := repository.CheckIfCategoryExists(category)
	err := foundCategory.Scan(&category)
	if err != nil {
		return false
	}
	return true
}

func GetAllCategories() []models.Category {
	foundCategories, err := repository.GetAllCategories()
	if err != nil {
		fmt.Println("Error getting categories from database. " + err.Error())
		return nil
	}

	var id int
	var name string
	categories := []models.Category{}

	for foundCategories.Next() {
		foundCategories.Scan(&id, &name)
		category := models.Category{
			Id:   id,
			Name: name,
		}
		categories = append(categories, category)
	}

	return categories
}

func GetCategoryById(id int) models.Category {
	foundCategory := repository.GetCategoryById(id)

	var name string
	foundCategory.Scan(&id, &name)
	category := models.Category{
		Id:   id,
		Name: name,
	}
	return category
}
