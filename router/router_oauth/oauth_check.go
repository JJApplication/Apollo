/*
   Create: 2024/2/2
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_oauth

import (
	"github.com/JJApplication/Apollo/app/oauth_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

// OAuthCheckLogin 无需经过auth中间件认证 直接查询内部的oauth expire有效期
func OAuthCheckLogin(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		router.Response(c, "", false)
		return
	}
	if expire := oauth_manager.ExpireOAuthUser(user); expire {
		router.Response(c, "", false)
		return
	}
	router.Response(c, "", true)
	return
}
