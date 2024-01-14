/*
   Create: 2024/1/14
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package noengine_manager

import (
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"sync"
)

// 写入NoEnine目录下的端口映射表
// 与内部manager中维护的同步

var (
	syncMapLock = sync.Mutex{}
)

func syncMap() {
	syncMapLock.Lock()
	defer syncMapLock.Unlock()
	var temp = make(map[string]string)
	NoEngineMap.Range(func(key, value any) bool {
		if key.(string) == "" {
			return false
		}
		app := value.(NoEngineTemplate)
		if len(app.Ports) > 0 {
			temp[key.(string)] = app.Ports[0].HostPort
		}
		return true
	})

	if err := utils.Save2JsonFile(temp, NoEngineAPPMap); err != nil {
		logger.LoggerSugar.Errorf("save %s error: %s", NoEngineAPPMap, err.Error())
		return
	}
	logger.LoggerSugar.Infof("save %s success", NoEngineAPPMap)
}
