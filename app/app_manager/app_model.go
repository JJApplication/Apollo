/*
Project: Apollo app_model.go
Created: 2021/11/20 by Landers
*/

package app_manager

import (
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/octopus_meta"
)

// App model for app
type App struct {
	Meta octopus_meta.App `json:"meta" bson:"meta"`
}

// Validate 适用于model的检查器
func (app *App) Validate() bool {
	if !app.Meta.Validate() {
		logger.Logger.Error(appCodeMsg("app validate failed"))
		return false
	}

	return true
}
