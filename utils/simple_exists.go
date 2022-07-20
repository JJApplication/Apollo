/*
Project: Apollo
Created: 2022/2/22 by Landers
*/

package utils

import "os"

// 简单的文件 文件夹存在性判断

func FileExist(p string) bool {
	if _, err := os.Stat(p); err == nil || os.IsExist(err) {
		return true
	}
	return false
}

func FileNotExist(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return true
	}
	return false
}

func FileExistOrCreate(p string) error {
	if FileNotExist(p) {
		return CreateFile(p)
	}
	return nil
}
