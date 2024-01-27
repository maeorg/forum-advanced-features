package repository

import (
	"database/sql"
	"forum/models"
)

func SaveCategory(category models.Category) error {
	query := `INSERT INTO categories (name) values (?)`
	_, err := Database.Exec(query, category.Name)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfCategoryExists(name string) *sql.Row {
	var foundCategory *sql.Row
	foundCategory = Database.QueryRow("SELECT name FROM categories WHERE name = ?", name)
	return foundCategory
}

func GetAllCategories() (*sql.Rows, error) {
	foundCategories, err := Database.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	return foundCategories, nil
}

func GetCategoryById(id int) *sql.Row {
	foundCategory := Database.QueryRow("SELECT * FROM categories WHERE id = ?", id)
	return foundCategory
}

func SaveCategoryPost(categoryPost models.CategoryPost) error {
	query := `INSERT INTO categoryPost (post_id, category_id) values (?, ?)`
	_, err := Database.Exec(query, categoryPost.PostId, categoryPost.CategoryId)
	if err != nil {
		return err
	}
	return nil
}

func GetPostCategories(post models.Post) (*sql.Rows, error) {
	foundCategories, err := Database.Query(`SELECT category_id FROM categoryPost WHERE post_id = ?`, post.Id)
	if err != nil {
		return nil, err
	}
	return foundCategories, nil
}

func DeleteCategoryPostInfoByPostId(postId int) (sql.Result, error) {
	result, err := Database.Exec(`DELETE FROM categoryPost WHERE post_id = ?`, postId)
	if err != nil {
		return result, err
	}
	return result, nil
}
