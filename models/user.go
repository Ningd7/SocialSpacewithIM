package models

type User struct {
	ID         int    `json:"id" form:"id"`                 // 用户编号
	Username   string `json:"username" form:"username"`     // 用户名
	Gender     string `json:"gender" form:"gender"`         // 性别
	Email      string `json:"email" form:"email"`           // 邮箱
	Password   string `json:"password" form:"password"`     // 密码
	CoverPic   string `json:"coverPic" form:"coverPic"`     // 背景图
	ProfilePic string `json:"profilePic" form:"profilePic"` // 头像
	City       string `json:"city" form:"city"`             // 城市
	WebSite    string `json:"webSite" form:"webSite"`       // 个人网站
}
