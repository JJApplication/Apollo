// Package router_repo
package router_repo

import (
	"github.com/JJApplication/Apollo/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerRepo := r.Group("/api/repo")
	{
		routerRepo.GET("/list", ListRepos)
		routerRepo.GET("/:org/:name", GetRepo)

	}
	routerRepoWithAuth := r.Group("/api/repo", middleware.MiddleWareAuth())
	{
		routerRepoWithAuth.POST("/commits/:org/:name", GetCommits)
		routerRepoWithAuth.POST("/sync/repos", SyncRepos)
		routerRepoWithAuth.POST("/sync/commits/:org/:name", SyncCommits)
	}
}
