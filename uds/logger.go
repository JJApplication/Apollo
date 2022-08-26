/*
Create: 2022/8/26
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

import (
	"go.uber.org/zap"
)

// 继承fushin Logger

type udsLogger struct {
	logger *zap.Logger
}

func (u udsLogger) Info(v ...interface{}) {
	u.logger.Sugar().Info(v...)
}

func (u udsLogger) Warn(v ...interface{}) {
	u.logger.Sugar().Warn(v...)
}

func (u udsLogger) Error(v ...interface{}) {
	u.logger.Sugar().Error(v...)
}

func (u udsLogger) InfoF(fmt string, v ...interface{}) {
	u.logger.Sugar().Infof(fmt, v...)
}

func (u udsLogger) WarnF(fmt string, v ...interface{}) {
	u.logger.Sugar().Warnf(fmt, v...)
}

func (u udsLogger) ErrorF(fmt string, v ...interface{}) {
	u.logger.Sugar().Errorf(fmt, v...)
}
