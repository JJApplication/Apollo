/*
Project: dirichlet global_cf.go
Created: 2021/11/20 by Landers
*/

package config

import (
	"reflect"
	"sync"

	"github.com/landers1037/configen"
	"github.com/landers1037/dirichlet/utils"
)

var DirichletConf DConfig

const (
	GlobalConfigRoot = "conf"
	GlobalConfigFile = "dirichlet.pig"
)

// DConfig 全局配置
type DConfig struct {
	// lock
	lock *sync.Mutex `json:"lock,omitempty"`

	ServiceRoot string `json:"service_root"`  // 服务架构的根目录
	APPRoot     string `json:"app_root"`      // 整个服务的根目录
	APPManager  string `json:"app_manager"`   // 管理服务所在目录
	APPCacheDir string `json:"app_cache_dir"` // 服务缓存目录
	APPLogDir   string `json:"app_log_dir"`   // 服务日志目录
	APPTmpDir   string `json:"app_tmp_dir"`   // 服务临时文件目录
	APPBackUp   string `json:"app_back_up"`   // 服务备份目录

	// logger
	Log DLog `json:"log"`

	// DB
	DB DDb `json:"db"`

	// Server
	Server Server `json:"server"`
}

// DLog log config
type DLog struct {
	EnableLog      string `json:"enable_log"`      // yes | no
	EnableStack    string `json:"enable_stack"`    // default disabled
	EnableFunction string `json:"enable_function"` // default disabled
	EnableCaller   string `json:"enable_caller"`   // default enabled
	LogFile        string `json:"log_file"`        // default stderr
	Encoding       string `json:"encoding"`        // default encoding json/console
}

// DDb Database Config
type DDb struct {
	Sqlite Sqlite `json:"sqlite"`
	Mongo  Mongo  `json:"mongo"`
	Redis  Redis  `json:"redis"`
}

type Sqlite struct {
}

type Mongo struct {
	URL    string `json:"url"`
	User   string `json:"user"`
	PassWd string `json:"pass_wd"`
}

type Redis struct {
}

// Server server Config
type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// Sync 从配置文件中同步加载
func (d *DConfig) Sync() {
	if d.lock == nil {
		return
	}

	d.lock.Lock()
	defer d.lock.Unlock()

	_ = configen.ParseConfig(
		&DirichletConf,
		configen.Pig,
		utils.CalDir(
			utils.GetAppDir(),
			GlobalConfigRoot,
			GlobalConfigFile))
}

// Update 安全的更新配置数据
func (d *DConfig) Update(v *DConfig) {
	if d.lock == nil {
		return
	}

	d.lock.Lock()
	vars := reflect.ValueOf(v).Elem()
	globals := reflect.ValueOf(&DirichletConf).Elem()

	for i := 0; i < vars.NumField(); i++ {
		if !vars.Field(i).IsZero() {
			// set pointer
			if globals.Field(i).CanSet() {
				globals.Field(i).Set(vars.Field(i))
			}
		}
	}

	defer d.lock.Unlock()
}

func (d *DConfig) ToJSON() string {
	return utils.PrettyJson(d)
}
