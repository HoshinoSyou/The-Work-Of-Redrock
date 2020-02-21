package dao

import (
	"bilibili/sqlConnection"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type myvideo struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type userData struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Videos   []gin.H
}

var myVideo myvideo
var myVideoSlice []myvideo
var mJSON gin.H
var mJSONs []gin.H
var dJSON gin.H
var d userData
var Data gin.H

func MyHomeDataQuery(m *gin.Context) bool {
	db := sqlConnection.SqlConn()
	username := m.Param("username")
	myHomeDataQuery := `select id,username,sex,age,address from userInformation where id = ?`
	stmt, err1 := db.Prepare(myHomeDataQuery)
	if err1 != nil {
		log.Printf("预处理失败喵！错误信息:%v\n", err1)
		return false
	}
	defer stmt.Close()
	rows, err2 := stmt.Query(username)
	if err2 != nil {
		log.Printf("查询个人信息失败喵TAT！错误信息:%v\n", err2)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&d.Id, &d.Username, &d.Sex, &d.Age, &d.Address)
		if err != nil {
			log.Printf("扫描个人详情页失败喵QAQ！错误信息:%v\n", err)
			return false
		}
	}
	MyVideo(m)
	d.Videos = mJSONs
	Data = MyHomeJSON()
	dJSON = nil
	return true
}

func MyVideo(m *gin.Context) bool {
	db := sqlConnection.SqlConn()
	username := m.Param("username")
	wordQuery := `select title,link from video where creator=?`
	stmt, err1 := db.Prepare(wordQuery)
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
		err := rows.Scan(&myVideo.Title, &myVideo.Link)
		if err != nil {
			log.Printf("扫描失败喵！错误信息:%v\n", err)
			return false
		} else {
			log.Printf("扫描成功喵！")
		}
		myVideoSlice = append(myVideoSlice, myVideo)
	}
	mJSONs = MyVideoJSON()
	myVideoSlice = nil
	return true
}

func MyVideoJSON() []gin.H {
	for _, myVideo := range myVideoSlice {
		mJSON = gin.H{
			"title": myVideo.Title,
			"link":  myVideo.Link}
		mJSONs = append(mJSONs, mJSON)
	}
	return mJSONs
}

func MyHomeJSON() gin.H {
	dJSON = gin.H{
		"id":       d.Id,
		"username": d.Username,
		"sex":      d.Sex,
		"age":      d.Age,
		"address":  d.Address,
		"video":    mJSONs}
	return dJSON
}
