package utils

import (
	"SocialSpace/models"
	"SocialSpace/repository"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func DecodeJSONBody(r *http.Request, dst interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err // 读取请求体失败
	}
	defer r.Body.Close()

	// 尝试将JSON数据解码到dst指向的结构体中
	if err := json.Unmarshal(body, dst); err != nil {
		return err // JSON解码失败
	}

	return nil // 数据解析成功
}

func EncryptPassword(password string) string {
	// password是明文的密码
	md := md5.New()
	md.Write([]byte(password))
	return hex.EncodeToString(md.Sum(nil))
}

// http://127.0.0.1:8080

func CheckToken(r *http.Request) (*models.User, bool) {
	// 获取cookie
	cookie := r.Header.Get("AccessToken")
	temp := strings.Split(cookie, "=")
	if len(temp) <= 1 {
		return nil, false
	}
	username := strings.ReplaceAll(temp[1], ";", "")
	// 根据cookie里的用户信息查询当前用户
	user, _, err := repository.GetUserByUsername(username)
	if err != nil {
		return nil, false
	}
	// 返回结果
	return user, true
}

func JSON(w http.ResponseWriter, x interface{}) {
	// 4. 返回结果
	// 4.1 响应数据做序列化
	res, _ := json.Marshal(x)
	// 4.2 设置content-type
	w.Header().Set("Content-Type", "application/json")
	// 4.3 设置状态码
	w.WriteHeader(http.StatusOK)
	// 4.4 返回结果。
	w.Write(res)
}

// 跨域处理
func CORS(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("access-control-allow-headers", "content-type,accesstoken,x-xsrf-token,authorization,token")
		w.Header().Set("access-control-allow-credentials", "true")
		w.Header().Set("access-control-allow-methods", "post,get,delete,put,options")
		w.Header().Set("access-type", "application/json;charset=utf8-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		f(w, r)
	}
}
