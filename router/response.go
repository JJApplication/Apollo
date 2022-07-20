/*
Project: Apollo response.go
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

func Response(c *gin.Context, data interface{}, isOK bool) {
	// 设置上下文
	var status string
	if isOK {
		status = OK
	} else {
		status = BAD
	}

	c.Set("data", data)
	c.JSON(http.StatusOK, gin.H{
		"data":   data,
		"status": status,
		"time":   time.Now().Unix(),
	})
	return
}

func Res4xx(c *gin.Context, data interface{}) {
	c.Set("data", data)
	c.JSON(http.StatusBadRequest, gin.H{
		"data":  data,
		"error": "bad request",
		"time":  time.Now().Unix(),
	})
	return
}

func Res5xx(c *gin.Context, data interface{}) {
	c.Set("data", data)
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":  data,
		"error": "inner error",
		"time":  time.Now().Unix(),
	})
	return
}

// todo 限流时的429状态返回
