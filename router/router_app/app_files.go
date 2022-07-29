/*
Create: 2022/7/29
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_app
package router_app

import (
	"path"

	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-gonic/gin"
)

// Upload 默认支持多文件上传
func Upload(c *gin.Context) {
	uploadPath := c.Query("path")
	if !isInAppRoot(uploadPath) {
		router.Response(c, "path not exists", false)
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		router.Response(c, err, false)
		return
	}
	for _, f := range form.File["files"] {
		// 文件保存到app的目录下 需要检验
		err := c.SaveUploadedFile(f, getSavePath(uploadPath, f.Filename))
		if err != nil {
			router.Response(c, err, false)
			return
		}
	}

	router.Response(c, "ok", true)
}

func Remove(c *gin.Context) {
	f := c.Query("file")
	if !isInAppRoot(f) {
		router.Response(c, "file not exists", false)
		return
	}
	err := utils.RemoveFile(getSavePath("", f))
	if err != nil {
		router.Response(c, err, false)
		return
	}
	router.Response(c, "", true)
}

func isInAppRoot(p string) bool {
	if p == "" {
		return utils.FileExist(config.ApolloConf.APPRoot)
	}
	realPath := path.Join(config.ApolloConf.APPRoot, p)
	return utils.FileExist(realPath)
}

func getSavePath(p, f string) string {
	if p == "" {
		return path.Join(config.ApolloConf.APPRoot, f)
	}
	return path.Join(config.ApolloConf.APPRoot, p, f)
}
