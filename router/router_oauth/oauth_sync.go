/*
   Create: 2024/2/2
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_oauth

import (
	"github.com/JJApplication/Apollo/app/oauth_manager"
	"github.com/JJApplication/Apollo/middleware"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

// OAuthLoginSync 传入access-token
//
// 通过token获取存储在缓存中的用户
// 使用用户数据与github同步，如果登录成功则认证有效，刷新缓存中的token有效期
func OAuthLoginSync(c *gin.Context) {
	token := c.Request.Header.Get(middleware.OAuthToken)
	if token == "" {
		router.Response(c, "", false)
		return
	}
	user := oauth_manager.ValidateOAuth(token)
	if expire := oauth_manager.SyncFromGithub(user.Username); expire {
		router.Response(c, "", false)
		return
	}
	router.Response(c, "", true)
	return
}
