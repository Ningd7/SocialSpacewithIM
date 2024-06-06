package handlers

import (
	"SocialSpace/models"
	"SocialSpace/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateRelationships(c *gin.Context) {
	var relationships models.Relationship
	var err error
	if relationships.FollowerUserId, err = strconv.Atoi(c.Query("followerUserId")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid followerUserId"})
		return
	}

	// 解析 followedUserId
	if relationships.FollowedUserId, err = strconv.Atoi(c.Query("followedUserId")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid followedUserId"})
		return
	}
	if relationships.FollowerUserId == relationships.FollowedUserId {
		fmt.Printf("不能自己关注自己！")
		return
	}
	err = repository.CreateFollowRelationship(relationships)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"relationships": relationships})
	return
}

func DeleteRelationships(c *gin.Context) {
	var relationships models.Relationship
	var err error
	if relationships.FollowerUserId, err = strconv.Atoi(c.Query("followerUserId")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid followerUserId"})
		return
	}

	// 解析 followedUserId
	if relationships.FollowedUserId, err = strconv.Atoi(c.Query("followedUserId")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid followedUserId"})
		return
	}
	err = repository.DeleteFollowRelationship(relationships)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "取消关注成功"})
	return
}

func GetFollowRelationships(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
		return
	}
	var followersCount *models.FollowCounts
	followersCount, err = repository.GetFollowRelationships(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"followersCount": followersCount})
}

func GetFollowersHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
		return
	}
	followers, err := repository.GetFollowersName(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, followers)
}

func GetFollowingHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
		return
	}
	following, err := repository.GetFollowedName(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, following)
}
