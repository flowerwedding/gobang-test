package control

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"test/dao"
	inits "test/model"
	"test/service"
)

func Balance(c *gin.Context){
	//s := c.PostForm("stepXY")

	s := c.Query("stepXY")
	session:=sessions.Default(c)
	ids := session.Get("room").(string)//?
	id, _ := strconv.Atoi(ids)
	//	id := 2
	prepares := session.Get("prepare").(string)
	prepare, _ := strconv.Atoi(prepares)
	if prepare == 1 {
		return
	}else{
		prepare = 1
	}

	var a = [16][16]int{{0}}
	var h inits.History
	dao.UnbalanceSelect(&h,id)
	arr := strings.Split(h.Step, " # ")
	for u,v := range arr[1:] {
		i1 , j1 ,_ := service.Fenjie(v)
		a[i1 - 1][j1 -1] = 2 - (u + 1) % 2
	}

	i , j , comma := service.Fenjie(s)

	if a[i- 1][j - 1] == 0 {
		statu := session.Get("statu").(string)
		a[i - 1][j - 1],_ = strconv.Atoi(statu)
		//	a[i - 1][j - 1] = 1

		if a[i - 1][j - 1] == 1{
			var result = 0
			for adjsame := 0; adjsame <= 3;adjsame ++{
				result += service.KeyPointForbCheck(a , i , j , adjsame , adjsame)
				result += service.KeyPointForbCheck(a , i , j , adjsame , adjsame + 4)
			}

			if(result != 0){
				service.ErrorJson(c,"禁手")
				return
			}
		}

		service.StepXY(a,s[1:comma],s[comma + 1:len(s)-1],c)

		err1 := dao.History(id, s)
		if err1 != nil {
			service.ErrorJson(c, err1.Error())
		}

		if service.Checkwin(a,i - 1,j - 1) {
			username := session.Get("username").(string)
			service.OtherOkJson(c,"你赢了","winner",username)
			//outroom()退赛
			Outroom(c)
		}else{
			service.OKJson(c,"/")
		}
	}else{
		service.ErrorJson(c,"此坐标已落子")
	}
}
