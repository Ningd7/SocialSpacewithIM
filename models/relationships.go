package models

type Relationship struct {
	ID             int `json:"id" form:"id" gorm:"AUTO_INCREMENT;primary_key;"`
	FollowerUserId int `json:"followerUserId" form:"followerUserId" gorm:"column:followerUserId"` // 关注人
	FollowedUserId int `json:"followedUserId" form:"followedUserId" gorm:"column:followedUserId"` // 被关注人
}

type FollowCounts struct {
	UserID         int `json:"userId" form:"userId"`
	FollowersCount int `json:"followersCount" form:"followersCount"`
	FollowedCount  int `json:"followedCount" form:"followedCount"`
}

type UserRelations struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
