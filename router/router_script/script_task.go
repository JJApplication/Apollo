package router_script

import (
	"github.com/JJApplication/Apollo/app/script_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

// ScriptTaskList 获取全部历史任务列表
func ScriptTaskList(c *gin.Context) {
	result := script_manager.GetScriptTasks()
	router.Response(c, result, true)
}

// ScriptTaskByName 根据脚本名称获取其列表
func ScriptTaskByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		router.Response(c, nil, true)
		return
	}
	result := script_manager.GetScriptTaskByName(name)
	router.Response(c, result, true)
}

// ScriptTaskStart 开始执行脚本
func ScriptTaskStart(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		router.Response(c, nil, false)
		return
	}
	if err := script_manager.ExecuteScript(name); err != nil {
		router.Response(c, nil, false)
		return
	}
	router.Response(c, nil, true)
}

// ScriptTaskStop 强制停止任务
// 如果对应的任务存在且正在执行则停止其上下文
func ScriptTaskStop(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		router.Response(c, nil, false)
		return
	}

	if err := script_manager.StopScript(name); err != nil {
		router.Response(c, nil, false)
		return
	}
	router.Response(c, nil, true)
}

// ScriptTaskDelete 删除某个任务
// 理论上只能删除状态停止的任务
func ScriptTaskDelete(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		router.Response(c, nil, false)
		return
	}
	if err := script_manager.DeleteTask(name); err != nil {
		router.Response(c, nil, false)
		return
	}
	router.Response(c, nil, true)
}
