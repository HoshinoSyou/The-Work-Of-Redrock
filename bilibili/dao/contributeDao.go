package dao

import (
	"bilibili/sqlConnection"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
)

var idSlice0 = []int{}
var maxId0 int

func FilesUpload(f *gin.Context) bool {
	file, err1 := f.FormFile("视频")
	if err1 != nil {
		log.Printf("读取视频失败辽喵！错误信息:%v\n", err1)
		return false
	}
	log.Println(file.Filename)
	catalogue := fmt.Sprintf("./files/%s", file.Filename)
	err2 := f.SaveUploadedFile(file, catalogue)
	if err2 != nil {
		log.Printf("上传文件失败喵！错误信息:%v\n", err2)
		return false
	}
	id := f.Param("id")
	err := os.Rename(file.Filename, "./files/"+id+".mp4")
	if err != nil {
		fmt.Printf("上传文件名为%s喵，记得更改名字为%s喵！", file.Filename, id)
	}
	return true
}

func VideoInformation(c *gin.Context) bool {
	db := sqlConnection.SqlConn()
	name := c.PostForm("name")
	creator := c.PostForm("username")
	information := c.PostForm("information")
	res := MaxId()
	if res {
		link := "/" + strconv.Itoa(maxId0+1)
		sqlStr := `Insert into video (name,creator,information,link) values(?,?,?)`
		stmt, err := db.Prepare(sqlStr)
		if err != nil {
			log.Printf("预处理失败喵！错误信息:%v\n", err)
			return false
		}
		defer stmt.Close()
		exec, err2 := stmt.Exec(name, creator, information, link)
		if err2 != nil {
			log.Printf("上传信息失败喵！错误信息:%v\n", err2)
			return false
		}
		id, err3 := exec.LastInsertId()
		if err3 != nil {
			log.Printf("获取ID失败喵！错误信息:%v\n", err3)
			return false
		}
		fmt.Println(id)
		return true
	} else {
		return false
	}
}

func MaxId() bool {
	db := sqlConnection.SqlConn()
	sqlStr := `select * from video`
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("查询数据库失败喵！错误信息:%v\n", err)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(idSlice0)
		if err != nil {
			log.Printf("扫描id失败喵！错误信息:%v\n", err)
			return false
		}
	}
	maxValue0 := idSlice0[0]
	maxIndex0 := 0
	for i := 1; i < len(idSlice0); i++ {
		if maxValue0 < idSlice0[i] {
			maxValue0 = idSlice0[i]
			maxIndex0 = i
		}
	}
	maxId0 = maxValue0
	log.Printf("最大id为%v,角标为%v", maxValue0, maxIndex0)
	return true
}
