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
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-gonic/gin"
)

type status struct {
	App    string
	Status string
}

type count struct {
	App  string
	Type string
}

func CountApp(c *gin.Context) {
	var res []count
	apps, err := app_manager.GetAllApp()
	if err != nil {
		router.Response(c, res, false)
		return
	}
	for _, app := range apps {
		res = append(res, count{
			App:  app.Meta.Name,
			Type: app.Meta.Type,
		})
	}
	router.Response(c, res, true)
	return
}

func StatusApp(c *gin.Context) {
	app := c.Query("name")
	if app != "" {
		appStatus, err := app_manager.Status(app)
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

func InfoApp(c *gin.Context) {
	app := c.Query("name")
	if app != "" {
		appInfo, err := app_manager.GetApp(app)
		if err != nil {
			router.Response(c, appInfo, false)
			return
		}
		router.Response(c, appInfo, true)
		return
	}
	router.Response(c, app_manager.App{}, false)
}

// FileTree 文件结构树
func FileTree(c *gin.Context) {
	var filesTree []utils.File
	filesTree = utils.GetFileTreeAllDepth(config.ApolloConf.APPRoot, config.ApolloConf.APPRoot)
	router.Response(c, filesTree, true)
}
