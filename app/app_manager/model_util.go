/*
Project: Apollo model_util.go
Created: 2021/11/20 by Landers
*/

package app_manager

import (
	"errors"

	"github.com/JJApplication/Apollo/utils"
	"github.com/JJApplication/octopus_meta"
)

// 模型的序列化 反序列化

// LoadToMap 直接更新到全局的管理字典中
func LoadToMap(appName string) error {
	return loadFromApp(appName)
}

// 单个服务配置文件的加载
func loadFromApp(appName string) error {
	var cf octopus_meta.App
	cf, err := octopus_meta.LoadApp(appName)
	if err != nil || cf.Validate() {
		return errors.New("can't load from config file")
	}

	APPManager.APPManagerMap.Store(appName, App{Meta: cf})

	return nil
}

// SaveToFile 保存配置到文件中
func SaveToFile(cf *App, appName string) error {
	if ok := cf.Validate(); !ok {
		return errors.New("validate failed")
	}
	return octopus_meta.SaveAppMeta(cf.Meta, appName)
}

// NewApp 新建一个APP配置文件
func NewApp(appName string) error {
	return octopus_meta.NewAppMeta(appName)
}

// NewAppScript 新增app 命令目录
// 默认只创建start stop check
func NewAppScript(appName string) error {
	for _, sh := range []string{"start.sh", "stop.sh", "check.sh"} {
		err := utils.CreateFileX(utils.CalDir(
			utils.GetAppDir(),
			APPScriptsRoot,
			appName,
			sh))
		if err != nil {
			return err
		}
	}
	return nil
}
