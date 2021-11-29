/*
Project: dirichlet load_cf.go
Created: 2021/11/20 by Landers
*/

package app_manager

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/landers1037/configen"
	"github.com/landers1037/dirichlet/logger"
	"github.com/landers1037/dirichlet/utils"
)

// LoadManagerCf 加载所有配置文件到全局的字典中
func LoadManagerCf() error {
	// 保证读取到配置后再刷新字典
	tm, ok := loadAllCfs(getAPPCfs())
	// 每次刷新
	if ok {
		AppManagerMap.Range(func(key, value interface{}) bool {
			AppManagerMap.Delete(key)
			return true
		})

		for k, v := range tm {
			logger.Logger.Info(fmt.Sprintf("store app [%s] config: %+v", k, v))
			AppManagerMap.Store(k, v)
		}
	}

	// 未刷新保持缓存的map

	return nil
}

// 从配置目录中获取配置文件
// 文件名为对应的APPName
func getAPPCfs() []string {
	var cfs []string
	p := utils.CalDir(utils.GetAppDir(), APPConfigsRoot)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return cfs
	}

	err := filepath.Walk(p, func(path string, info fs.FileInfo, err error) error {
		if err == nil {
			if !info.IsDir() {
				cfs = append(cfs, path)
			}
		}

		return err
	})

	if err != nil {
		//todo log
	}

	return cfs
}

func loadAllCfs(cfs []string) (map[string]App, bool) {
	fmt.Println(cfs)
	var loadStatus = true
	tm := make(map[string]App, 0)
	for _, c := range cfs {
		var appCfg App
		err := configen.ParseConfig(&appCfg, configen.Config, c)
		if err != nil || reflect.DeepEqual(appCfg, App{}) {
			loadStatus = false
			continue
		}

		// get name
		name := strings.Trim(filepath.Base(c), filepath.Ext(c))
		// save to map
		tm[name] = appCfg
	}

	return tm, loadStatus
}
