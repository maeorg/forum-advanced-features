package repository

import (
	"database/sql"
	"forum/models"
)

func SavePost(post models.Post) (*sql.Rows, error) {
	query := `INSERT INTO posts (title, content, created_at, image_url, user_id) values (?, ?, ?, ?, ?)`
	_, err := Database.Exec(query, post.Title, post.Content, post.CreatedAt, post.ImageUrl, post.UserId)
	if err != nil {
		return nil, err
	}
	savedPost, _ := Database.Query("SELECT * FROM posts WHERE title = ? AND content = ? AND created_at = ? AND user_id = ?", post.Title, post.Content, post.CreatedAt, post.UserId)
	return savedPost, nil
}

func GetAllPosts() (*sql.Rows, error) {
	foundPosts, err := Database.Query("SELECT * FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	return foundPosts, nil
}

func GetPostsByUserId(userId int) (*sql.Rows, error) {
	foundPosts, err := Database.Query("SELECT * FROM posts WHERE user_id = ? ORDER BY created_at DESC", userId)
	if err != nil {
		return nil, err
	}
	return foundPosts, nil
}

func GetPostById(postId int) (*sql.Rows, error) {
	foundPost, err := Database.Query("SELECT * FROM posts WHERE id = ?", postId)
	if err != nil {
		return nil, err
	}
	return foundPost, nil
}

func GetPostIdsByCategoryId(categoryId int) (*sql.Rows, error) {
	foundPostIds, err := Database.Query("SELECT * FROM categoryPost WHERE category_id = ?", categoryId)
	if err != nil {
		return nil, err
	}
	return foundPostIds, nil
}
