/*
Create: 2022/7/24
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_app
package router_app

import (
	"sort"
	"strings"

	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

type status struct {
	App    string
	Status string
}

func StatusApp(c *gin.Context) {
	app := c.Query("name")
	if app != "" {
		appStatus, err := app_manager.GetApp(app)
		if err != nil {
			router.Response(c, appStatus, false)
			return
		}
		router.Response(c, appStatus, true)
		return
	}

	apps, _ := app_manager.StatusAll()
	var stat []status
	for _, app := range apps {
		s := status{}
		s.App = strings.Trim(strings.Split(app, ":")[0], "[]")
		s.Status = strings.TrimSpace(strings.Split(app, ":")[1])
		stat = append(stat, s)
	}
	sort.SliceStable(stat, func(i, j int) bool {
		return stat[i].App < stat[j].App
	})

	router.Response(c, stat, true)
}
