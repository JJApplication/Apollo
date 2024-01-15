/*
   Create: 2024/1/10
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package engine

import (
	"github.com/gin-gonic/gin"
)

// 本地APP互访中间件

const (
	XLocal = "X-Gateway-Local"
	YES    = "yes"
	No     = "no"
)

func MiddleWareXLocal() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.RemoteIP() != "127.0.0.1" && c.RemoteIP() != "localhost" {
			c.Abort()
			return
		}
		if c.Request.Header.Get(XLocal) != YES {
			c.Abort()
			return
		}
		c.Next()
	}
}
