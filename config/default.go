/*
Project: dirichlet default.go
Created: 2021/11/20 by Landers
*/

package config

import (
	"os"
)

// GetDefault 设置默认配置
func GetDefault(ck string) string {
	return os.Getenv(ck)
}
