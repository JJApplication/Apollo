/*
Project: Apollo logger_check.go
Created: 2021/11/26 by Landers
*/

package logger

import (
	"fmt"
	"time"

	"github.com/JJApplication/Apollo/config"
)

func Check() {
	fmt.Printf("[Apollo] ššš\nā StartTime: %s\nš„ Listening on %s\nš ServiceRoot: %s\nš AppRoot: %s\nš AppManger: %s\nš AppLog: %s\nš BackUpDir: %s\nš CacheDir: %s\n\n",
		time.Now().String(),
		fmt.Sprintf("http://%s:%d", config.ApolloConf.Server.Host, config.ApolloConf.Server.Port),
		config.ApolloConf.ServiceRoot,
		config.ApolloConf.APPRoot,
		config.ApolloConf.APPManager,
		config.ApolloConf.APPLogDir,
		config.ApolloConf.APPBackUp,
		config.ApolloConf.APPCacheDir)
}
