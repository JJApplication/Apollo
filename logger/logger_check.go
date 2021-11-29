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
	fmt.Printf("[Dirichlet]\nStartTime: %s\nListening on %s\nServiceRoot: %s\nAppRoot: %s\nAppManger: %s\nAppLog: %s\nBackUpDir: %s\nCacheDir: %s\n\n",
		time.Now().String(),
		fmt.Sprintf("http://%s:%d", config.DirichletConf.Server.Host, config.DirichletConf.Server.Port),
		config.DirichletConf.ServiceRoot,
		config.DirichletConf.APPRoot,
		config.DirichletConf.APPManager,
		config.DirichletConf.APPLogDir,
		config.DirichletConf.APPBackUp,
		config.DirichletConf.APPCacheDir)
}
