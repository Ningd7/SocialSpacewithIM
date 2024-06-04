package handler

import (
	"SocialSpace/models"
	"SocialSpace/repository"
	"SocialSpace/utils"
	"encoding/json"
	"log"
	"net/http"
)

// http://127.0.0.1:8080/login POST
func Login(w http.ResponseWriter, r *http.Request) {
	// 1. 判断请求方式
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// 2. 获取参数
	u := models.User{}
	err := utils.DecodeJSONBody(r, &u)

	// 3. 调用controller
	user, _, err := repository.GetUserByUsername(u.Username)
	if err != nil {
		http.Error(w, "user not found!", http.StatusBadRequest)
		return
	}
	// 密码验证
	if !VerifyPassword(u.Password, user.Password) {
		http.Error(w, "Wrong password or username!", http.StatusBadRequest)
		return
	}
	// 4. 签发token
	//cookie := http.Cookie{
	//	Name:  "username",
	//	Value: user.Username,
	//}
	//w.Header().Set("accessToken", cookie.String()) // 将用户信息存入cookie中
	// 5. 返回请求结果
	respUser := map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"gender":     user.Gender,
		"email":      user.Email,
		"profilePic": user.ProfilePic,
		"coverPic":   user.CoverPic,
		"city":       user.City,
		"website":    user.WebSite,
	}

	// 序列化响应数据
	respJSON, err := json.MarshalIndent(respUser, "", "  ")
	if err != nil {
		http.Error(w, "生成响应数据失败", http.StatusInternalServerError)
		log.Printf("序列化响应数据错误: %v", err)
		return
	}

	// 写入响应
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(respJSON)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "请求方法必须为POST", http.StatusMethodNotAllowed)
		return
	}

	// 使用utils.DecodeJSONBody解码请求体到User结构体
	var user models.User
	if err := utils.DecodeJSONBody(r, &user); err != nil {
		http.Error(w, "无法解析请求体", http.StatusBadRequest)
		log.Printf("解析请求体错误: %v", err)
		return
	}

	// 查询用户是否已存在
	existingUser, _, err := repository.GetUserByUsername(user.Username)
	if err != nil {
		http.Error(w, "查询用户信息时发生错误", http.StatusInternalServerError)
		log.Printf("查询用户错误: %v", err)
		return
	}
	if existingUser != nil {
		http.Error(w, "用户名已存在", http.StatusConflict)
		return
	}

	// 加密密码
	hashedPassword := utils.EncryptPassword(user.Password)
	//if err != nil {
	//	http.Error(w, "密码加密失败", http.StatusInternalServerError)
	//	log.Printf("密码加密错误: %v", err)
	//	return
	//}
	user.Password = hashedPassword

	// 设置默认头像
	user.CoverPic = "./assets/upload/default-cover-pic.jpg"
	user.ProfilePic = "./assets/upload/default-profile-pic.jpg"

	// 创建用户
	createdUser, _, err := repository.CreateUser(user)
	if err != nil {
		http.Error(w, "创建用户失败", http.StatusInternalServerError)
		log.Printf("创建用户错误: %v", err)
		return
	}

	// 准备响应数据
	respUser := map[string]interface{}{
		"id":         createdUser.ID,
		"username":   createdUser.Username,
		"gender":     createdUser.Gender,
		"email":      createdUser.Email,
		"profilePic": createdUser.ProfilePic,
		"coverPic":   createdUser.CoverPic,
		"city":       createdUser.City,
		"website":    createdUser.WebSite,
	}

	// 序列化响应数据
	respJSON, err := json.MarshalIndent(respUser, "", "  ")
	if err != nil {
		http.Error(w, "生成响应数据失败", http.StatusInternalServerError)
		log.Printf("序列化响应数据错误: %v", err)
		return
	}

	// 写入响应
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(respJSON)
}

func VerifyPassword(password, encryptPassword string) bool {
	// password是明文的密码， encryptPassword是加密后的密码

	return encryptPassword == utils.EncryptPassword(password)
}
