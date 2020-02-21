package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginMiddleWare() gin.HandlerFunc {
	return func(w *gin.Context) {
		if cookie,err :=w.Cookie("loginCookie") ; err == nil{
			if cookie == "68s49/3"{
				w.Next()
				return
			}
		}
		w.JSON(401,gin.H{
			"status":http.StatusUnauthorized,
			"error":"err",
			"message":"登录后才能查看喵！"})
		w.Abort()
		return
	}
}
