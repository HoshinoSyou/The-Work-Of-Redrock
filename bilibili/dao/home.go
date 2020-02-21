package dao

import (
	"bilibili/sqlConnection"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"time"
)

type video3 struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Creator string `json:"creator"`
	Link    string `json:"link"`
}

var id int
var maxId int
var hVideo video3
var hVideoSlice []video3
var idSlice = []int{}
var idArray = [4]int{}
var hVideoJSON gin.H
var HVideoJSONs []gin.H

func RandomNumber() {
	rand.Seed(time.Now().UnixNano())
	for j := 1; j <= 4; j++ {
		id := rand.Intn(maxId-0+1) + 0
		idArray[j] = id
	}
	return
}

func GetMaxId() bool {
	db := sqlConnection.SqlConn()
	sqlStr := `select * from video`
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("查询数据库失败喵！错误信息:%v\n", err)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Printf("扫描id失败喵！错误信息:%v\n", err)
			return false
		}
		idSlice = append(idSlice, id)
	}
	maxValue := idSlice[0]
	maxIndex := 0
	for i := 1; i < len(idSlice); i++ {
		if maxValue < idSlice[i] {
			maxValue = idSlice[i]
			maxIndex = i
		}
	}
	maxId = maxValue
	log.Printf("最大id为%v,角标为%v", maxValue, maxIndex)
	return true
}

func RandomVideo() bool {
	res := GetMaxId()
	if res {
		RandomNumber()
		db := sqlConnection.SqlConn()
		sqlStr := `select id,title,creator,link where id =? `
		stmt, err := db.Prepare(sqlStr)
		if err != nil {
			log.Printf("预处理失败喵！错误信息:%v\n", err)
			return false
		}
		defer stmt.Close()
		for k := 0; k <= 3; k++ {
			rows, err := stmt.Query(idArray[k])
			if err != nil {
				log.Printf("查询数据库失败喵！错误信息:%v\n", err)
				return false
			}
			for rows.Next() {
				err := rows.Scan(&hVideo.Id, &hVideo.Title, &hVideo.Creator, &hVideo.Link)
				if err != nil {
					log.Printf("扫描数据失败喵！错误信息:%v\n", err)
					return false
				}
				log.Printf("扫描成功喵！")
				hVideoSlice = append(hVideoSlice, hVideo)
			}
		}
		HVideoJSONs = HomeVideoJSON()
		hVideoSlice = nil
		return true
	} else {
		return false
	}
}

func HomeVideoJSON() []gin.H {
	for _, v := range hVideoSlice {
		hVideoJSON = gin.H{
			"id":      v.Id,
			"title":   v.Title,
			"creator": v.Creator,
			"link":    v.Link}
		HVideoJSONs = append(HVideoJSONs, hVideoJSON)
	}
	return HVideoJSONs
}
