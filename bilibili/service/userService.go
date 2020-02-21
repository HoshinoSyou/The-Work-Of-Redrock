package service

import (
	"bilibili/dao"
	"github.com/gin-gonic/gin"
)

func Register(i *gin.Context) bool {
	username := i.PostForm("username")
	registerQuery := dao.RegisterInformationQuery(username)
	registerInsert := dao.RegisterInformationInsert(i)
	if registerQuery && registerInsert {
		return true
	}
	return false
}

func Login(u *gin.Context) bool {
	username := u.PostForm("username")
	password := u.PostForm("password")
	loginInformation := dao.LoginInformationQuery(username, password)
	return loginInformation
}

func ChangePasswordQuery(up *gin.Context) bool {
	password := up.PostForm("oldpassword")
	username := up.PostForm("username")
	res := dao.OldPasswordQuery(username, password)
	return res
}

func UpdatePassword(up *gin.Context) bool {
	res := dao.PasswordUpdate(up)
	return res
}

func Logout(d *gin.Context) bool {
	username := d.PostForm("username")
	password := d.PostForm("password")
	res1 := dao.LogoutQuery(username, password)
	res2 := dao.DeleteUser(d)
	if res1 && res2 {
		return true
	} else {
		return false
	}
}

func Search(s *gin.Context) bool {
	res := dao.WordSearch(s)
	return res
}

func MyHome(m *gin.Context) bool {
	res := dao.MyHomeDataQuery(m)
	return res
}

func Contribute(f *gin.Context) bool {
	res1 := dao.FilesUpload(f)
	res2 := dao.VideoInformation(f)
	if res1 && res2 {
		return true
	}
	return false
}

func DeleteVideo(d *gin.Context) bool {
	res := dao.DeleteVideo(d)
	return res
}

func Home() bool {
	res := dao.RandomVideo()
	return res
}

func VideoPage(v *gin.Context) bool {
	res := dao.VideoQuery(v)
	return res
}
