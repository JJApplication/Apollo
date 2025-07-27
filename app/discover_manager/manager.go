/*
Create: 2023/3/17
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

package discover_manager

import (
	"path/filepath"

	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
)

const DiscoverManagerPrefix = "[APPDiscover Manager]"

// 服务发现
//
// 原理：根据octopusTree的模型在服务目录下进行自动发现
// 如果存在模型文件 但是服务目录下不存在此模型对应的微服务则从发现的服务中卸载已经注册的服务

func InitDiscoverManager() {
	logger.LoggerSugar.Infof("%s start", DiscoverManagerPrefix)
	logger.LoggerSugar.Infof("%s find APP_ROOT: %s", DiscoverManagerPrefix, config.ApolloConf.APPRoot)
	createAppDiscover()
	createNoEngineDiscover()
}

var (
	appDiscover      *DiscoverRule
	noEngineDiscover *DiscoverRule
)

func createAppDiscover() {
	appDiscover = NewDiscoverRule(DiscoverRule{
		WatchRoot:   filepath.Join(config.ApolloConf.APPRoot, ".octopus"),
		Include:     nil,
		Exclude:     nil,
		CompareFlag: FlagAll,
	})
}

func createNoEngineDiscover() {
	noEngineDiscover = NewDiscoverRule(DiscoverRule{
		WatchRoot:   filepath.Join(config.ApolloConf.APPRoot, "NoEngine"),
		Include:     nil,
		Exclude:     []string{"nginx.conf", "conf", "conf.d", "cache", "apps.json", "noengine.app.json", "domain.json", "noengine.domain.json"},
		CompareFlag: FlagAll,
	})
}

func GetAppDiscover() *DiscoverRule {
	return appDiscover
}

func GetNoEngineDiscover() *DiscoverRule {
	return noEngineDiscover
}
