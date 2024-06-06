package main

import (
	"SocialSpace/config"
	"SocialSpace/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	config.InitDB()

	router := gin.Default()

	// 用户操作路由
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
	router.GET("/user", handlers.GetUser)
	router.GET("/users/:method", handlers.GetSomeUsers)
	router.PUT("/user", handlers.UpdateUser)

	router.POST("/create-relationships", handlers.CreateRelationships)
	router.DELETE("/delete-relationships", handlers.DeleteRelationships)
	router.GET("/get-follow-relationships", handlers.GetFollowRelationships)

	router.GET("/followers", handlers.GetFollowersHandler)
	router.GET("/following", handlers.GetFollowingHandler)

	// 启动服务
	router.Run(":8080")
}
