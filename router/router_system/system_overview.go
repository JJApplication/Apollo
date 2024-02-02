/*
   Create: 2023/9/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_system

import (
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/vars"
	"github.com/gin-gonic/gin"
)

type overview struct {
	GOOS      string `json:"goos"`
	GOARCH    string `json:"goarch"`
	GOVersion string `json:"goVersion"`
	BuildDate string `json:"buildDate"`
	GitCommit string `json:"gitCommit"`
}

func SystemOverview(c *gin.Context) {
	router.Response(c, overview{
		GOOS:      vars.GOOS,
		GOARCH:    vars.GOARCH,
		GOVersion: vars.GOVersion,
		BuildDate: vars.BuildDate,
		GitCommit: vars.GitCommit,
	}, true)
}
