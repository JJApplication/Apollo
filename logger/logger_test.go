/*
Project: Apollo logger_test.go
Created: 2021/11/22 by Landers
*/

package logger

import (
	"testing"

	"github.com/JJApplication/Apollo/config"
)

func TestInitLogger(t *testing.T) {
	err := InitLogger()
	t.Log(err)
}

func TestInitLogger2(t *testing.T) {
	config.ApolloConf.Log.EnableLog = "yes"
	config.ApolloConf.Log.EnableStack = "yes"
	config.ApolloConf.Log.EnableFunction = "yes"
	config.ApolloConf.Log.LogFile = "test.log"
	config.ApolloConf.Log.Encoding = "console"
	err := InitLogger()
	t.Log(err)
}
