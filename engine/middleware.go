/*
Project: dirichlet middleware.go
Created: 2021/11/26 by Landers
*/

package engine

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/landers1037/dirichlet/logger"
	"github.com/landers1037/dirichlet/utils"
)

// 基于本地配置加载中间件
// 优先被初始化以让所有中间件判断自己是否被激活
// 两种模式：全局global 路由route 在路由模式下需要指定路由

const (
	MiddleWare = "[MiddleWare]"
	PluginsPath = "lib"
)

var PreInjectMiddle []MiddleWareConfig
var DefaultMiddleWare = []MiddleWareConfig{
	{
		Name:   "log",
		Mode:   GlobalMode,
		Active: true,
	},
	{
		Name:   "recovery",
		Mode:   GlobalMode,
		Active: true,
	},
	{
		Name:   "plnack",
		Mode:   GlobalMode,
		Active: false,
	},
	{
		Name:   "cors",
		Mode:   GlobalMode,
		Active: true,
	},
}
var MiddleWareMap = map[string]gin.HandlerFunc{
	"log":      gin.Logger(),
	"recovery": gin.Recovery(),
	"cors":     MiddleWareCors(),
	"plnack":   MiddlewarePlnack(),
}

// 中间件加载
func loadMiddleWare(g *gin.Engine) {
	PreInjectMiddle = LoadMiddleWareConfig()
	if len(PreInjectMiddle) == 0 {
		PreInjectMiddle = DefaultMiddleWare
	}

	for i, m := range PreInjectMiddle {
		logger.Logger.Info(fmt.Sprintf("%s (%d) %s loaded", MiddleWare, i, m.Name))
		g.Use(MiddleWareMap[m.Name])
	}
}

// load from ./lib/*.so
func loadMiddlePlugins(g *gin.Engine) {
	plugins := findAllPlugins()
	for i, p := range plugins {
		pl, err := plugin.Open(p)
		if err != nil {
			logger.Logger.Info(fmt.Sprintf("%s (%d) %s plugin loaded failed", MiddleWare, i, p))
			continue
		}
		syb, err := pl.Lookup("Patch")
		if err != nil {
			logger.Logger.Info(fmt.Sprintf("%s (%d) %s plugin loaded failed", MiddleWare, i, p))
		}
		g.Use(syb.(gin.HandlerFunc))
		logger.Logger.Info(fmt.Sprintf("%s (%d) %s plugin loaded", MiddleWare, i, p))
	}
}

func findAllPlugins() []string {
	var plugins []string
	libPath := utils.CalDir(utils.GetAppDir(), PluginsPath)
	if _, e := os.Stat(libPath); os.IsNotExist(e) {
		return nil
	}
	err := filepath.Walk(libPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.Contains(info.Name(), ".so") {
			plugins = append(plugins, path)
		}
		return nil
	})
	if err != nil {
		logger.Logger.Info(fmt.Sprintf("%s failed to find plugins: %s", MiddleWare, err.Error()))
	}

	return plugins
}