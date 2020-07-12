package control

import (
	"github.com/gin-gonic/gin"
	"test/service"
)

func Intofight (c *gin.Context){//进房打架
	service.Random(c,"gobanger")
}

func Intowatch (c *gin.Context){//进房打架
	service.Random(c,"audience")
}