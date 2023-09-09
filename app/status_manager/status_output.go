/*
   Create: 2023/7/31
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package status_manager

import (
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/utils"
	"path/filepath"
	"sync"
)

// 状态树的输出文件

const (
	StatusOptFile = "status_tree.json"
)

func GetStatusOptFile() string {
	return filepath.Join(config.ApolloConf.APPCacheDir, StatusOptFile)
}

func writeStatusOptFile(data StatusTree) error {
	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()

	return utils.Save2JsonFile(data, GetStatusOptFile())
}
