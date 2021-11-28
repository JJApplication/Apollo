/*
Project: dirichlet load_cf_test.go
Created: 2021/11/20 by Landers
*/

package app_manager

import (
	"testing"
)

// 测试加载配置文件
func TestLoadManagerCf(t *testing.T) {
	err := LoadManagerCf()
	if err != nil {
		t.Error(err)
	}
}
