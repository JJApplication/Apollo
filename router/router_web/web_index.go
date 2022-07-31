/*
Project: Apollo web_index.go
Created: 2021/12/22 by Landers
*/

package router_web

import (
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
	c.File("./web/index.html")
}

func Favicon(c *gin.Context) {
	c.File("./web/favicon.ico")
}
