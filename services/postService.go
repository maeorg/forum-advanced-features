package services

import (
	"database/sql"
	"fmt"
	"forum/models"
	"forum/repository"
	"sort"
)

func SavePost(post models.Post) models.Post {
	savedPost, err := repository.SavePost(post)
	if err != nil {
		fmt.Println("Error saving post to database. " + err.Error())
		return models.Post{}
	} else {
		fmt.Println("Saved post to database")
		return formatPosts(savedPost)[0]
	}
}

func UpdatePost(post models.Post) models.Post {
	savedUpdatedPost, err := repository.UpdatePost(post)
	if err != nil {
		fmt.Println("Error saving post to database. " + err.Error())
		return models.Post{}
	} else {
		fmt.Println("Saved post to database")
		return formatPosts(savedUpdatedPost)[0]
	}
}

func DeletePostById(postId int) error {

	// delete post
	_, err := repository.DeletePostById(postId)
	if err != nil {
		fmt.Println("Error removing post from database. " + err.Error())
		return err
	} else {
		fmt.Println("Removed from database post with id ", postId)
	}

	// delete comments of that post
	_, err = repository.DeleteCommentsByPostId(postId)
	if err != nil {
		fmt.Println("Error removing comments from database. " + err.Error())
		return err
	} else {
		fmt.Println("Removed comments from database for post ", postId)
	}

	// delete postCategory info for that post
	_, err = repository.DeleteCategoryPostInfoByPostId(postId)
	if err != nil {
		fmt.Println("Error removing categoryPost info from database. " + err.Error())
		return err
	} else {
		fmt.Println("Removed categoryPost info from database for post ", postId)
	}

	// delete the likes of that post
	_, err = repository.DeleteLikesByPostId(postId)
	if err != nil {
		fmt.Println("Error removing likes from database. " + err.Error())
		return err
	} else {
		fmt.Println("Removed likes from database for post ", postId)
	}

	return nil
}

func GetPostById(postId int) models.Post {
	foundPost, err := repository.GetPostById(postId)
	if err != nil {
		fmt.Println("Error getting post from database. " + err.Error())
		return models.Post{}
	}
	formatedPosts := formatPosts(foundPost)
	if len(formatedPosts) <= 0 {
		fmt.Println("Did not find post with that id")
		return models.Post{}
	}
	return formatedPosts[0]
}

func GetAllPosts() []models.Post {
	foundPosts, err := repository.GetAllPosts()
	if err != nil {
		fmt.Println("Error getting posts from database. " + err.Error())
		return nil
	}
	return formatPosts(foundPosts)
}

func GetPostsByUserId(userId int) []models.Post {
	foundPosts, err := repository.GetPostsByUserId(userId)
	if err != nil {
		fmt.Println("Error getting posts from database. " + err.Error())
		return nil
	}
	return formatPosts(foundPosts)
}

func formatPosts(foundPosts *sql.Rows) []models.Post {
	var id, userId int
	var title, content, createdAt, imageUrl string

	posts := []models.Post{}

	for foundPosts.Next() {
		foundPosts.Scan(&id, &title, &content, &createdAt, &imageUrl, &userId)
		post := models.Post{
			Id:        id,
			Title:     title,
			Content:   content,
			CreatedAt: createdAt,
			ImageUrl:  imageUrl,
			UserId:    userId,
		}
		posts = append(posts, post)
	}

	return posts
}

func GetAllLikedPostsByUserId(userId int) []models.Post {
	foundLikes := GetAllLikesByUserId(userId)

	var likedPosts []models.Post

	for _, like := range foundLikes {
		if like.PostId == 0 {
			continue
		}
		foundPost, err := repository.GetPostById(like.PostId)
		if err != nil {
			fmt.Println("Error getting liked post from database. " + err.Error())
			return nil
		}

		likedPosts = append(likedPosts, formatPosts(foundPost)[0])
	}

	sort.Slice(likedPosts, func(i, j int) bool {
		return likedPosts[i].CreatedAt > likedPosts[j].CreatedAt
	})

	return likedPosts
}

func GetAllDislikedPostsByUserId(userId int) []models.Post {
	foundDislikes := GetAllDislikesByUserId(userId)

	var dislikedPosts []models.Post

	for _, dislike := range foundDislikes {
		if dislike.PostId == 0 {
			continue
		}
		foundPost, err := repository.GetPostById(dislike.PostId)
		if err != nil {
			fmt.Println("Error getting disliked post from database. " + err.Error())
			return nil
		}

		dislikedPosts = append(dislikedPosts, formatPosts(foundPost)[0])
	}

	sort.Slice(dislikedPosts, func(i, j int) bool {
		return dislikedPosts[i].CreatedAt > dislikedPosts[j].CreatedAt
	})

	return dislikedPosts
}

func GetPostsBycategoryId(categoryId int) []models.Post {
	foundPostIds, err := repository.GetPostIdsByCategoryId(categoryId)
	if err != nil {
		fmt.Println("Error getting posts from database. " + err.Error())
		return nil
	}

	posts := []models.Post{}
	var id, postId int
	for foundPostIds.Next() {
		foundPostIds.Scan(&id, &postId, &categoryId)
		foundPost, err := repository.GetPostById(postId)
		if err != nil {
			fmt.Println("Error getting post from database. " + err.Error())
			return nil
		}

		posts = append(posts, formatPosts(foundPost)[0])
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt > posts[j].CreatedAt
	})

	return posts
}

func SaveCategoryPost(categoryPost models.CategoryPost) {
	err := repository.SaveCategoryPost(categoryPost)
	if err != nil {
		fmt.Println("Error saving categoryPost to database. " + err.Error())
	} else {
		fmt.Println("Saved categoryPost to database")
	}
}

func GetPostCategories(post models.Post) []models.Category {
	foundCategories, err := repository.GetPostCategories(post)
	if err != nil {
		fmt.Println("Error getting posts from database. " + err.Error())
		return nil
	}
	return categoryMapper(foundCategories)
}

func categoryMapper(foundCategories *sql.Rows) []models.Category {
	var id int

	categories := []models.Category{}

	for foundCategories.Next() {
		foundCategories.Scan(&id)
		category := GetCategoryById(id)
		categories = append(categories, category)
	}
	return categories
}
