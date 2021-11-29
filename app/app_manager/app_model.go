/*
Project: dirichlet app_model.go
Created: 2021/11/20 by Landers
*/

package app_manager

import (
	"github.com/gookit/validate"
	"github.com/landers1037/dirichlet/logger"
)

// App model for app
type App struct {
	Name          string `json:"name" validate:"required"`
	ID            string `json:"id" validate:"required"`
	Type          string `json:"type"`           // service | middleware
	ReleaseStatus string `json:"release_status"` // published | pending | testing
	EngDes        string `json:"eng_des"`
	CHSDes        string `json:"chs_des"`

	// 管理项
	ManageCMD CMD `json:"manage_cmd"`
	// 元数据
	Meta Meta `json:"meta"`
	// 动态依赖配置
	RunData RunData `json:"run_data"`
}

// CMD 服务的管理脚本
type CMD struct {
	Start     string `json:"start"`
	Stop      string `json:"stop"`
	Restart   string `json:"restart"`
	ForceKill string `json:"force_kill"`
	Check     string `json:"check"`
}

// Meta 服务的元数据
type Meta struct {
	Author      string   `json:"author"`
	Language    []string `json:"language"`
	CreateDate  string   `json:"create_date"`
	Version     string   `json:"version"`
	DynamicConf bool     `json:"dynamic_conf"` // 是否需要生成配置文件
	ConfType    string   `json:"conf_type"`    // nginx | gunicorn
	ConfPath    string   `json:"conf_path"`    // 支持绝对和相对路径
}

// RunData 运行时依赖
type RunData struct {
	Envs       []string `json:"envs"` // just like `Name=Diri`
	Ports      []int    `json:"ports"`
	RandomPort bool     `json:"random_port"` // if using random port
	Host       string   `json:"host"`        // always must be localhost
}

// Validate 适用于model的检查器
func (app *App) Validate() bool {
	if !validate.Struct(app).Validate() {
		logger.Logger.Error(appCodeMsg("app validate failed") + validate.Struct(app).Errors.String())
		return false
	}

	return true
}
