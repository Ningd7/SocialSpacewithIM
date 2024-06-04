package repository

import (
	"SocialSpace/config"
	"SocialSpace/models"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

var db *sql.DB

//type models.User struct {
//	ID         int    `json:"id" form:"id"`                 // 用户编号
//	Username   string `json:"username" form:"username"`     // 用户名
//	Gender     string `json:"gender" form:"gender"`         // 性别
//	Email      string `json:"email" form:"email"`           // 邮箱
//	Password   string `json:"password" form:"password"`     // 密码
//	CoverPic   string `json:"coverPic" form:"coverPic"`     // 背景图
//	ProfilePic string `json:"profilePic" form:"profilePic"` // 头像
//	City       string `json:"city" form:"city"`             // 城市
//	WebSite    string `json:"webSite" form:"webSite"`       // 个人网站
//}

func init() {
	db = config.GetDB()
}

func CreateUser(u models.User) (*models.User, *sql.DB, error) {
	userStr := "insert into users(username, gender, email, password, name, coverPic, profilePic, city, website) value(?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(userStr)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Username, u.Gender, u.Email, u.Password, u.CoverPic, u.ProfilePic, u.City, u.WebSite)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	id, err := result.LastInsertId()
	u.ID = int(id)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	return &u, db, err
}

func GetUserbyID(uid int) (*models.User, *sql.DB, error) {
	userstr := "select * from users where id=? limit 1"
	stmt, err := db.Prepare(userstr)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer stmt.Close()
	row, err := stmt.Query(uid)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	user := models.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Gender, &user.Email, &user.Password, &user.CoverPic, &user.ProfilePic, &user.City, &user.WebSite)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	return &user, db, err
}
func GetUserByUsername(username string) (*models.User, *sql.DB, error) {
	userstr := "select * from users where username=? limit 1"
	stmt, err := db.Prepare(userstr)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer stmt.Close()
	row, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	user := models.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Gender, &user.Email, &user.Password, &user.CoverPic, &user.ProfilePic, &user.City, &user.WebSite)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	return &user, db, err
}

func GetUsers(method string, searchInformation string) ([]*models.User, error) {
	users := make([]*models.User, 0)

	// 安全地构建SQL查询，防止SQL注入
	var query string
	switch method {
	case "city", "gender":
		query = fmt.Sprintf("SELECT * FROM users WHERE %s = ? LIMIT 100", method)
	default:
		return nil, fmt.Errorf("unsupported search method: %s", method)
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Error preparing query:", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(searchInformation)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := new(models.User)
		if err := rows.Scan(&user.ID, &user.Username, &user.Gender, &user.Email, &user.Password, &user.CoverPic, &user.ProfilePic, &user.City, &user.WebSite); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error during rows iteration:", err)
		return nil, err
	}

	return users, nil
}

func UpdateUser(user models.User) (*models.User, *sql.DB, error) {
	userstr := "update users set"
	var params []interface{}
	if user.Username != "" {
		userstr = userstr + " ,name=?" // update user set name=?
		params = append(params, user.Username)
	}
	if user.City != "" {
		userstr = userstr + " ,city=?"
		params = append(params, user.City)
	}
	if user.WebSite != "" {
		userstr = userstr + " ,website=?"
		params = append(params, user.WebSite)
	}
	if user.ProfilePic != "" {
		userstr = userstr + " ,profilePic=?"
		params = append(params, user.ProfilePic)
	}
	if user.CoverPic != "" {
		userstr = userstr + " ,coverPic=?"
		params = append(params, user.CoverPic)
	}
	userstr = userstr + " where id=?" //  update user set name=? where username=?
	params = append(params, user.ID)
	userstr = strings.Replace(userstr, ",", "", 1)

	stmt, err := db.Prepare(userstr)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(params...)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err

	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	user.ID = int(id)
	return &user, db, err
}
