package handlers

import (
	"SocialSpace/models"
	"SocialSpace/repository"
	"SocialSpace/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var loginDetails models.User
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法解析请求体"})
		return
	}

	user, err := repository.GetUserByUsername(loginDetails.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户未找到"})
		return
	}

	if !utils.VerifyPassword(loginDetails.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "错误的用户名或密码"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"username":   user.Username,
		"gender":     user.Gender,
		"email":      user.Email,
		"profilePic": user.ProfilePic,
		"coverPic":   user.CoverPic,
		"city":       user.City,
		"website":    user.WebSite,
	})
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法解析请求体"})
		return
	}

	if _, err := repository.GetUserByUsername(user.Username); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "用户名已存在"})
		return
	}

	user.Password = utils.EncryptPassword(user.Password)
	user.CoverPic = "./assets/upload/default-cover-pic.jpg"
	user.ProfilePic = "./assets/upload/default-profile-pic.jpg"

	if _, err := repository.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         user.ID,
		"username":   user.Username,
		"gender":     user.Gender,
		"email":      user.Email,
		"profilePic": user.ProfilePic,
		"coverPic":   user.CoverPic,
		"city":       user.City,
		"website":    user.WebSite,
	})
}
