package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"test/dao"
	inits "test/model"
	"test/service"
	"time"
)

func TUnbalance(c *gin.Context){
	//	timer1 := time.NewTimer(time.Second * 2)
	//	<-timer1.C
	var s string
	var flag = 0
	ticket := time.NewTicker(time.Millisecond * 10)//限时三分钟，十秒更新一次
	n := 0
	for t := range ticket.C{
		s = c.Query("stepXY")
		if s != "" { flag = 1 }
		fmt.Println(s)
		fmt.Println(t)
		if flag == 1 || n >= 18000 {
			ticket.Stop()
			break
		}
		n++
	}
	id := 4

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
		a[i - 1][j - 1] = 2

		service.StepXY(a,s[1:comma],s[comma + 1:len(s)-1],c)

		fmt.Println(s)
		err1 := dao.History(id, s)
		if err1 != nil {
			service.ErrorJson(c,err1.Error())
		}

		if service.Checkwin(a,i - 1,j - 1) {
			service.OKJson(c,"你赢了")
			//outroom()退赛
			Outroom(c)
		}else{
			service.OKJson(c,"/")
		}
	}else{
		service.ErrorJson(c,"此坐标已落子")
	}
}