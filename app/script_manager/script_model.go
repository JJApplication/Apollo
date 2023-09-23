/*
   Create: 2023/9/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package script_manager

import (
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/utils"
	"path/filepath"
)

// 脚本配置文件模型

type scriptModel struct {
	Script     string            `json:"script"`
	ScriptName string            `json:"scriptName"`
	ScriptDes  string            `json:"scriptDes"`
	Args       []string          `json:"args"`
	EnvsAdd    []string          `json:"envsAdd"`
	Envs       map[string]string `json:"envs"`
	Workdir    string            `json:"workdir"`
}

// Run 异步任务运行脚本
func (s *scriptModel) Run() (string, error) {
	return "", nil
}

// RunCustom 自定义参数运行脚本
func (s *scriptModel) RunCustom(args []string, envs map[string]string, workdir string) (string, error) {
	s.Args = args
	s.Envs = envs
	s.Workdir = workdir

	return "", nil
}

// Check 检查脚本对应的运行文件是否存在
func (s *scriptModel) Check() bool {
	return utils.FileExist(filepath.Join(config.GlobalScriptRoot, s.Script))
}
