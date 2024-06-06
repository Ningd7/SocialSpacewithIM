package repository

import (
	"SocialSpace/config"
	"SocialSpace/models"
	"fmt"
	"log"
)

func CreateUser(u models.User) (*models.User, error) {
	result := config.DB.Create(&u)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &u, nil
}

func GetUserByID(uid int) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, uid)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &user, nil
}

func GetUsers(method, searchInformation string) ([]*models.User, error) {
	var users []*models.User
	query := fmt.Sprintf("%s = ?", method)
	result := config.DB.Where(query, searchInformation).Limit(100).Find(&users)
	if result.Error != nil {
		log.Println("Error executing query:", result.Error)
		return nil, result.Error
	}
	return users, nil
}

func UpdateUser(user models.User) (*models.User, error) {
	updatedFields := make(map[string]interface{})

	if user.Username != "" {
		updatedFields["username"] = user.Username
	}
	if user.Gender != "" {
		updatedFields["gender"] = user.Gender
	}
	if user.Email != "" {
		updatedFields["email"] = user.Email
	}
	if user.Password != "" {
		updatedFields["password"] = user.Password
	}
	if user.CoverPic != "" {
		updatedFields["coverPic"] = user.CoverPic
	}
	if user.ProfilePic != "" {
		updatedFields["profilePic"] = user.ProfilePic
	}
	if user.City != "" {
		updatedFields["city"] = user.City
	}
	if user.WebSite != "" {
		updatedFields["webSite"] = user.WebSite
	}

	// 仅更新提供的字段
	result := config.DB.Model(&user).Where("id = ?", user.ID).Updates(updatedFields)
	if result.Error != nil {
		log.Println("Error updating user:", result.Error)
		return nil, result.Error
	}

	return &user, nil
}
