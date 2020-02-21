package dao

import (
	"bilibili/sqlConnection"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type video1 struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Creator      string `json:"creator"`
	Introduction string `json:"introduction"`
}

var Video video1
var videoSlice []video1
var videoJSON gin.H
var VideoJSONs []gin.H

func VideoQuery(v *gin.Context) bool {
	db := sqlConnection.SqlConn()
	id := v.Param("id")
	sqlStr := `select id,title,creator,introduction from video where id=?`
	stmt, err1 := db.Prepare(sqlStr)
	if err1 != nil {
		log.Printf("预处理失败喵！错误信息:%v\n", err1)
		return false
	}
	defer stmt.Close()
	rows, err2 := stmt.Query(id)
	if err2 != nil {
		log.Printf("查询失败喵！错误信息:%v\n", err2)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		err3 := rows.Scan(&Video.Id, &Video.Title, &Video.Creator, &Video.Introduction)
		if err3 != nil {
			log.Printf("获取信息失败喵！错误信息:%v\n", err3)
			return false
		}
		videoSlice = append(videoSlice, Video)
	}
	VideoJSONs = VideoJSON()
	videoSlice = nil
	return true
}

func VideoJSON() []gin.H {
	for _, v := range videoSlice {
		videoJSON = gin.H{
			"id":           v.Id,
			"title":        v.Title,
			"creator":      v.Creator,
			"introduction": v.Introduction}
		VideoJSONs = append(VideoJSONs, videoJSON)
	}
	return VideoJSONs
}
