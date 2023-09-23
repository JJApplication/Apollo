/*
   Create: 2023/9/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package script_manager

import (
	"github.com/JJApplication/Apollo/logger"
	"time"
)

// 每隔1天重新加载一次脚本

func autoLoadScripts() {
	ticker := time.NewTicker(time.Duration(time.Second * 24 * 60 * 60))
	go func() {
		for {
			select {
			case <-ticker.C:
				logger.LoggerSugar.Infof("%s autoload scripts", ScriptManagerPrefix)
				UpdateScript(LoadScripts())
			}
		}
	}()
}
