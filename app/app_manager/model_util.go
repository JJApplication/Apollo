/*
Project: dirichlet model_util.go
Created: 2021/11/20 by Landers
*/

package app_manager

import (
	"errors"
	"reflect"
	"strings"

	"github.com/landers1037/configen"
	"github.com/landers1037/dirichlet/utils"
)

// 模型的序列化 反序列化

// LoadToMap 直接更新到全局的管理字典中
func LoadToMap(appName string) error {
	return loadFromApp(appName)
}

// 单个服务配置文件的加载
func loadFromApp(appName string) error {
	var cf App
	p := utils.CalDir(utils.GetAppDir(), APPConfigsRoot, appName+ConfigSuffix)
	err := configen.ParseConfig(&cf, configen.Pig, p)
	if err != nil || reflect.DeepEqual(cf, App{}) {
		return errors.New("can't load from config file")
	}

	APPManager.APPManagerMap.Store(appName, cf)

	return nil
}

// SaveToFile 保存配置到文件中
func SaveToFile(cf *App, appName string) error {
	if ok := cf.Validate(); !ok {
		return errors.New("validate failed")
	}

	return configen.SaveConfig(
		cf,
		configen.Pig,
		utils.CalDir(
			utils.GetAppDir(), APPConfigsRoot, appName+ConfigSuffix))
}

// NewApp 新建一个APP配置文件
func NewApp(appName string) error {
	return configen.SaveConfig(
		App{
			Name:          appName,
			ID:            "app_" + strings.ToLower(appName),
			Type:          TypeService,
			ReleaseStatus: Published,
			EngDes:        "default english description",
			CHSDes:        "默认中文描述",
			ManageCMD: CMD{
				Start:     "start.sh",
				Stop:      "stop.sh",
				Restart:   "restart.sh",
				ForceKill: "kill.sh",
				Check:     "check.sh",
			},
			Meta: Meta{
				Author:      "",
				Domain:      "",
				Language:    []string{},
				CreateDate:  "",
				Version:     "1.0.0",
				DynamicConf: false,
				ConfType:    "",
				ConfPath:    "",
			},
			RunData: RunData{
				Envs:       []string{},
				Ports:      []int{},
				RandomPort: true,
				Host:       "localhost",
			},
		},
		configen.Pig,
		utils.CalDir(
			utils.GetAppDir(), APPConfigsRoot, appName+ConfigSuffix))
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
