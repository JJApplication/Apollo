/*
Project: dirichlet init.go
Created: 2021/12/22 by Landers
*/

package router_web

import (
	"github.com/gin-gonic/gin"
)

var routerWeb *gin.RouterGroup

func Init(r *gin.Engine) {
	routerWeb = r.Group("")
	{
		routerWeb.GET("/", Index)
	}
}