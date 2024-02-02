/*
Create: 2022/7/28
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_app
package router_app

import (
	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func StopApp(c *gin.Context) {
	app := c.Query("app")
	if app == "" {
		router.Response(c, "", false)
	}

	isOk, err := app_manager.Stop(app)
	if err != nil || !isOk {
		router.Response(c, err, false)
		return
	}

	router.Response(c, err, true)
}

func StopAll(c *gin.Context) {
	stopList, err := app_manager.StopAll()
	if err != nil {
		router.Response(c, err, false)
		return
	}

	router.Response(c, stopList, true)
}
