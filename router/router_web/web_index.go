/*
Project: Apollo web_index.go
Created: 2021/12/22 by Landers
*/

package router_web

import (
	"net/http"
	"sort"
	"strings"

	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/gin-gonic/gin"
)

type status struct {
	App    string
	Status string
}

// Index 主页
// @Summary 主页面
// @Description 主页
// @Tags Home
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router / [get]
func Index(c *gin.Context) {
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
	c.HTML(http.StatusOK, "base/index.tmpl", map[string]interface{}{
		"Apps": stat,
	})
}
