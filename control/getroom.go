package control

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/dao"
	inits "test/model"
	"test/service"
	"time"
)

func Getroom (c *gin.Context) {//开房成功
	var room inits.Room
	perid := c.DefaultPostForm("number"," ")
	room.Password = c.DefaultPostForm("password"," ")
	if perid == " " {
		perid = strconv.FormatInt( time.Now().Unix(), 10)
	}
	room.Id,_ = strconv.Atoi(perid)

	dao.GetroomSearch(&room)
	if room.Member1 != ""{
		service.ErrorJson(c,"房间已占用")
		return
	}

	session :=sessions.Default(c)
	room.Member1 = session.Get("username").(string)
	room.Member2 = " "
	room.Statu = 0
	room.Balance = c.DefaultPostForm("balance","unbalance")
	session.Set("statu", 1)
	_ = session.Save()

	err := dao.GetroomInsert(&room)
	if err != nil {
		service.OKJson(c,err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": "申请成功!"})
}