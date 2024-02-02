/*
Project: Apollo init.go
Created: 2021/12/22 by Landers
*/

package router_web

import (
	"github.com/JJApplication/Apollo/config"
	_ "github.com/JJApplication/Apollo/docs"
	"github.com/JJApplication/Apollo/middleware"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var routerWeb *gin.RouterGroup

func Init(r *gin.Engine) {
	routerWeb = r.Group("")
	{
		routerWeb.GET("/", middleware.MiddleCache(Index))
		routerWeb.GET("/favicon.ico", middleware.MiddleCache(Favicon))
		routerWeb.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		routerWeb.GET("heartbeat", func(c *gin.Context) {
			router.Response(c, "", true)
		})
		handleUIRouter()
	}
}

func handleUIRouter() {
	for _, r := range config.ApolloConf.Server.UIRouter {
		routerWeb.GET(r, middleware.MiddleCache(Index))
	}
}
