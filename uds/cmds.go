/*
Project: Apollo cmds.go
Created: 2021/12/9 by Landers
*/

package uds

import (
	"fmt"
	"strings"

	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/utils"
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

var cmdsMap = map[string]func(args ...string) (result, err string){
	"show": func(args ...string) (result, err string) {
		var s []string
		for _, cmd := range cmds {
			s = append(s, fmt.Sprintf("[%-7s] %s\n", cmd, helpMsg[cmd]))
		}

		return strings.Join(s, ""), ""
	},
	"app": func(args ...string) (result, err string) {
		if len(args) == 0 {
			return "empty app, do you mean all?", ""
		}
		appName := args[0]
		switch appName {
		case "help":
			return "list\nall\nsync\nsyncdb\n[app]\n", ""
		case "list":
			var appList []string
			app_manager.APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
				appList = append(appList, key.(string))
				return true
			})
			return strings.Join(appList, "\n"), ""
		case "all":
			apps, err := app_manager.GetAllApp()
			if err != nil {
				return "", fmt.Sprintf("failed: %s", err.Error())
			}
			res := utils.PrettyJson(apps)
			return res, ""
		case "sync":
			var res []string
			app_manager.APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
				app := value.(app_manager.App)
				ok, _ := app.Sync()
				if !ok {
					res = append(res, fmt.Sprintf("[%s]: BAD", app.Meta.Name))
				} else {
					res = append(res, fmt.Sprintf("[%s]: OK", app.Meta.Name))
				}
				return true
			})
			return "sync to config\n" + strings.Join(res, "\n"), ""
		case "syncdb":
			app_manager.SaveToDB()
			return "sync to db", ""
		default:
			app, err := app_manager.GetApp(appName)
			if err != nil {
				return "", fmt.Sprintf("failed: %s", err.Error())
			}
			res := utils.PrettyJson(app)
			return res, ""
		}
	},
	"start": func(args ...string) (result, err string) {
		if len(args) == 0 {
			return "empty app, do you mean all?", ""
		}
		appName := args[0]
		switch appName {
		case "all":
			res, err := app_manager.StartAll()
			if err != nil {
				return strings.Join(res, "\n"), fmt.Sprintf("failed: %s", err.Error())
			}
			return strings.Join(res, "\n"), ""
		default:
			app, err := app_manager.GetApp(appName)
			if err != nil {
				return "", fmt.Sprintf("failed: %s", err.Error())
			}
			ok, err := app.Start()
			if err != nil || !ok {
				return "", fmt.Sprintf("failed: %s", err.Error())
			}
			return SUCCESS, ""
		}
	},
	"stop": func(args ...string) (result, err string) {
		if len(args) == 0 {
			return "empty app, do you mean all?", ""
		}
		appName := args[0]
		switch appName {
		case "all":
			res, err := app_manager.StopAll()
			if err != nil {
				return strings.Join(res, "\n"), fmt.Sprintf("failed: %s", err.Error())
			}
			return strings.Join(res, "\n"), ""
		default:
			app, err := app_manager.GetApp(appName)
			if err != nil {
				return "", fmt.Sprintf("failed: %s", err.Error())
			}
			ok, err := app.Stop()
			if err != nil || !ok {
				return "", fmt.Sprintf("failed: %s", err.Error())
			}
			return SUCCESS, ""
		}
	},
	"restart": func(args ...string) (result, err string) {
		return "restart", ""
	},
	"status": func(args ...string) (result, err string) {
		if len(args) == 0 {
			return "empty app, do you mean all?", ""
		}
		appName := args[0]
		switch appName {
		case "all":
			res, err := app_manager.StatusAll()
			if err != nil {
				return "", fmt.Sprintf("failed: %s\n%s", err.Error(), strings.Join(res, "\n"))
			}
			return fmt.Sprintf("%s\n%s", SUCCESS, strings.Join(res, "\n")), ""
		default:
			app, err := app_manager.GetApp(appName)
			if err != nil {
				return "", fmt.Sprintf("failed: %s", err.Error())
			}
			ok, err := app.Check()
			if err != nil || !ok {
				return "", fmt.Sprintf("failed: %s", err.Error())
			}
			return SUCCESS, ""
		}
	},
	"config": func(args ...string) (result, err string) {
		return config.ApolloConf.ToJSON(), ""
	},
}
