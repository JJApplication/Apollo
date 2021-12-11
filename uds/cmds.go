/*
Project: dirichlet cmds.go
Created: 2021/12/9 by Landers
*/

package uds

import (
	"fmt"
	"strings"

	"github.com/landers1037/dirichlet/app/app_manager"
	"github.com/landers1037/dirichlet/config"
	"github.com/landers1037/dirichlet/utils"
)

// 支持的命令
var cmds = []string{
	"show",
	"app",
	"start",
	"stop",
	"restart",
	"status",
	"config",

	"task",
}

const (
	SUCCESS = "success"
)

func supportCmds(cmd string) bool {
	cmdsRaw := strings.Fields(cmd)
	if len(cmdsRaw) <= 0 {
		return false
	}
	for _, c := range cmds {
		if c == cmdsRaw[0] {
			return true
		}
	}
	return false
}

var cmdsMap = map[string]func(args ...string) (result string){
	"show": func(args ...string) (result string) {
		var s []string
		for _, cmd := range cmds {
			s = append(s, fmt.Sprintf("[%-7s] %s\n", cmd, helpMsg[cmd]))
		}

		return strings.Join(s, "")
	},
	"app": func(args ...string) (result string) {
		if len(args) == 0 {
			return "empty app, do you mean all?"
		}
		appName := args[0]
		if appName == "all" {
			apps, err := app_manager.GetAllApp()
			if err != nil {
				return fmt.Sprintf("failed: %s", err.Error())
			}
			res := utils.PrettyJson(apps)
			return res
		} else {
			app, err := app_manager.GetApp(appName)
			if err != nil {
				return fmt.Sprintf("failed: %s", err.Error())
			}
			res := utils.PrettyJson(app)
			return res
		}
	},
	"start": func(args ...string) (result string) {
		if len(args) == 0 {
			return "empty app, do you mean all?"
		}
		appName := args[0]
		if appName == "all" {
			err := app_manager.StartAll()
			if err != nil {
				return fmt.Sprintf("failed: %s", err.Error())
			}
			return SUCCESS
		} else {
			app, err := app_manager.GetApp(appName)
			if err != nil {
				return fmt.Sprintf("failed: %s", err.Error())
			}
			ok, err := app.Start()
			if err != nil || !ok {
				return fmt.Sprintf("failed: %s", err.Error())
			}
			return SUCCESS
		}
	},
	"stop": func(args ...string) (result string) {
		if len(args) == 0 {
			return "empty app, do you mean all?"
		}
		appName := args[0]
		if appName == "all" {
			err := app_manager.StopAll()
			if err != nil {
				return fmt.Sprintf("failed: %s", err.Error())
			}
			return SUCCESS
		} else {
			app, err := app_manager.GetApp(appName)
			if err != nil {
				return fmt.Sprintf("failed: %s", err.Error())
			}
			ok, err := app.Stop()
			if err != nil || !ok {
				return fmt.Sprintf("failed: %s", err.Error())
			}
			return SUCCESS
		}
	},
	"restart": func(args ...string) (result string) {
		return "restart"
	},
	"status": func(args ...string) (result string) {
		if len(args) == 0 {
			return "empty app, do you mean all?"
		}
		appName := args[0]
		if appName == "all" {
			err := app_manager.StatusAll()
			if err != nil {
				return fmt.Sprintf("failed: %s", err.Error())
			}
			return SUCCESS
		} else {
			app, err := app_manager.GetApp(appName)
			if err != nil {
				return fmt.Sprintf("failed: %s", err.Error())
			}
			ok, err := app.Check()
			if err != nil || !ok {
				return fmt.Sprintf("failed: %s", err.Error())
			}
			return SUCCESS
		}
	},
	"config": func(args ...string) (result string) {
		return config.DirichletConf.ToJSON()
	},
}
