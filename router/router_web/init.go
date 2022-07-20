/*
Project: Apollo init.go
Created: 2021/12/22 by Landers
*/

package router_web

import (
	"github.com/gin-gonic/gin"
	_ "github.com/JJApplication/Apollo/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var routerWeb *gin.RouterGroup

func Init(r *gin.Engine) {
	routerWeb = r.Group("")
	{
		routerWeb.GET("/", Index)
		routerWeb.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
