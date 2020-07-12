package service

import "github.com/gin-gonic/gin"

func StepXY(a [16][16]int,i string,j string,c *gin.Context){ //输出在postman
	size := 0.28//显示时的位置
	n := 1
	stepX := 0.04
	stepY := stepX // 一行的高度大概能竖着排四个`*`

	var pri [16] string
	// 坐标系 从上到下y是递减的，从左到右x是递增的
	message :="对方落子于坐标：( "+i+" , "+j+" ) ,请你落子！"
	pri[0] = " 1  2  3  4  5  6  7  8  9  10 11 12 13 14 15"
	for y := size + 0.02; y >= -size; y -= stepY {
		k := 0
		for x := -size; x <= size; x += stepX {
			if a[n - 1][k] == 1 {
				pri[n] += " & "
			}else if a[n - 1][k] == 2{
				pri[n] += " # "
			}else{
				pri[n] += " _ "
			}
			k++
		}
		n++
	}
	c.JSON(200,gin.H{"message":message,"  ":pri[0]," 1":pri[1]," 2":pri[2]," 3":pri[3]," 4":pri[4]," 5":pri[5]," 6":pri[6]," 7":pri[7]," 8":pri[8]," 9":pri[9],"10":pri[10],"11":pri[11],"12":pri[12],"13":pri[13],"14":pri[14],"15":pri[15]})
}