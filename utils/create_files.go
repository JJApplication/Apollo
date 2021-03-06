/*
Project: Apollo create_files.go
Created: 2021/11/30 by Landers
*/

package utils

import (
	"os"
	"path/filepath"
)

// CreateFile 多级目录和文件的创建
func CreateFile(f string) error {
	if _, err := os.Stat(f); os.IsExist(err) {
		return err
	}

	dir := filepath.Dir(f)
	err := os.MkdirAll(dir, 0640)
	if err != nil {
		return err
	}

	file, err := os.Create(f)
	if err != nil {
		return err
	}

	return file.Close()
}

// CreateFileX 可执行文件创建 需要额外的x权限
func CreateFileX(f string) error {
	if _, err := os.Stat(f); os.IsExist(err) {
		return err
	}

	dir := filepath.Dir(f)
	err := os.MkdirAll(dir, 0640)
	if err != nil {
		return err
	}

	file, err := os.Create(f)
	if err != nil {
		return err
	}

	_ = file.Chmod(0750)
	return file.Close()
}
