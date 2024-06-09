/*
Create: 2022/8/20
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_noengine
package router_noengine

import (
	"github.com/JJApplication/Apollo/app/noengine_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func GetAllNoEngineApp(c *gin.Context) {
	router.Response(c, noengine_manager.GetAllNoEngineAPPsRt(), true)
}

func GetNoEngineApp(c *gin.Context) {
	app := c.Query("app")
	if app == "" {
		router.Response(c, "", false)
		return
	}
	router.Response(c, noengine_manager.GetNoEngineAPPRt(app), true)
}

func GetNoEngineStatus(c *gin.Context) {
	app := c.Query("app")
	router.Response(c, noengine_manager.StatusNoEngineApp(app), true)
}

func StartNoEngineApp(c *gin.Context) {
	app := c.Query("app")
	err := noengine_manager.StartNoEngineApp(app)
	if err != nil {
		router.Response(c, "", false)
		return
	}
	router.Response(c, "", true)
}

func StopNoEngineApp(c *gin.Context) {
	app := c.Query("app")
	err := noengine_manager.StopNoEngineApp(app)
	if err != nil {
		router.Response(c, "", false)
		return
	}
	router.Response(c, "", true)
}

// RestartNoEngineApp remove and start
func RestartNoEngineApp(c *gin.Context) {
	app := c.Query("app")
	err := noengine_manager.StopNoEngineApp(app)
	if err != nil {
		router.Response(c, "", false)
		return
	}
	err = noengine_manager.RemoveNoEngineApp(app)
	if err != nil {
		router.Response(c, "", false)
		return
	}
	err = noengine_manager.StartNoEngineApp(app)
	if err != nil {
		router.Response(c, "", false)
		return
	}
	router.Response(c, "", true)
}

func PauseNoEngineApp(c *gin.Context) {
	app := c.Query("app")
	err := noengine_manager.PauseNoEngineApp(app)
	if err != nil {
		router.Response(c, "", false)
		return
	}
	router.Response(c, "", true)
}

func ResumeNoEngineApp(c *gin.Context) {
	app := c.Query("app")
	err := noengine_manager.ResumeNoEngineApp(app)
	if err != nil {
		router.Response(c, "", false)
		return
	}
	router.Response(c, "", true)
}

func RemoveNoEngineApp(c *gin.Context) {
	app := c.Query("app")
	err := noengine_manager.RemoveNoEngineApp(app)
	if err != nil {
		router.Response(c, "", false)
		return
	}
	router.Response(c, "", true)
}

func RefreshNoEngineApp(c *gin.Context) {
	_, err := noengine_manager.RefreshNoEngineMap()
	if err != nil {
		router.Response(c, "", false)
		return
	}
	router.Response(c, "", true)
}
