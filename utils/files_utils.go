/*
Project: Apollo create_files.go
Created: 2021/11/30 by Landers
*/

package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	ArchiveType = ".tar.gz"
	RotateSize  = 20 << 20 // 20mb
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

// CreateOrCoverFile 覆盖文件或者创建文件
func CreateOrCoverFile(f string) error {
	return nil
}

// RemoveFile 尽力删除 文件不存在 不报错
func RemoveFile(f string) error {
	if !FileExist(f) {
		return nil
	}
	return os.RemoveAll(f)
}

// ArchiveFile 打包文件 默认为tar.gz
// fileName 要打包的文件
// dst 要生成的文件名
// autoClear 自动清空文件
// 写时复制 会导致writeToolong 将当前文件备份到tmp后压缩并删除
func ArchiveFile(fileName string, dst string, autoClear bool) error {
	if !strings.HasSuffix(dst, ArchiveType) {
		dst = dst + ArchiveType
	}
	_, err := os.Stat(fileName)
	if err != nil {
		return err
	}

	// copy when write
	fileTmpName := copyWhenWrite(fileName)
	fileInfo, err := os.Stat(fileTmpName)
	if err != nil {
		return err
	}

	f, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer f.Close()

	// gzip
	gzw := gzip.NewWriter(f)
	defer gzw.Close()

	// tar
	tarw := tar.NewWriter(gzw)
	defer tarw.Close()

	// start to zip
	file, err := os.Open(fileTmpName)
	if err != nil {
		return err
	}
	defer file.Close()

	header := new(tar.Header)
	// name需要为相对路径 受限于tar
	header.Name = filepath.Base(filepath.Clean(fileTmpName))
	header.Size = fileInfo.Size()
	header.Mode = int64(fileInfo.Mode())
	header.ModTime = fileInfo.ModTime()

	err = tarw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tarw, file)
	if err != nil {
		return err
	}

	// 自动清理不保证成功
	if autoClear {
		_ = RemoveFile(fileTmpName)
		return ClearFile(fileName)
	}
	return nil
}

// ClearFile 清空文件
func ClearFile(f string) error {
	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	return ioutil.WriteFile(f, []byte(""), 0664)
}

// GetFileSize 获取文件大小 不存在时为0
func GetFileSize(file string) int64 {
	if FileExist(file) {
		info, err := os.Stat(file)
		if err != nil {
			return 0
		}
		return info.Size()
	}
	return 0
}

// 不保证成功的写时复制
// 默认的目录/tmp
// 返回tmpDst
func copyWhenWrite(src string) string {
	tmpFile := filepath.Join(os.TempDir(), filepath.Base(src))
	_ = copyN(src, tmpFile)
	return tmpFile
}

// Rotate 日志绕接
// 计算大小阈值 超过才会压缩 否则跳过
func Rotate(appLog string) error {
	if GetFileSize(appLog) < RotateSize {
		return nil
	}
	gzName := fmt.Sprintf("%s-%s", appLog, TimeNowBetterSep())
	return ArchiveFile(appLog, gzName, true)
}
