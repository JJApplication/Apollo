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

func OAuthLogout(c *gin.Context) {
	oauth_manager.RemoveOAuthUser(c.Request.Header.Get(middleware.OAuthToken))
	router.Response(c, "", true)
	return
}
