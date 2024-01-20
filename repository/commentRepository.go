package repository

import (
	"database/sql"
	"forum/models"
)

func SaveComment(comment models.Comment) error {
	query := `INSERT INTO comments (content, created_at, user_id, post_id) values (?, ?, ?, ?)`
	_, err := Database.Exec(query, comment.Content, comment.CreatedAt, comment.UserId, comment.PostId)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCommentsByPostId(postId int) (*sql.Rows, error) {
	foundComments, err := Database.Query("SELECT * FROM comments WHERE post_id = ?  ORDER BY created_at DESC", postId)
	if err != nil {
		return nil, err
	}
	return foundComments, nil
}

func GetCommentById(commentId int) (*sql.Rows, error) {
	foundComment, err := Database.Query("SELECT * FROM comments WHERE id = ?", commentId)
	if err != nil {
		return nil, err
	}
	return foundComment, nil
}