/*
Project: Apollo back_up.go
Created: 2021/12/29 by Landers
*/

package utils

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gookit/goutil/fsutil"
	copyfs "github.com/otiai10/copy"
)

// 全局的备份方法

var (
	BackBase   = "backup-%s.zip" // 拼接ServiceRoot
	BackupFlag = "/var/.backup"
	BackTmp    = "/tmp/Apollo"
)

// Backup 备份初始化，全局只能有一个备份任务,使用标志位判断
// src 为要备份的目录
func Backup(src, back string) error {
	if checkFlag() {
		return errors.New("backup progress is running")
	}
	setFlag()
	err := startBackup(src, back)
	defer rmFlag()
	if err != nil {
		return err
	}
	return nil
}

// 检查标志文件
func checkFlag() bool {
	if FileExist(BackupFlag) {
		return true
	}
	return false
}

// 设置标志文件
func setFlag() {
	_ = os.WriteFile(BackupFlag, []byte(""), 0644)
}

// 删除标志文件
func rmFlag() {
	_ = os.Remove(BackupFlag)
}

// 开始备份，避免数据io异常先cp到/tmp下操作
func startBackup(src, back string) error {
	if FileNotExist(src) {
		return errors.New("backup src is not exist")
	}
	if FileNotExist(back) {
		return errors.New("backup dst is not exist")
	}
	// 已经存在则删除原有的文件
	err := os.RemoveAll(BackTmp)
	if err != nil {
		return err
	}
	err = os.MkdirAll(BackTmp, 0644)
	if err != nil {
		return err
	}
	if err = copyDir(src); err != nil {
		return err
	}
	return zipDir(back)
}

// io cp操作
// 拷贝要备份的目录到临时目录下
func copyDir(src string) error {
	return copyfs.Copy(src, BackTmp)
}

// 压缩目录
// src为要压缩的原目录
// 已经存在备份文件时删除后进行压缩
func zipDir(back string) error {
	BackDir := path.Join(back, fmt.Sprintf(BackBase, TimeNowBetterSep()))
	if fsutil.FileExists(BackDir) {
		err := fsutil.DeleteIfFileExist(BackDir)
		if err != nil {
			return err
		}
	}
	return zipFunc(BackDir, BackTmp)
}

func zipFunc(dst, src string) error {
	zipFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	err = filepath.Walk(BackTmp, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filter(info.Name()) {
			return filepath.SkipDir
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(path, filepath.Dir(src)+"/")
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})
	return err
}

var filterFiles = []string{}

// 过滤文件 文件夹
// 不过滤时返回true
func filter(name string) bool {
	for _, f := range filterFiles {
		if f == name {
			return false
		}
	}
	return true
}
