/*
   Create: 2024/1/28
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_oauth

import (
	"github.com/JJApplication/Apollo/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerOAuth := r.Group("/api/oauth")
	{
		// 获取认证github oauth的地址
		routerOAuth.GET("/github", GetGithubOAuthApi)

		// github 第一步oauth的回调
		// github会在query中传递code=$authorized_code
		// 再callback中请求github获取access_token
		routerOAuth.GET("/login", OAuthLogin)

		// 根据当前登录的用户名 根据access token查询是否仍存在登录状态
		// 不存在登录状态时 重定向到登录页面并退出当前用户
		routerOAuth.GET("/check", OAuthCheckLogin)

		// 同github oauth同步状态
		routerOAuth.POST("/sync", middleware.MiddleWareAuth(), OAuthLoginSync)

		// 退出登录
		routerOAuth.POST("/logout", middleware.MiddleWareAuth(), OAuthLogout)
	}
}
