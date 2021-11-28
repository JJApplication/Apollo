/*
Project: dirichlet logger_test.go
Created: 2021/11/22 by Landers
*/

package logger

import (
	"testing"

	"github.com/landers1037/dirichlet/config"
)

func TestInitLogger(t *testing.T) {
	err := InitLogger()
	t.Log(err)
}

func TestInitLogger2(t *testing.T) {
	config.DirichletConf.Log.EnableLog = "yes"
	config.DirichletConf.Log.EnableStack = "yes"
	config.DirichletConf.Log.EnableFunction = "yes"
	config.DirichletConf.Log.LogFile = "test.log"
	config.DirichletConf.Log.Encoding = "console"
	err := InitLogger()
	t.Log(err)
}
