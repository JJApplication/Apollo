/*
   Create: 2024/1/6
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package log_manager

import (
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/utils"
	"path/filepath"
	"strings"
	"sync"
)

const MaxLogSize = 100 * 1024

type Collector struct {
	readLock  *sync.Mutex
	AppName   string
	AppLogDir string
}

func (c *Collector) init() {
	c.AppLogDir = filepath.Join(config.ApolloConf.APPLogDir, c.AppName)
	c.readLock = new(sync.Mutex)
}

func (c *Collector) checkDir() bool {
	return c.AppLogDir != ""
}

// 读取实时日志
// 全量读取时最大200kb 超过时仅读取最后200kb
func (c *Collector) readLog() string {
	if !c.checkDir() {
		return ""
	}
	c.readLock.Lock()
	defer c.readLock.Unlock()
	logFile := filepath.Join(c.AppLogDir, c.AppName+".log")
	if !utils.FileExist(logFile) {
		return ""
	}
	size := utils.GetFileSize(logFile)
	// 小于等于2mb时 直接读取
	if size <= MaxLogSize {
		return utils.ReadFileString(logFile)
	}
	return utils.ReadFileSlice(logFile, MaxLogSize, true)
}

// 获取日志目录下的文件
func (c *Collector) getLogDir() []utils.FileDetail {
	if !c.checkDir() {
		return nil
	}
	return utils.ReadDirDetail(c.AppLogDir)
}

// 判断合法性后获取压缩日志或日志包的根下载路径
// logfile为前台返回的不带路径的日志名称 需要拼接上日志路径
func (c *Collector) getLogFile(logFile string) string {
	if !c.checkDir() {
		return ""
	}
	if logFile == "" || strings.Contains(logFile, "../") {
		return ""
	}
	return filepath.Join(c.AppLogDir, logFile)
}
