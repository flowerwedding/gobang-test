package control

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"test/dao"
	inits "test/model"
	"test/service"
)

func Register (c *gin.Context) {//登录成功
	var user inits.Player
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")

	dao.RegisterSearch(&user)
	if user.Id == 0 {
		service.ErrorJson(c,"密码错误")
		return
	}

	session:=sessions.Default(c)
	session.Set("username", user.Username)
	session.Set("room", 0)
	session.Set("statu", 0)
	session.Set("id", user.Id)
	_ = session.Save()

	service.OKJson(c,"登陆成功，欢迎回来！")
}
