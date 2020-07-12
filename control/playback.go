package control

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"test/dao"
	"test/service"
)

func Playback(c *gin.Context){
	preid := c.Query("number")
	id,_ := strconv.Atoi(preid)
	steps := dao.SelectHistory(id)
	var a = [16][16]int{{0}}

	for n := 1; n < len(steps);n++{
		for u,v := range steps[1:n] {
			i1 , j1 ,_ := service.Fenjie(v)
			a[i1 - 1][j1 -1] = 2 - (u + 1) % 2
		}
		s := steps[n]
		_ , _ , comma := service.Fenjie(s)
		service.StepXY(a,s[1:comma],s[comma + 1:len(s)-1],c)
	}
}