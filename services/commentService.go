package services

import (
	"database/sql"
	"fmt"
	"forum/models"
	"forum/repository"
)

func SaveComment(comment models.Comment) {
	err := repository.SaveComment(comment)
	if err != nil {
		fmt.Println("Error saving comment to database. " + err.Error())
	} else {
		fmt.Println("Saved comment to database")
	}
}

func GetAllCommentsByPostId(postId int) []models.Comment {
	foundComments, err := repository.GetAllCommentsByPostId(postId)
	if err != nil {
		fmt.Println("Error getting comments from database. " + err.Error())
		return nil
	}
	return formatComments(foundComments)
}

func formatComments(foundComments *sql.Rows) []models.Comment {
	var id, userId, postId int
	var content, createdAt string

	comments := []models.Comment{}

	for foundComments.Next() {
		foundComments.Scan(&id, &content, &createdAt, &userId, &postId)
		comment := models.Comment{
			Id:        id,
			Content:   content,
			CreatedAt: createdAt,
			UserId:    userId,
			PostId:    postId,
		}
		comments = append(comments, comment)
	}

	return comments
}

func GetCommentById(commentId int) models.Comment {
	foundComment, err := repository.GetCommentById(commentId)
	if err != nil {
		fmt.Println("Error getting comment from database. " + err.Error())
		return models.Comment{}
	}
	return formatComments(foundComment)[0]
}

func GetCommentsByUserId(userId int) []models.Comment {
	foundComments, err := repository.GetCommentsByUserId(userId)
	if err != nil {
		fmt.Println("Error getting comments from database. " + err.Error())
		return nil
	}
	return formatComments(foundComments)
}

func DeleteCommentById(commentId int) error {
	_, err := repository.DeleteCommentById(commentId)
	if err != nil {
		fmt.Println("Error removing comment from database. " + err.Error())
		return err
	} else {
		fmt.Println("Removed from database comment with id ", commentId)
	}

	repository.DeleteLikesByCommentId(commentId)

	return nil
}

func UpdateComment(comment models.Comment) models.Comment {
	savedUpdatedComment, err := repository.UpdateComment(comment)
	if err != nil {
		fmt.Println("Error updating comment in database. " + err.Error())
		return models.Comment{}
	} else {
		fmt.Println("Updated comment in database")
		return formatComments(savedUpdatedComment)[0]
	}
}