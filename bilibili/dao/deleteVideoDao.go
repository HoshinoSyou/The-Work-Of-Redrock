package dao

import (
	"bilibili/sqlConnection"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func DeleteVideo(d *gin.Context) bool {
	db := sqlConnection.SqlConn()
	id := d.Param("id")
	sqlStr := `delete from video where id =?`
	stmt, err1 := db.Prepare(sqlStr)
	if err1 != nil {
		log.Printf("预处理失败喵！错误信息:%v\n", err1)
		return false
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(id)
	if err2 != nil {
		log.Printf("删除数据失败喵！错误信息:%v\n", err2)
		return false
	}
	log.Printf("删除id为%v行的数据成功", id)
	file := "./files/" + id + ".mp4"
	err3 := os.Remove(file)
	if err3 != nil {
		log.Printf("删除id为%v的文件失败喵！，记得删除喵！", id)
	}
	return true
}
