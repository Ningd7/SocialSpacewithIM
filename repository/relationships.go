package repository

import (
	"SocialSpace/config"
	"SocialSpace/models"
	"fmt"
	"log"
)

//var relationships models.Relationship

func CreateFollowRelationship(relationships models.Relationship) error {

	result, err := ExistRelationships(relationships)
	if result {
		return err
	} else {
		result := config.DB.Create(&relationships)
		if result.Error != nil {
			log.Println(result.Error)
			return err
		}
	}
	return nil

}

func DeleteFollowRelationship(relationships models.Relationship) error {

	result, err := ExistRelationships(relationships)
	if !result {
		return err
	} else {
		result := config.DB.Delete(&relationships)
		if result.Error != nil {
			log.Println(result.Error)
			return err
		}
	}
	return nil

}

func GetFollowRelationships(userID int) (*models.FollowCounts, error) {
	var followCounts models.FollowCounts
	followCounts.UserID = userID

	// 查询关注者数量，使用创建的视图 followers_count
	err := config.DB.Table("followers_count"). // 使用正确的视图名
							Select("followersCount"). // 选择正确的列名
							Where("userId = ?", userID).
							Row().Scan(&followCounts.FollowersCount)
	if err != nil {
		log.Println("Error querying followers count from view:", err)
		return nil, err
	}

	// 查询被关注者数量，使用创建的视图 followered_count
	err = config.DB.Table("followered_count"). // 使用正确的视图名
							Select("followeredCount"). // 选择正确的列名
							Where("userId = ?", userID).
							Row().Scan(&followCounts.FollowedCount)
	if err != nil {
		log.Println("Error querying followed count from view:", err)
		return nil, err
	}

	return &followCounts, nil
}

func GetFollowersName(userID int) ([]models.UserRelations, error) {
	var users []models.UserRelations
	err := config.DB.Table("relationships").
		Select("users.id, users.username").
		Joins("join users on users.id = relationships.followerUserId").
		Where("relationships.followedUserId = ?", userID).
		Scan(&users).Error
	if err != nil {
		log.Println("Error querying followers:", err)
		return nil, err
	}
	return users, nil
}

func GetFollowedName(userID int) ([]models.UserRelations, error) {
	var users []models.UserRelations
	err := config.DB.Table("relationships").
		Select("users.id, users.username").
		Joins("join users on users.id = relationships.followedUserId").
		Where("relationships.followerUserId = ?", userID).
		Scan(&users).Error
	if err != nil {
		log.Println("Error querying followed users:", err)
		return nil, err
	}
	return users, nil
}

func ExistRelationships(relationships models.Relationship) (bool, error) {
	result := config.DB.Where("followerUserId = ? AND followedUserId = ?", relationships.FollowerUserId, relationships.FollowedUserId).First(&relationships)
	if result.Error == nil {
		fmt.Printf("关注关系存在！")
		return true, nil
	}
	fmt.Printf("关注关系不存在！")
	return false, result.Error
}
