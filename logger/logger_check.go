/*
Project: dirichlet logger_check.go
Created: 2021/11/26 by Landers
*/

package logger

import (
	"fmt"
	"time"

	"github.com/landers1037/dirichlet/config"
)

func Check() {
	fmt.Printf("[Dirichlet] 🚀🚀🚀\n⌛ StartTime: %s\n🔥 Listening on %s\n📁 ServiceRoot: %s\n📁 AppRoot: %s\n📁 AppManger: %s\n📁 AppLog: %s\n📁 BackUpDir: %s\n📁 CacheDir: %s\n\n",
		time.Now().String(),
		fmt.Sprintf("http://%s:%d", config.DirichletConf.Server.Host, config.DirichletConf.Server.Port),
		config.DirichletConf.ServiceRoot,
		config.DirichletConf.APPRoot,
		config.DirichletConf.APPManager,
		config.DirichletConf.APPLogDir,
		config.DirichletConf.APPBackUp,
		config.DirichletConf.APPCacheDir)
}
