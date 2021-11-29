/*
Project: dirichlet start_app.go
Created: 2021/11/30 by Landers
*/

package router_app

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/dirichlet/app/app_manager"
	"github.com/landers1037/dirichlet/router"
)

// StartApp 启动app
func StartApp(r *gin.RouterGroup) {
	r.GET("", func(c *gin.Context) {
		err := app_manager.StartAll()
		router.Response(c, err)
	})
}
