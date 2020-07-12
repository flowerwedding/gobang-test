package service

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"test/dao"
	inits "test/model"
)

func Random(c *gin.Context,identify string){
	session :=sessions.Default(c)
	prestatu := session.Get("statu").(int)
	if prestatu != 0{
		ErrorJson(c,"您已进入其他房间")
		return
	}
	session.Set("statu", 2)

	number := c.PostForm("number")
	var r inits.Room
	dao.RandomSelect(number ,&r)

	if (number != "")||(r.Balance != "") {
		fmt.Println("aaaa")
		session.Set("room", number)
		_ = session.Save()
		if r.Password != " " && r.Password !=c.PostForm("password"){
			ErrorJson(c,"房间密码错误")
			return
		}
	}else{
		balance := c.DefaultPostForm("balance","unbalance")
		dao.RandomSelect2(balance,&r)

		premember2 := session.Get("username")
		session.Set("room", r.Id)
		_ = session.Save()
		if identify == "gobanger"{
			err1,err2 := dao.RandomUpdate(&r,premember2)
			if err1 != nil || err2 != nil{
				ErrorJson(c,err1.Error() + err2.Error())
				return
			}
		}
		err3 := dao.RandomUpdate2(premember2,identify)
		if err3 != nil {
			ErrorJson(c,err3.Error())
			return
		}
	}

	var verb string
	if identify == "gobanger"{ verb = "fight"} else { verb = "watch"}
	OtherOkJson(c,"进房成功","number",verb)
	return
}