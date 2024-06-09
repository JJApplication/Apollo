/*
   Create: 2024/1/10
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package noengine_manager

import (
	"errors"
	"github.com/JJApplication/Apollo/utils"
	"sync"
)

// NoEngine的端口映射缓存
// 用于sandwich的转发
// 数据会存储在noengine.map文件下

var (
	loadLock   = sync.Mutex{}
	localCache = "noengine.map"
)

// LoadNoEngineMap 供外部接口调用
func LoadNoEngineMap() (map[string]NoEngineTemplate, error) {
	loadLock.Lock()
	defer loadLock.Unlock()
	var res map[string]NoEngineTemplate
	if err := utils.ParseJsonFile(localCache, &res); err != nil {
		return res, err
	}
	return res, nil
}

// RefreshNoEngineMap 从内置缓存刷新Map到文件
func RefreshNoEngineMap() (map[string]NoEngineTemplate, error) {
	loadLock.Lock()
	defer loadLock.Unlock()
	data := GetAllNoEngineAPPsRt()
	if err := utils.Save2JsonFile(data, localCache); err != nil {
		return data, err
	}
	return data, nil
}

// ReloadNoEngineMap 重启时加载文件缓存到内置缓存
func ReloadNoEngineMap() (map[string]NoEngineTemplate, error) {
	if utils.FileNotExist(localCache) {
		return nil, errors.New("noengine cache not exist")
	}
	return LoadNoEngineMap()
}
