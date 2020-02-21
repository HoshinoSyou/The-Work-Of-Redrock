package dao

import (
	"bilibili/sqlConnection"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type user struct {
	id       int
	username string
	password string
}

func LogoutQuery(username string, password string) bool {
	db := sqlConnection.SqlConn()
	sqlStr := `select id from user where username =?, password =?`
	stmt, err1 := db.Prepare(sqlStr)
	if err1 != nil {
		log.Printf("预处理失败喵！错误信息:%v\n", err1)
		return false
	}
	defer stmt.Close()
	rows, err2 := stmt.Query(username, password)
	if err2 != nil {
		log.Printf("查询数据库失败喵！错误信息:%v\n", err2)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id)
		if err != nil {
			log.Printf("获取数据失败喵！错误信息:%v\n", err)
			return false
		}
		if u.id >= 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func DeleteUser(d *gin.Context) bool {
	username := d.PostForm("username")
	password := d.PostForm("password")
	db := sqlConnection.SqlConn()
	deleteStr := `delete from user where username =?, password =?`
	stmt, err := db.Prepare(deleteStr)
	if err != nil {
		log.Printf("预处理失败喵！错误信息:%v\n", err)
		return false
	}
	defer stmt.Close()
	result, err1 := stmt.Exec(username, password)
	if err1 != nil {
		log.Printf("删除数据行失败喵！错误信息:%v\n", err1)
		return false
	}
	i, err2 := result.RowsAffected()
	if err2 != nil {
		log.Printf("查询影响行失败喵！错误信息:%v\n", err2)
	}
	log.Printf("影响行数为:%d\n", i)
	return true
}
