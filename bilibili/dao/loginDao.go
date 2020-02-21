package dao

import (
	"bilibili/sqlConnection"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type user2 struct {
	username string
	password string
	id       int
}

func LoginInformationQuery(username string, password string) bool {
	db := sqlConnection.SqlConn()
	sqlQuery := `select id from user where username=? and password=?;`
	stmt, err1 := db.Prepare(sqlQuery)
	if err1 != nil {
		log.Printf("预处理失败喵！错误信息:%v\n", err1)
		return false
	}
	defer stmt.Close()
	row, err2 := stmt.Query(username, password)
	if err2 != nil {
		log.Printf("查询失败喵！错误信息:%v\n", err2)
		return false
	}
	defer row.Close()
	for row.Next() {
		var u user2
		err := row.Scan(&u.id)
		if err != nil {
			log.Printf("扫描id失败喵！错误信息:%v\n",err)
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
