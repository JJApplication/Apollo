/*
Project: Apollo middleware.go
Created: 2021/11/26 by Landers
*/

package middleware

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"plugin"
	"strings"
	"time"

	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// 基于本地配置加载中间件
// 优先被初始化以让所有中间件判断自己是否被激活
// 两种模式：全局global 路由route 在路由模式下需要指定路由

const (
	MiddleWare  = "[MiddleWare]"
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
	"log": gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}
		return fmt.Sprintf("[Apollo] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	}),
	"recovery": gin.Recovery(),
	"cors":     MiddleWareCors(),
	"plnack":   MiddlewarePlnack(),
	"gzip":     gzip.Gzip(gzip.BestCompression, gzip.WithExcludedPaths([]string{"/api/"})),
}

// LoadMiddleWare 中间件加载
func LoadMiddleWare(g *gin.Engine) {
	PreInjectMiddle = LoadMiddleWareConfig()
	if len(PreInjectMiddle) == 0 {
		PreInjectMiddle = DefaultMiddleWare
	}

	for i, m := range PreInjectMiddle {
		if m.Active {
			logger.LoggerSugar.Infof("%s <%d> (%s) loaded", MiddleWare, i, m.Name)
			g.Use(MiddleWareMap[m.Name])
		} else {
			logger.LoggerSugar.Infof("%s <%d> (%s) disabled", MiddleWare, i, m.Name)
		}
	}
}

// LoadMiddlePlugins load from ./lib/*.so
func LoadMiddlePlugins(g *gin.Engine) {
	plugins := findAllPlugins()
	for i, p := range plugins {
		pl, err := plugin.Open(p)
		if err != nil {
			logger.LoggerSugar.Infof("%s (%d) %s plugin loaded failed", MiddleWare, i, p)
			continue
		}
		syb, err := pl.Lookup("Patch")
		if err != nil {
			logger.LoggerSugar.Infof("%s (%d) %s plugin loaded failed", MiddleWare, i, p)
		}
		g.Use(syb.(gin.HandlerFunc))
		logger.LoggerSugar.Infof("%s (%d) %s plugin loaded", MiddleWare, i, p)
	}
}

func findAllPlugins() []string {
	var plugins []string
	libPath := utils.CalDir(utils.GetAppDir(), PluginsPath)
	if utils.FileNotExist(libPath) {
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
		logger.LoggerSugar.Infof("%s failed to find plugins: %s", MiddleWare, err.Error())
	}

	return plugins
}
