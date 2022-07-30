/*
Create: 2022/7/31
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_container
package router_container

import (
	"github.com/JJApplication/Apollo/app/docker_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func GetAllImages(c *gin.Context) {
	res, err := docker_manager.ImageList()
	if err != nil {
		router.Response(c, err, false)
		return
	}

	router.Response(c, res, true)
}
