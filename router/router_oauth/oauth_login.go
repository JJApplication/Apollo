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
	"time"
)

func OAuthLogin(c *gin.Context) {
	code := c.Query(middleware.OAuthCode)
	token, err := oauth_manager.GetAccessToken(code)
	if err != nil {
		router.Response(c, "", false)
		return
	}
	user, err := oauth_manager.GetGithubUser(token)
	// 获取用户信息后返回前台 同时存储到manager表示登录成功
	if err != nil {
		router.Response(c, "", false)
		return
	}
	// store
	oauth_manager.AddOAuthUser(oauth_manager.OAuthUser{
		Username:  user.Login,
		Token:     token,
		HomeUrl:   user.HomeUrl,
		Avatar:    user.AvatarUrl,
		LoginTime: time.Now().Unix(),
	})
	// 作为oauth认证客户端响应, 返回要重定向的url
	router.Response(c, oauth_manager.ApolloOAuthUser{
		Login:       user.Login,
		AvatarUrl:   user.AvatarUrl,
		HomeUrl:     user.HomeUrl,
		AccessToken: token,
	}, true)
	return
}
