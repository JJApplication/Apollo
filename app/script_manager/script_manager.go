/*
   Create: 2023/9/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package script_manager

import (
	"errors"
	"github.com/JJApplication/Apollo/logger"
	"os/exec"
	"sync"
)

var (
	scriptLock    sync.Mutex
	scriptManager map[string]struct {
		uuid   string
		status string
		cmd    *exec.Cmd
	}
)

var scripts []scriptModel

func init() {
	UpdateScript(LoadScripts())
	autoLoadScripts()
}

func UpdateScript(data []scriptModel) {
	lock := sync.Mutex{}
	lock.Lock()
	scripts = data
	defer lock.Unlock()
}

func GetScripts() []scriptModel {
	return scripts
}

// ExecuteScript 执行脚本
func ExecuteScript(name, args string) error {
	var sc scriptModel
	for _, script := range scripts {
		if script.ScriptName == name {
			sc = script
		}
	}
	if sc.Script == "" {
		return errors.New("script not found")
	}
	return runWithTimeout(sc, args)
}

// StopScript 强行停止脚本
func StopScript(name string) error {
	var sc scriptModel
	for _, script := range scripts {
		if script.ScriptName == name {
			sc = script
		}
	}
	if sc.Script == "" {
		return errors.New("script not found")
	}
	return stopScript(sc)
}

// GetScriptTasks 获取所有执行的脚本任务
func GetScriptTasks() []ScriptTask {
	return getTaskList()
}

// GetScriptTaskByName 获取某个脚本的任务列表
func GetScriptTaskByName(name string) []ScriptTask {
	return getTaskListByName(name)
}

func DeleteTask(uuid string) error {
	if err := deleteTask(uuid); err != nil {
		logger.LoggerSugar.Errorf("%s delete task [%s] error:%v", ScriptManagerPrefix, uuid, err)
		return err
	}
	logger.LoggerSugar.Infof("%s delete task [%s]", ScriptManagerPrefix, uuid)
	return nil
}
