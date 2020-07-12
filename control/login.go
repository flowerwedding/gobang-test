package control

import (
	"github.com/gin-gonic/gin"
	"test/dao"
	inits "test/model"
	"test/service"
)

func Login (c *gin.Context){//注册成功
	var user inits.Player
	user.Username=c.PostForm("username")

	dao.LoginSearch(&user)
	if user.Password != ""{
		service.ErrorJson(c,"用户已注册")
		return
	}

	user = inits.Player{
		Password : c.PostForm("password"),
		Username : c.PostForm("username"),
	}

	err := dao.LoginInsert(&user)
	if err != nil {
		service.ErrorJson(c,err.Error())
		return
	}

	service.OtherOkJson(c,"注册成功","your id",user.Id)
}