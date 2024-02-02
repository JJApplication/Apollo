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

func GetGithubOAuthApi(c *gin.Context) {
	router.Response(c, oauth_manager.GetGithubOAuthUrl(), true)
	return
}
