package router_database

import (
	"fmt"
	"github.com/JJApplication/Apollo/app/database_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func GetDatabaseList(c *gin.Context) {
	fmt.Println("GetDatabaseList", database_manager.Get())
	res := database_manager.Get().List()
	router.Response(c, res, true)
}

func GetDatabaseInfo(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		router.Response(c, nil, false)
		return
	}
	res := database_manager.Get().Get(name)
	router.Response(c, res, true)
}
