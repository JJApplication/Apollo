/*
   Create: 2024/1/6
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package log_manager

import (
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
)

type LogDetail struct {
	Name       string
	FileSize   string
	ModifyTime string
}

// HasCollector 是否已经创建采集器
func HasCollector(app string) bool {
	_, ok := LogManagerPool.Load(app)
	return ok
}

func MustGetCollector(app string) (Collector, error) {
	col, err := getCollector(app)
	if err != nil {
		logger.LoggerSugar.Errorf("%s get collector: %s error: %s", LogManager, app, err.Error())
		logger.LoggerSugar.Infof("%s create collector: %s", LogManager, app)
		if err = initCollector(app); err != nil {
			logger.LoggerSugar.Errorf("%s create collector: %s error: %s", LogManager, app, err.Error())
			return Collector{}, err
		}
		col, err := getCollector(app)
		if err != nil {
			return Collector{}, err
		}
		return col, nil
	}

	return col, nil
}

// GetCollectors 获取已经创建的采集器列表
func GetCollectors() []string {
	var list []string
	LogManagerPool.Range(func(key, value any) bool {
		list = append(list, key.(string))
		return true
	})

	return list
}

// GetAPPLog 获取APP实时日志
func GetAPPLog(app string) string {
	col, err := MustGetCollector(app)
	if err != nil {
		logger.LoggerSugar.Errorf("%s get %s's log error: %s", LogManager, app, err.Error())
		return ""
	}

	return col.readLog()
}

// DownloadAPPLogFile 以文件名方式直接由引擎处理返回文件路径下载
func DownloadAPPLogFile(app, logName string) string {
	col, err := MustGetCollector(app)
	if err != nil {
		logger.LoggerSugar.Errorf("%s get %s's logfile error: %s", LogManager, app, err.Error())
		return ""
	}

	return col.getLogFile(logName)
}

// GetAPPLogList 返回APP日志目录下的全部日志文件
func GetAPPLogList(app string) []LogDetail {
	col, err := MustGetCollector(app)
	if err != nil {
		logger.LoggerSugar.Errorf("%s get %s's logList error: %s", LogManager, app, err.Error())
		return nil
	}

	var result []LogDetail
	list := col.getLogDir()
	for _, detail := range list {
		result = append(result, LogDetail{
			Name:       detail.Name,
			FileSize:   utils.CalcFileSize(detail.FileSize),
			ModifyTime: utils.GetTimeByFormat(utils.TimeFormat, detail.ModifyTime),
		})
	}

	return result
}
