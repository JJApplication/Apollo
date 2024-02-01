/*
   Create: 2024/1/28
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_oauth

import (
	"github.com/JJApplication/Apollo/app/oauth_manager"
	"github.com/JJApplication/Apollo/engine"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
	"time"
)

func Init(r *gin.Engine) {
	routerOAuth := r.Group("/api/oauth")
	{
		// 获取认证github oauth的地址
		routerOAuth.GET("/github", func(c *gin.Context) {
			router.Response(c, oauth_manager.GetGithubOAuthUrl(), true)
		})

		// github 第一步oauth的回调
		// github会在query中传递code=$authorized_code
		// 再callback中请求github获取access_token
		routerOAuth.GET("/login", func(c *gin.Context) {
			code := c.Query("code")
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
		})

		// 根据当前登录的用户名 根据access token查询是否仍存在登录状态
		// 不存在登录状态时 重定向到登录页面并退出当前用户
		routerOAuth.GET("/check", func(c *gin.Context) {
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
		})

		// 同github oauth同步状态
		routerOAuth.POST("/sync", func(c *gin.Context) {
			user := c.Query("user")
			if user == "" {
				router.Response(c, "", false)
				return
			}
			if expire := oauth_manager.SyncFromGithub(user); expire {
				router.Response(c, "", false)
				return
			}
			router.Response(c, "", true)
			return
		})

		// 退出登录
		routerOAuth.POST("/logout", engine.MiddleWareAuth(), func(c *gin.Context) {
			oauth_manager.RemoveOAuthUser(c.Request.Header.Get("access-token"))
			router.Response(c, "", true)
			return
		})
	}
}
