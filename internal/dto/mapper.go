package dto

import "blog-rest/internal/models"

func ToUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToCreateUserRequest(userRequest CreateUserRequest) models.User {
	return models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
}

func ToUpdateUserRequest(userRequest UpdateUserRequest) models.User {
	return models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
}

func ToCategoryResponse(category models.Category) CategoryResponse {
	return CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToCreateCategoryRequest(categoryRequest CreateCategoryRequest) models.Category {
	return models.Category{
		Name: categoryRequest.Name,
	}
}

func ToUpdateCategoryRequest(categoryRequest UpdateCategoryRequest) models.Category {
	return models.Category{
		Name: categoryRequest.Name,
	}
}

func ToPostResponse(post models.Post) PostResponse {
	return PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		User:      ToUserResponse(post.User),
		Category:  ToCategoryResponse(post.Category),
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToCreatePostRequest(postRequest CreatePostRequest) models.Post {
	return models.Post{
		Title:      postRequest.Title,
		Body:       postRequest.Body,
		UserID:     postRequest.UserID,
		CategoryID: postRequest.CategoryID,
	}
}

func ToUpdatePostRequest(postRequest UpdatePostRequest) models.Post {
	return models.Post{
		Title:      postRequest.Title,
		Body:       postRequest.Body,
		CategoryID: postRequest.CategoryID,
	}
}
