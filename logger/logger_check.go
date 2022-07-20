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
	fmt.Printf("[Apollo] 🚀🚀🚀\n⌛ StartTime: %s\n🔥 Listening on %s\n📁 ServiceRoot: %s\n📁 AppRoot: %s\n📁 AppManger: %s\n📁 AppLog: %s\n📁 BackUpDir: %s\n📁 CacheDir: %s\n\n",
		time.Now().String(),
		fmt.Sprintf("http://%s:%d", config.ApolloConf.Server.Host, config.ApolloConf.Server.Port),
		config.ApolloConf.ServiceRoot,
		config.ApolloConf.APPRoot,
		config.ApolloConf.APPManager,
		config.ApolloConf.APPLogDir,
		config.ApolloConf.APPBackUp,
		config.ApolloConf.APPCacheDir)
}
