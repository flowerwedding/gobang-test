package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorJson(c *gin.Context,message interface{}){
	c.JSON(http.StatusOK, gin.H{"code": 10001, "message": message})
}

func OKJson(c *gin.Context,message interface{}){
	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": message})
}

func OtherOkJson(c *gin.Context,message string,word string,other interface{}){
	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": message,word:other})
}