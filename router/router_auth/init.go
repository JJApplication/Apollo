/*
   Create: 2023/9/22
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_auth

import (
	"github.com/JJApplication/Apollo/engine"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerAuth := r.Group("/api/auth")
	{
		routerAuth.GET("/history", History)
		routerAuth.GET("/current", Current)
		routerAuth.POST("/login", Login)

		routerAuth.POST("/check", engine.MiddleWareAuth(), Check)
		routerAuth.POST("/logout", engine.MiddleWareAuth(), Logout)
	}
}
