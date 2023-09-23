/*
   Create: 2023/9/22
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_auth

import (
	"github.com/JJApplication/Apollo/app/token_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

// Logout 登出仅针对登入token 不针对认证码
func Logout(c *gin.Context) {
	token_manager.DisActiveAllToken()
	c.SetCookie("token", "", config.ApolloConf.Server.AuthExpire, "", "", false, false)
	router.Response(c, true, true)
}
