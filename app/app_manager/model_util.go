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
	err := configen.ParseConfig(&cf, configen.Config, p)
	if err != nil || reflect.DeepEqual(cf, App{}) {
		return errors.New("can't load from config file")
	}

	AppManagerMap.Store(appName, cf)

	return nil
}

// SaveToFile 保存配置到文件中
func SaveToFile(cf *App, appName string) error {
	if ok := cf.Validate(); !ok {
		return errors.New("validate failed")
	}

	return configen.SaveConfig(
		cf,
		configen.Config,
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
				Start:     []string{"start.sh"},
				Stop:      []string{"stop.sh"},
				Restart:   []string{"restart.sh"},
				ForceKill: []string{"kill.sh"},
				Check:     "check.sh",
			},
			Meta: Meta{
				Author:      "",
				Language:    []string{},
				CreateDate:  "",
				Version:     "1.0.0",
				DynamicConf: false,
				ConfType:    "",
				ConfPath:    "",
			},
			RunData: RunData{
				Envs:  []string{},
				Ports: []int{},
				Host:  "localhost",
			},
		},
		configen.Config,
		utils.CalDir(
			utils.GetAppDir(), APPConfigsRoot, appName+ConfigSuffix))
}
