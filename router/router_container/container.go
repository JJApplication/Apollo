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

// 容器操作

func GetAllContainer(c *gin.Context) {
	res, err := docker_manager.ContainerList()
	if err != nil {
		router.Response(c, err, false)
		return
	}

	router.Response(c, res, true)
}

func StartContainer(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		router.Response(c, "", false)
		return
	}
	err := docker_manager.ContainerStart(id)
	if err != nil {
		router.Response(c, err, false)
		return
	}
	router.Response(c, "", true)
}

func StopContainer(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		router.Response(c, "", false)
		return
	}
	err := docker_manager.ContainerStop(id)
	if err != nil {
		router.Response(c, err, false)
		return
	}
	router.Response(c, "", true)
}

func RemoveContainer(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		router.Response(c, "", false)
		return
	}
	err := docker_manager.ContainerRemove(id)
	if err != nil {
		router.Response(c, err, false)
		return
	}
	router.Response(c, "", true)
}
