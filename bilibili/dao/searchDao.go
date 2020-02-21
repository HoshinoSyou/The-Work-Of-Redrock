package dao

import (
	"bilibili/sqlConnection"
	"github.com/gin-gonic/gin"
	"log"
)

type video2 struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Creator string `json:"creator"`
	Link    string `json:"link"`
}

var sVideo video2
var SVideoSlice []video2
var sVideoJSON gin.H
var SVideoJSONs []gin.H
var SVideos []gin.H

func WordSearch(s *gin.Context) bool {
	db := sqlConnection.SqlConn()
	word := s.Param("word")
	wordQuery := `select id,title,creator,link from video where title like ? or introduction like ?`
	stmt, err1 := db.Prepare(wordQuery)
	if err1 != nil {
		log.Printf("预处理失败喵！错误信息:%v\n", err1)
		return false
	}
	defer stmt.Close()
	rows, err2 := stmt.Query("%"+word+"%", "%"+word+"%")
	if err2 != nil {
		log.Printf("查询失败喵！错误信息:%v\n", err2)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&sVideo.Id, &sVideo.Title, &sVideo.Creator, &sVideo.Link)
		if err != nil {
			log.Printf("扫描失败喵！错误信息:%v\n", err)
			return false
		} else {
			log.Printf("扫描成功喵！")
			SVideoSlice = append(SVideoSlice, sVideo)
		}
	}
	SVideos = ResultJSON()
	return true
}

func ResultJSON() []gin.H {
	for _, v := range SVideoSlice {
		sVideoJSON = gin.H{
			"id":      v.Id,
			"title":   v.Title,
			"creator": v.Creator,
			"link":    v.Link}
		SVideoJSONs = append(SVideoJSONs, sVideoJSON)
	}
	return SVideoJSONs
}
