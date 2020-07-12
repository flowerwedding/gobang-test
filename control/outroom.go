package control

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"test/dao"
	inits "test/model"
	"test/service"
)

func Outroom (c *gin.Context) {//退房成功
	session :=sessions.Default(c)
	session.Set("statu", 0)
	_ = session.Save()

	perid := c.PostForm("number")

	if perid == ""{
		perid = session.Get("room").(string)
	}
	var r inits.Room

	dao.OutroomDelete(&r,perid)

	service.OKJson(c,"退房成功")
}