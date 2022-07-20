/*
Project: Apollo model_util_test.go
Created: 2021/11/20 by Landers
*/

package app_manager

import (
	"testing"

	"github.com/JJApplication/octopus_meta"
)

// 加载模型文件
func TestLoadToMap(t *testing.T) {
	t.Logf("%+v", APPManager.APPManagerMap)
	err := loadFromApp("TestService")
	t.Log(err)
	// d, _ := AppManagerMap.Load("TestService")

	// range all
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		t.Logf("%s, %+v", key, value)
		return true
	})
	// t.Logf("%+v", d)
}

// 输出模型文件
func TestSaveToFile(t *testing.T) {
	err := octopus_meta.NewAppMeta("TestService")
	if err != nil {
		t.Log(err)
	}
}
