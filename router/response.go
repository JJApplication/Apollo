/*
Project: dirichlet response.go
Created: 2021/11/30 by Landers
*/

package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	OK  = "ok"  // 常规正确响应
	BAD = "bad" // 常规异常响应
)

func Response(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"time": time.Now().Unix(),
	})
}

func Res4xx(c *gin.Context, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"data":  data,
		"error": "bad request",
		"time":  time.Now().Unix(),
	})
}

func Res5xx(c *gin.Context, data interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":  data,
		"error": "inner error",
		"time":  time.Now().Unix(),
	})
}
