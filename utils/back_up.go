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

	"github.com/JJApplication/Apollo/config"
	"github.com/gookit/goutil/fsutil"
)

// 全局的备份方法

var (
	BackDir    = path.Join(config.ApolloConf.APPBackUp, fmt.Sprintf("backup-%s.zip", TimeNowBetterSep())) // 拼接ServiceRoot
	BackDirOld = BackDir + ".old"
	BackupFlag = "/var/.backup"
	BackTmp    = "/tmp/Apollo"
)

// Backup 备份初始化，全局只能有一个备份任务,使用标志位判断
func Backup(src string) error {
	if checkFlag() {
		return errors.New("backup progress is running")
	}
	setFlag()
	err := startBackup(src)
	if err != nil {
		return err
	}
	rmFlag()
	return nil
}

// 检查标志文件
func checkFlag() bool {
	if FileExist(BackupFlag) {
		return false
	}
	if FileNotExist(BackupFlag) {
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
func startBackup(src string) error {
	if FileNotExist(BackTmp) {

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
	return zipDir()
}

// io cp操作
func copyDir(src string) error {
	return fsutil.CopyFile(src, BackTmp)
}

// 压缩目录
func zipDir() error {
	if fsutil.FileExists(BackDir) && !fsutil.FileExists(BackDirOld) {
		_ = fsutil.CopyFile(BackDir, BackDirOld)
		err := fsutil.DeleteIfFileExist(BackDir)
		if err != nil {
			return err
		}
	} else if fsutil.FileExists(BackDir) && fsutil.FileExists(BackDirOld) {
		err := fsutil.DeleteIfFileExist(BackDirOld)
		if err != nil {
			return err
		}
		_ = fsutil.CopyFile(BackDir, BackDirOld)
		err = fsutil.DeleteIfFileExist(BackDir)
		if err != nil {
			return err
		}
	}
	return zipFunc(BackTmp)
}

func zipFunc(src string) error {
	zipFile, err := os.Create(BackDir)
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
