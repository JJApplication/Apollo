/*
Create: 2022/7/25
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package engine
package engine

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	ApolloAuthCode = "ApolloAuthCode"
)

// MiddleWareAuth 前置用户校验
// 只有部分API需要校验 需要校验的接口组，单独使用此中间件
func MiddleWareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authCode := c.Query("auth")
		if authCode == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !isAuthVerify(authCode) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

// 计算authcode是否有效
// authcode 形如code_time time为创建的时间 默认失效日期为60s
// 生效时间可以配置
func isAuthVerify(code string) bool {
	envCode := os.Getenv(ApolloAuthCode)

	return code != "" && code == envCode
}
