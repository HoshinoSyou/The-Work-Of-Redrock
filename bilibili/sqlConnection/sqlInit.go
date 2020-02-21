package sqlConnection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SqlInit() (err error) {
	datebase := "root:@tcp(127.0.0.1:3306)/bilibiliproject?charset=utf8"
	db, err = sql.Open("mysql", datebase)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return
}

func SqlConn() *sql.DB {
	err := SqlInit()
	if err != nil {
		fmt.Printf("数据库连接失败喵！, err:%v\n", err)
		return db
	}
	fmt.Println("数据库连接成功喵！")
	return db
}
