package dao

import (
	"bilibili/sqlConnection"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type user3 struct {
	id          int
	username    string
	oldpassword string
	newpassword string
}

var u user3

func OldPasswordQuery(username string, password string) bool {
	db := sqlConnection.SqlConn()
	sqlString := `select id,password from service where username =? and password =?`
	prepare, err1 := db.Prepare(sqlString)
	if err1 != nil {
		log.Printf("预处理失败喵！错误信息:%v\n", err1)
		return false
	}
	defer prepare.Close()
	rows, err2 := prepare.Query(username, password)
	if err2 != nil {
		log.Printf("查询失败喵！错误信息:%v\n", err2)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.id)
		if err != nil {
			log.Printf("扫描id失败喵！错误信息:%v\n", err)
			return false
		}
		if u.id >= 0 {
			return true
		} else {
			log.Printf("找不到用户信息喵！错误信息:%v\n", err)
		}
	}
	return false
}

func PasswordUpdate(up *gin.Context) bool {
	username := up.PostForm("username")
	newpassword := up.PostForm("newpassword")
	db := sqlConnection.SqlConn()
	sqlString := `update service set password =? where username =?`
	prepare, err1 := db.Prepare(sqlString)
	if err1 != nil {
		log.Printf("预处理失败喵！错误信息：%d\n", err1)
		return false
	}
	defer prepare.Close()
	result, err2 := prepare.Exec(newpassword, username)
	if err2 != nil {
		log.Printf("更新密码失败喵！错误信息:%d\n", err2)
		return false
	}
	i, _ := result.RowsAffected()
	log.Printf("修改的行数喵：%d\n", i)
	return true
}
