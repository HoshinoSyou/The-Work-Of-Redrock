package enter

import (
	"bilibili/dao"
	"bilibili/middleware"
	"bilibili/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"strconv"
)

func Entrance() {
	r := gin.Default()
	r.POST("/register/:username/:password/：oncepassword", func(r *gin.Context) {
		username := r.PostForm("username")
		password := r.PostForm("password")
		oncepassword := r.PostForm("oncepassword")
		if username == "" {
			r.JSON(200, gin.H{
				"status":  http.StatusOK,
				"message": "用户名不能为空喵！"})
		}
		if username != "" {
			if password != oncepassword {
				r.JSON(200, gin.H{
					"status":  http.StatusOK,
					"message": "两次输入的密码不一致喵！"})
			} else {
				row := service.Register(r)
				if row {
					r.JSON(203, gin.H{
						"status":      http.StatusNonAuthoritativeInfo,
						"message":     "注册失败喵！",
						"possibility": "1.用户名已存在，重新输入一个用户名吧喵\n2.系统崩溃辽，请和我们的客服亲联系喵"})
				} else {
					r.JSON(200, gin.H{
						"status":  http.StatusOK,
						"message": "恭喜你~注册成功喵！"})
				}
			}
		}
	}) //注册

	r.GET("/login/:username/:password", func(l *gin.Context) {
		username := l.PostForm("username")
		if username == "" {
			l.JSON(200, gin.H{
				"status":  http.StatusOK,
				"message": "用户名不能为空喵！"})
		}
		if username != "" {
			row := service.Login(l)
			if row {
				l.JSON(200, gin.H{
					"status":  http.StatusOK,
					"message": "欢迎回来" + username + "喵！"})
				l.SetCookie("loginCookie", "68s49/3", 1296000,
					"/login", "localhost",
					false, false)
			} else {
				l.JSON(203, gin.H{
					"status":  http.StatusNonAuthoritativeInfo,
					"message": "用户名或密码错误喵！"})
			}
		}
	}) //登录

	r.PUT("/changepassword/:username/:oldpassword/:newpassword/:oncenew", middleware.LoginMiddleWare(), func(c *gin.Context) {
		oldpassword := c.Param("oldpassword")
		username := c.Param("username")
		if username == "" {
			c.JSON(200, gin.H{
				"status":  http.StatusOK,
				"message": "用户名不能为空喵！"})
		} else {
			if oldpassword == "" {
				c.JSON(200, gin.H{
					"status":  http.StatusOK,
					"message": "旧密码不能为空喵！"})
			} else {
				res := service.ChangePasswordQuery(c)
				if res {
					newpassword := c.Param("newpassword")
					if newpassword == "" {
						c.JSON(200, gin.H{
							"status":  http.StatusOK,
							"message": "新密码不能为空喵！"})
					} else {
						oncenew := c.Param("oncenew")
						if oncenew == newpassword {
							res := dao.PasswordUpdate(c)
							if res {
								c.JSON(200, gin.H{
									"status":  http.StatusOK,
									"message": "修改密码成功喵！要好好记住你的新密码喵"})
							} else {
								c.JSON(403, gin.H{
									"status":  http.StatusForbidden,
									"message": "修改密码失败喵！请联系我们的管理员亲喵！"})
							}
						} else {
							c.JSON(200, gin.H{
								"status":  http.StatusForbidden,
								"message": "两次输入的密码不一致喵！"})
						}
					}
				} else {
					c.JSON(203, gin.H{
						"status":  http.StatusNonAuthoritativeInfo,
						"message": "用户名或密码不正确喵！"})
				}
			}
		}
	}) //修改密码

	r.DELETE("/logout/:username/:password", middleware.LoginMiddleWare(), func(d *gin.Context) {
		username := d.PostForm("username")
		password := d.PostForm("password")
		if username != "" && password != "" {
			res := service.Logout(d)
			if res {
				d.JSON(200, gin.H{
					"status":  http.StatusOK,
					"message": "注销账号成功喵！不要忘记我们喵(＞﹏＜)"})
			} else {
				d.JSON(203, gin.H{
					"status":  http.StatusNonAuthoritativeInfo,
					"message": "注销账号失败喵！可能因为：\n1.密码或用户名错误喵\n2.系统被不明力量干扰了喵！一定要和我们客服亲反馈喵！"})
			}
		} else {
			d.JSON(200, gin.H{
				"status":  http.StatusOK,
				"message": "用户名和密码不能为空喵！"})
		}
	}) //注销账户

	r.GET("/search/:word", func(s *gin.Context) {
		word := s.Param("word")
		if word == "" {
			s.JSON(200, gin.H{
				"status":  http.StatusOK,
				"message": "搜索内容不能为空喵！"})
		}
		if word != "" {
			if dao.SVideoSlice == nil {
				s.JSON(200, gin.H{
					"status":  http.StatusOK,
					"word":    word,
					"message": "库存里面找不到了喵！(＞﹏＜)"})
			} else {
				res := service.Search(s)
				if res == false {
					s.JSON(403, gin.H{
						"status":  http.StatusForbidden,
						"message": "系统被不明力量干扰了喵！"})
				} else {
					s.JSON(200, gin.H{
						"status":  http.StatusOK,
						"message": "以下是搜索" + word + "相关内容喵！",
						"result":  dao.SVideos})
					dao.SVideos = nil
					dao.SVideoJSONs = nil
					dao.SVideoSlice = nil
				}
			}
		}
	}) //搜索

	r.GET("/myhome", middleware.LoginMiddleWare(), func(m *gin.Context) {
		res := service.MyHome(m)
		if res {
			m.JSON(200, gin.H{
				"status": http.StatusOK,
				"data":   dao.Data})
			dao.Data = nil
		} else {
			m.JSON(403, gin.H{
				"status":  http.StatusForbidden,
				"message": "页面丢失喵！"})
		}
	}) //个人主页

	r.PUT("/contribute", middleware.LoginMiddleWare(), func(c *gin.Context) {
		res := service.Contribute(c)
		if res {
			c.JSON(200, gin.H{
				"status":  http.StatusOK,
				"message": "投稿成功喵！"})
		} else {
			c.JSON(403, gin.H{
				"status":  http.StatusForbidden,
				"message": "投稿失败喵！"})
		}
	}) //投稿

	r.DELETE("/home/deleteVideo/:id", middleware.LoginMiddleWare(), func(d *gin.Context) {
		res := service.DeleteVideo(d)
		if res {
			d.JSON(200, gin.H{
				"status":  http.StatusOK,
				"message": "删除该视频成功喵！"})
		} else {
			d.JSON(200, gin.H{
				"status":  http.StatusOK,
				"message": "删除该视频失败喵！"})
		}
	}) //删除投稿

	r.GET("/home", func(h *gin.Context) {
		res := service.Home()
		if res {
			h.JSON(200, gin.H{
				"status":  http.StatusOK,
				"message": "推荐",
				"data":    dao.HVideoJSONs})
			dao.HVideoJSONs = nil
		} else {
			h.JSON(403, gin.H{
				"status":  http.StatusForbidden,
				"message": "系统崩溃辽，哭唧唧"})
		}
	}) //首页推荐

	r.GET("/video/=id", func(v *gin.Context) {
		res := service.VideoPage(v)
		if res {
			path := "./files/" + strconv.Itoa(dao.Video.Id) + ".mp4"
			file, err := os.Open(path)
			if err != nil {
				v.JSON(403, gin.H{
					"status":  http.StatusForbidden,
					"message": "找不到页面喵！"})
			} else {
				v.JSON(200, gin.H{
					"status": http.StatusOK,
					"data":   dao.VideoJSONs})
			}
			defer file.Close()
		} else {
			v.JSON(403, gin.H{
				"status":  http.StatusForbidden,
				"message": "找不到页面喵！"})
		}
	}) //视频详情页

	r.Run(":17736")
}
