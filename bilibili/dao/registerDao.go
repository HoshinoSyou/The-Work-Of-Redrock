package dao

import (
	"bilibili/sqlConnection"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type user1 struct {
	id       int
	username string
	password string
}

func RegisterInformationQuery(username string) bool {
	db := sqlConnection.SqlConn()
	sqlQuery := `select id from user where username=?;`

	stmt, err1 := db.Prepare(sqlQuery)
	if err1 != nil {
		log.Printf("预处理失败喵！错误信息:%v\n", err1)
		return false
	}
	defer stmt.Close()
	rows, err2 := stmt.Query(username)
	if err2 != nil {
		log.Printf("查询失败喵！错误信息:%v\n", err2)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		var u user1
		err := rows.Scan(&u.id)
		if err != nil {
			log.Printf("扫描id失败喵！错误信息:%v\n", err)
			return false
		}
		if u.id >= 0 {
			return true
		}
		if u.id < 0 {
			return true
		}
	}
	return false
}

func RegisterInformationInsert(r *gin.Context) bool {
	db := sqlConnection.SqlConn()
	sqlInformationInsert := `insert into user (username, password) values (?, ?)`
	username := r.PostForm("username")
	password := r.PostForm("password")
	stmt, err1 := db.Prepare(sqlInformationInsert)
	if err1 != nil {
		log.Printf("预处理失败喵！错误信息:%v\n", err1)
		return false
	}
	defer stmt.Close()
	insert, err2 := stmt.Exec(username, password)
	if err2 != nil {
		log.Printf("注册失败喵！错误信息:%v\n", err2)
		return false
	}
	id, err3 := insert.LastInsertId()
	if err3 != nil {
		log.Printf("获取ID失败喵！错误信息:%v\n", err3)
	} else {
		log.Printf("注册的ID是:%v\n", id)
	}
	return true
}
