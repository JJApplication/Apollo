/*
   Create: 2023/9/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package script_manager

import (
	"errors"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/fushin/utils/cmd"
	"os/exec"
	"path/filepath"
	"strings"
)

// 脚本的执行全部为异步任务
// 脚本执行的结果会记录到数据库中

// RunWithTimeout 带上下文超时的运行回调
// 考虑到存在大备份任务，超时时间设置为6h
func runWithTimeout(script scriptModel, args string) error {
	if hasTaskRunning(script.ScriptName) {
		return errors.New("task is running")
	}
	if !script.Check() {
		return errors.New("script not found")
	}
	scriptLock.Lock()
	defer scriptLock.Unlock()

	// 创建任务
	uuid, err := startTask(script.ScriptName)
	if err != nil {
		return err
	}

	// 拼接参数
	var sh string
	if args == "" {
		sh = filepath.Join(config.GlobalScriptRoot, script.Script)
	} else {
		sh = strings.Join([]string{filepath.Join(config.GlobalScriptRoot, script.Script), args}, " ")
	}

	// 创建cmd进程
	c := cmd.NewCmder().Run(sh)
	go func() {
		if err = c.Run(); err != nil {
			logger.LoggerSugar.Errorf("%s run script: %s err:%v", ScriptManagerPrefix, script.ScriptName, err)
			updateTaskStatus(StatusFail, uuid, err.Error())
		} else {
			updateTaskStatus(StatusSuccess, uuid, "")
		}
		// 停止任务时删除manager中的实例
		scriptLock.Lock()
		delete(scriptManager, script.ScriptName)
		scriptLock.Unlock()
	}()
	scriptManager[script.ScriptName] = struct {
		uuid   string
		status string
		cmd    *exec.Cmd
	}{uuid: uuid, status: StatusRunning, cmd: c}

	logger.LoggerSugar.Infof("%s start script: %s, uuid: %s", ScriptManagerPrefix, script.ScriptName, uuid)
	return nil
}

// 根据运行时中存储的cmd上下文停止
// 当服务丢失上下文时直接根据进程kill
func stopScript(script scriptModel) error {
	if !hasTaskRunning(script.ScriptName) {
		return errors.New("task is not running")
	}
	if !script.Check() {
		return errors.New("script not found")
	}
	scriptLock.Lock()
	defer scriptLock.Unlock()

	sc, ok := scriptManager[script.ScriptName]
	if !ok {
		return nil
	}
	if err := sc.cmd.Process.Kill(); err != nil {
		updateTaskStatus(StatusError, sc.uuid, err.Error())
	} else {
		updateTaskStatus(StatusError, sc.uuid, "")
	}

	logger.LoggerSugar.Infof("%s force stop script: %s, uuid: %s", ScriptManagerPrefix, script.ScriptName, sc.uuid)
	// 清理
	delete(scriptManager, script.ScriptName)
	return nil
}
