package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"test/chat"
	"test/control"
)

func main(){
	router:=gin.Default()
	router.Use(cors.Default())

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("know-about-session",store))

	router.POST("/GoBang/login", control.Login)       //注册
	router.POST("/GoBang/register", control.Register) //登录

	router.POST("/GoBang/getroom",control.Getroom) //开房
	router.DELETE("/GoBang/outroom",control.Outroom) //退房

	router.POST("/GoBang/intofight",control.Intofight) //进房打架
	router.POST("/GoBang/intowatch",control.Intowatch) //进房观战

	router.GET("/GoBang/unbalance",control.Unbalance)//不禁手
	router.GET("/GoBang/balance",control.Balance) //禁手
	router.GET("/GoBang/TUnbalance",control.TUnbalance) //副本
	router.GET("/GoBang/TBalance",control.Tbalance)//副本

	router.GET("/GoBang/playback",control.Playback) //回放
	router.GET("/GoBang/other",control.Unbalance) //悔棋 1 、求和 2 、认输 3 ,多一个other参数

	router.GET("/GoBang/ws2", func(c *gin.Context) {
		chat.HandleConnections(2, c.Writer , c.Request)
	})
	router.GET("/GoBang/wsn", func(c *gin.Context) {
		chat.HandleConnections(0, c.Writer , c.Request)
	})

	go chat.HandleMessages()

	_ = router.Run(":8080")
}