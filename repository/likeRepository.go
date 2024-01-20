package repository

import (
	"database/sql"
	"forum/models"
)

func GetLike(userId, postId, commentId int, likeType string) *sql.Row {
	var foundLike *sql.Row
	if postId != 0 {
		foundLike = Database.QueryRow("SELECT * FROM likes WHERE type = ? AND user_id = ? AND post_id = ?", likeType, userId, postId)
	} else if commentId != 0 {
		foundLike = Database.QueryRow("SELECT * FROM likes WHERE type = ? AND user_id = ? AND comment_id = ?", likeType, userId, commentId)
	}
	return foundLike
}

func SaveLike(like models.Like) error {
	query := `INSERT INTO likes (type, user_id, post_id, comment_id) values (?, ?, ?, ?)`
	_, err := Database.Exec(query, like.Type, like.UserId, like.PostId, like.CommentId)
	if err != nil {
		return err
	}
	return nil
}

func RemoveLike(like models.Like) error {
	query := `DELETE FROM likes WHERE id = ?`
	_, err := Database.Exec(query, like.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllLikesByUserId(userId int) (*sql.Rows, error) {
	foundLikes, err := Database.Query(`SELECT * FROM likes WHERE user_id = ? AND type = "like"`, userId)
	if err != nil {
		return nil, err
	}
	return foundLikes, nil
}

func GetAllLikesByCommentId(commentId int, likeType string) (*sql.Rows, error) {
	foundLikes, err := Database.Query(`SELECT * FROM likes WHERE comment_id = ? AND type = ?`, commentId, likeType)
	if err != nil {
		return nil, err
	}
	return foundLikes, nil
}

func GetAllLikesByPostId(postId int, likeType string) (*sql.Rows, error) {
	foundLikes, err := Database.Query(`SELECT * FROM likes WHERE post_id = ? AND type = ?`, postId, likeType)
	if err != nil {
		return nil, err
	}
	return foundLikes, nil
}

func GetNumberOfLikesByPostId(postId int, likeType string) *sql.Row {
	foundLikes := Database.QueryRow(`SELECT COUNT(*) FROM likes WHERE post_id = ? AND type = ?`, postId, likeType)
	return foundLikes
}

func GetNumberOfLikesByCommentId(commentId int, likeType string) *sql.Row {
	foundLikes := Database.QueryRow(`SELECT COUNT(*) FROM likes WHERE comment_id = ? AND type = ?`, commentId, likeType)
	return foundLikes
}
