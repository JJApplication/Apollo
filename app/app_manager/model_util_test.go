/*
Project: dirichlet model_util_test.go
Created: 2021/11/20 by Landers
*/

package app_manager

import (
	"testing"
)

// 加载模型文件
func TestLoadToMap(t *testing.T) {
	t.Logf("%+v", &AppManagerMap)
	err := loadFromApp("TestService")
	t.Log(err)
	// d, _ := AppManagerMap.Load("TestService")

	// range all
	AppManagerMap.Range(func(key, value interface{}) bool {
		t.Logf("%s, %+v", key, value)
		return true
	})
	// t.Logf("%+v", d)
}

// 输出模型文件
func TestSaveToFile(t *testing.T) {
	err := SaveToFile(&App{
		Name:          "TestService",
		ID:            "testService",
		Type:          "Service",
		ReleaseStatus: "",
		EngDes:        "",
		CHSDes:        "",
		ManageCMD:     CMD{},
		Meta:          Meta{},
		RunData:       RunData{},
	}, "TestService")

	if err != nil {
		t.Log(err)
	}
}
