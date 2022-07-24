/*
Create: 2022/7/24
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package utils
package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type File struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Files []File `json:"files"`
	Extra string `json:"extra"`
}

const (
	// MaxDepth 可以通过环境变量ApolloMaxDepth自定义
	MaxDepth = 5
	KB       = 1024
	MB       = 2 << 20
	GB       = 2 << 30
)

var ApolloMaxDepth int

func init() {
	env := os.Getenv("ApolloMaxDepth")
	if env != "" {
		d, e := strconv.Atoi(env)
		if e == nil && d > 0 {
			ApolloMaxDepth = d
		} else {
			ApolloMaxDepth = MaxDepth
		}
	} else {
		ApolloMaxDepth = MaxDepth
	}
}

// GetFileTreeAllDepth 递归返回路径树 为避免卡顿最大遍历深度为5
func GetFileTreeAllDepth(offset, dir string) []File {
	if CalcPathDepth(offset, dir) >= ApolloMaxDepth {
		return nil
	}

	var filesTree []File
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	for _, file := range files {
		if file.IsDir() {
			filesTree = append(filesTree, File{
				Type:  "directory",
				Name:  file.Name(),
				Files: GetFileTreeAllDepth(offset, path.Join(dir, file.Name())),
				Extra: "#",
			})
		} else {
			filesTree = append(filesTree, File{
				Type:  "file",
				Name:  file.Name(),
				Extra: CalcFileSize(file.Size()),
			})
		}
	}

	return filesTree
}

// CalcFileSize 统计byte kb mb gb
func CalcFileSize(s int64) string {
	if s == 0 {
		return "0kb"
	} else if s < KB {
		return fmt.Sprintf("%sb", strconv.FormatInt(s, 10))
	} else if s >= KB && s < MB {
		return fmt.Sprintf("%skb", strconv.FormatInt(s/KB, 10))
	} else if s >= MB && s < GB {
		return fmt.Sprintf("%smb", strconv.FormatInt(s/MB, 10))
	} else {
		return fmt.Sprintf("%sgb", strconv.FormatInt(s/GB, 10))
	}
}

// CalcPathDepth 计算路径深度
// offsetDir 初始目录需要减去这个长度
// eg: offset: /code/gen dir: /code/gen/map
// depth = 4 - 3 = 1
func CalcPathDepth(offsetDir, dir string) int {
	offset := strings.Split(offsetDir, string(os.PathSeparator))
	l := strings.Split(dir, string(os.PathSeparator))
	return len(l) - len(offset) + 1
}
