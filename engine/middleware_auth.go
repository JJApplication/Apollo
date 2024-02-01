/*
Create: 2022/7/25
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package engine
package engine

import (
	"github.com/JJApplication/Apollo/app/oauth_manager"
	"github.com/JJApplication/Apollo/app/token_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	ApolloAuthCode = "ApolloAuthCode"
	OAuthToken     = "access-token"
)

// MiddleWareAuth 前置用户校验
// 只有部分API需要校验 需要校验的接口组，单独使用此中间件
// 优先级认证码 > 传入的headers > cookie
//
// OAuth检查成功后跳过内部认证
func MiddleWareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get(OAuthToken)
		if user := oauth_manager.ValidateOAuth(accessToken); user.Username != "" {
			if oauth_manager.ValidateOAuthAdmin(user.Username) {
				c.Next()
				return
			}
			if isAllowAPI(c.Request.URL.Path) {
				c.Next()
				return
			}
		}

		// 内部认证
		authCode := c.Query("auth")
		if authCode != "" {
			// 存在认证码 但是认证码校验失败时不会继续校验
			if !isAuthVerify(authCode) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			// 存在认证码 校验成功直接返回
			c.Next()
			return
		}

		// 从headers中取token校验
		headerToken := c.GetHeader("token")
		validate := token_manager.ValidateToken(utils.GetRemoteIP(c.Request), headerToken)
		if headerToken != "" && validate {
			// header校验成功直接返回
			c.Next()
			return
		}

		// 从cookie中去token校验
		cookieToken, err := c.Cookie("token")
		if err != nil || cookieToken == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token_manager.ValidateToken(utils.GetRemoteIP(c.Request), cookieToken) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

// 计算authcode是否有效
// authcode 形如code_time time为创建的时间 默认失效日期为60s
// 生效时间可以配置
// 优先级环境变量 > 配置文件
func isAuthVerify(code string) bool {
	envCode := os.Getenv(ApolloAuthCode)

	if envCode == "" {
		envCode = config.ApolloConf.Server.AuthCode
	}

	return code != "" && code == envCode
}

func isAllowAPI(url string) bool {
	if len(config.ApolloConf.Server.OAuth.AllowApiList) <= 0 {
		return false
	}
	for _, api := range config.ApolloConf.Server.OAuth.AllowApiList {
		if url == api {
			return true
		}
	}

	return false
}
