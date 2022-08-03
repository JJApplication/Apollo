/*
Project: Apollo create_files.go
Created: 2021/11/30 by Landers
*/

package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	ArchiveType = ".tar.gz"
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
func ArchiveFile(fileName string, dst string, autoClear bool) error {
	if !strings.HasSuffix(dst, ArchiveType) {
		dst = dst + ArchiveType
	}
	fileInfo, err := os.Stat(fileName)
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
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	header := new(tar.Header)
	// name需要为相对路径 受限于tar
	header.Name = filepath.Base(filepath.Clean(fileName))
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
	if autoClear {
		return ClearFile(fileName)
	}
	return nil
}

// ClearFile 清空文件
func ClearFile(f string) error {
	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	return ioutil.WriteFile(f, []byte(""), 0644)
}
