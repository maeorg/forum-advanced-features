package services

import (
	"database/sql"
	"fmt"
	"forum/models"
	"forum/repository"
)

func SaveLike(like models.Like) {
	err := repository.SaveLike(like)
	if err != nil {
		fmt.Println("Error saving like to database. " + err.Error())
	} else {
		fmt.Println("Saved like to database")
	}
}

func RemoveLike(like models.Like) {
	err := repository.RemoveLike(like)
	if err != nil {
		fmt.Println("Error removing like from database. " + err.Error())
	} else {
		fmt.Println("Removed like from database")
	}
}

func GetAllLikesByUserId(userId int) []models.Like {
	foundLikes, err := repository.GetAllLikesByUserId(userId)
	if err != nil {
		fmt.Println("Error getting likes from database. " + err.Error())
		return nil
	}
	return formatLikes(foundLikes)
}

func GetAllDislikesByUserId(userId int) []models.Like {
	foundDislikes, err := repository.GetAllDislikesByUserId(userId)
	if err != nil {
		fmt.Println("Error getting dislikes from database. " + err.Error())
		return nil
	}
	return formatLikes(foundDislikes)
}

func formatLikes(foundLikes *sql.Rows) []models.Like {
	var id, userId, postId, commentId int
	var likeType string

	likes := []models.Like{}

	for foundLikes.Next() {
		foundLikes.Scan(&id, &likeType, &userId, &postId, &commentId)
		like := models.Like{
			Id:        id,
			Type:      likeType,
			UserId:    userId,
			PostId:    postId,
			CommentId: commentId,
		}
		likes = append(likes, like)
	}

	return likes
}

func formatLike(foundLikes *sql.Row) models.Like {
	var id, userId, postId, commentId int
	var likeType string

	foundLikes.Scan(&id, &likeType, &userId, &postId, &commentId)
	like := models.Like{
		Id:        id,
		Type:      likeType,
		UserId:    userId,
		PostId:    postId,
		CommentId: commentId,
	}

	return like
}

func GetLike(userId, postId, commentId int, likeType string) models.Like {
	return formatLike(repository.GetLike(userId, postId, commentId, likeType))
}

func GetNumberOfLikesByPostId(postId int, likeType string) int {
	foundLikes := repository.GetNumberOfLikesByPostId(postId, likeType)
	var count int
	if err := foundLikes.Scan(&count); err != nil {
		fmt.Println("Error getting number of likes from database. " + err.Error())
		return -1
	}
	return count
}

func GetNumberOfLikesByCommentId(commentId int, likeType string) int {
	foundLikes := repository.GetNumberOfLikesByCommentId(commentId, likeType)
	var count int
	if err := foundLikes.Scan(&count); err != nil {
		fmt.Println("Error getting number of likes from database. " + err.Error())
		return -1
	}
	return count
}
