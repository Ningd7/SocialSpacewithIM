package config

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func init() {
	dbSource := "root:123456@tcp(192.168.43.116:3306)/SocialSpace?charset=utf8mb4&parseTime=True&loc=Local\"
	dbcnn, err := sql.Open("mysql", dbSource)
	if err != nil {
		fmt.Println(err)
	}
	db = dbcnn
}

func GetDB() *sql.DB {
	return db
}
