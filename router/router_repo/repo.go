// Package router_repo
package router_repo

import (
	"fmt"

	"github.com/JJApplication/Apollo/app/repo_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

type reqSetEnv struct {
	Service string `json:"service"`
	Envs    []struct {
		Key     string `json:"key"`
		Value   string `json:"value"`
		Encrypt bool   `json:"encrypt"`
	} `json:"envs"`
}

func ListRepos(c *gin.Context) {
	repoManager := repo_manager.GetRepoManager()
	router.Response(c, repoManager.GetRepos(), true)
}

func GetRepo(c *gin.Context) {
	org := c.Param("org")
	name := c.Param("name")
	fullName := fmt.Sprintf("%s/%s", org, name)
	if fullName == "" {
		router.Response(c, nil, false)
		return
	}
	repoManager := repo_manager.GetRepoManager()
	data := repoManager.GetRepo(fullName)
	router.Response(c, data, true)
}

func SyncRepos(c *gin.Context) {
	repoManager := repo_manager.GetRepoManager()
	data := repoManager.SyncRepos()
	router.Response(c, data, true)
}

func GetCommits(c *gin.Context) {
	repoManager := repo_manager.GetRepoManager()
	org := c.Param("org")
	name := c.Param("name")
	fullName := fmt.Sprintf("%s/%s", org, name)
	if fullName == "" {
		router.Response(c, nil, false)
		return
	}
	data := repoManager.GetCommits(fullName)
	router.Response(c, data, true)
}

func SyncCommits(c *gin.Context) {
	repoManager := repo_manager.GetRepoManager()
	org := c.Param("org")
	name := c.Param("name")
	fullName := fmt.Sprintf("%s/%s", org, name)
	if fullName == "" {
		router.Response(c, nil, false)
		return
	}
	data := repoManager.SyncCommits(fullName)
	router.Response(c, data, true)
}
