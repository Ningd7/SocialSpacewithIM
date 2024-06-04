package handler

import (
	"SocialSpace/models"
	"SocialSpace/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "Method ot allowed", http.StatusMethodNotAllowed)
		return
	}
	searchmethod, userdate, err := SearchUserMethod(r)
	if err != nil {
		http.Error(w, "Error occured while searching user", http.StatusInternalServerError)
		return
	}
	var getuser *models.User
	switch searchmethod {
	case "username":
		// 确保userdate是string类型，这里假设SearchUserMethod已经正确处理了类型
		if username, ok := userdate.(string); ok {
			getuser, _, err = repository.GetUserByUsername(username)
		} else {
			http.Error(w, "Invalid data type for username search", http.StatusBadRequest)
			return
		}
	case "userid":
		// 确保userdate是int类型，进行类型断言并转换
		if userId, ok := userdate.(float64); ok { // JSON解码时int会被解码为float64
			getuser, _, err = repository.GetUserbyID(int(userId))
		} else {
			http.Error(w, "Invalid data type for getuser ID search", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Unsupported search method", http.StatusBadRequest)
		return
	}
	respJSON, err := json.MarshalIndent(getuser, "", "  ")
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

func GetSomeUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		http.Error(w, "Method ot allowed", http.StatusMethodNotAllowed)
		return
	}
	method, searchInformation, err := SearchsomeUsersMethod(r)
	if err != nil {
		http.Error(w, "Error occured while searching users", http.StatusInternalServerError)
		return
	}

	getusers, err := repository.GetUsers(method, searchInformation)
	if err != nil {
		http.Error(w, "Error occured while searching users", http.StatusInternalServerError)
		return
	}
	respJSON, err := json.MarshalIndent(getusers, "", "  ")
	if err != nil {
		http.Error(w, "生成响应数据失败", http.StatusInternalServerError)
		log.Printf("序列化响应数据错误: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var newuser models.User
	err := json.NewDecoder(r.Body).Decode(newuser)
	if err != nil {
		http.Error(w, "Error occured while decoding body", http.StatusInternalServerError)
		return
	}
	user, _, err := repository.UpdateUser(newuser)
	if err != nil {
		http.Error(w, "Error occured while updating user", http.StatusInternalServerError)
		return
	}
	respJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		http.Error(w, "Error occured while marshaling body", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func SearchUserMethod(r *http.Request) (string, interface{}, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", nil, fmt.Errorf("读取请求体失败: %w", err)
	}
	defer r.Body.Close()

	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		return "", nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	var methodType string
	var userData interface{}

	if username, ok := requestData["username"]; ok {
		methodType = "username"
		userData = username
	} else if userid, ok := requestData["userid"]; ok {
		methodType = "userid"
		userData = userid
	} else {
		return "", nil, fmt.Errorf("请求中必须包含username或userid")
	}

	return methodType, userData, nil
}

func SearchsomeUsersMethod(r *http.Request) (string, string, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", "", fmt.Errorf("读取请求体失败: %w", err)
	}
	defer r.Body.Close()

	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		return "", "", fmt.Errorf("解析JSON失败: %w", err)
	}

	var methodType string
	var searchdata string

	if _, ok := requestData["gender"]; ok {
		methodType = "gender"
		searchdata = requestData["gender"].(string)
	} else if _, ok := requestData["city"]; ok {
		methodType = "city"
		searchdata = requestData["city"].(string)

	} else {
		return "", "", fmt.Errorf("请求中必须包含gender或city")
	}

	return methodType, searchdata, nil
}
