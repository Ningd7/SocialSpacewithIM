package handlers

import (
	"SocialSpace/models"
	"SocialSpace/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser handles GET requests for user information based on username or user ID
func GetUser(c *gin.Context) {
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析JSON失败"})
		return
	}

	var user *models.User
	var err error

	if username, ok := requestData["username"].(string); ok {
		user, err = repository.GetUserByUsername(username)
	} else if userID, ok := requestData["userid"].(float64); ok { // JSON解码时int会被解码为float64
		user, err = repository.GetUserByID(int(userID))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求中必须包含username或userid"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetSomeUsers handles batch user retrieval based on gender or city
func GetSomeUsers(c *gin.Context) {
	method := c.Param("method") // expects 'gender' or 'city'
	searchInformation := c.Query("searchInformation")

	users, err := repository.GetUsers(method, searchInformation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser handles user updates
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析JSON失败"})
		return
	}

	updatedUser, err := repository.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
