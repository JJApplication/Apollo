/*
Project: Apollo global_cf.go
Created: 2021/11/20 by Landers
*/

package config

import (
	"reflect"
	"sync"

	"github.com/JJApplication/Apollo/utils"
	"github.com/landers1037/configen"
)

var ApolloConf DConfig

const (
	GlobalConfigRoot = "conf"
	GlobalConfigFile = "apollo.pig"
	GlobalScriptRoot = "script"
)

// DConfig 全局配置
type DConfig struct {
	// lock
	lock *sync.Mutex

	Debug       bool   `json:"debug"`         // 开启gin debug
	ServiceRoot string `json:"service_root"`  // 服务架构的根目录
	APPRoot     string `json:"app_root"`      // 整个服务的根目录
	APPManager  string `json:"app_manager"`   // 管理服务所在目录
	APPCacheDir string `json:"app_cache_dir"` // 服务缓存目录
	APPLogDir   string `json:"app_log_dir"`   // 服务日志目录
	APPTmpDir   string `json:"app_tmp_dir"`   // 服务临时文件目录
	APPBackUp   string `json:"app_back_up"`   // 服务备份目录
	APPPidDir   string `json:"app_pid_dir"`   // 服务运行时的pid
	APPBridge   string `json:"app_bridge"`    // JJAPP服务的全局自定义网卡地址
	SSLRoot     string `json:"ssl_root"`      // SSL证书目录
	SSLCert     string `json:"ssl_cert"`      // SSL的证书名称

	// logger
	Log DLog `json:"log"`

	// DB
	DB DDb `json:"db"`

	// Server
	Server Server `json:"server"`

	// CI
	CI CI `json:"ci"`

	// Module 动态注册模块
	Module Module `json:"module"`

	// 管理扩展
	Manager Manager `json:"manager"`

	// 定时任务
	Task Task `json:"task"`

	// GRPC
	GRPC GRPC `json:"grpc"`

	// AES
	AES AES `json:"aes"`

	// 实验特性
	Experiment Experiment `json:"experiment"`
}

// DLog log config
type DLog struct {
	EnableLog      bool   `json:"enable_log"`      // yes | no
	EnableStack    bool   `json:"enable_stack"`    // default disabled
	EnableFunction bool   `json:"enable_function"` // default disabled
	EnableCaller   bool   `json:"enable_caller"`   // default enabled
	LogFile        string `json:"log_file"`        // default stderr
	Encoding       string `json:"encoding"`        // default encoding json/console
}

// DDb Database Config
type DDb struct {
	Sqlite Sqlite `json:"sqlite"`
	Mongo  Mongo  `json:"mongo"`
	Redis  Redis  `json:"redis"`
	KV     string `json:"kv"`
}

type Sqlite struct {
	// todo
}

type Mongo struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	User   string `json:"user"`
	PassWd string `json:"passwd"`
}

type Redis struct {
}

type TiDB struct {
	DB     string `json:"db"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	PassWd string `json:"passwd"`
}

// Server server Config
type Server struct {
	Host        string   `json:"host"`
	Port        int      `json:"port"`
	Uds         string   `json:"uds"`           // Octopus通信的UDS地址
	GRPC        string   `json:"grpc"`          // GRPC的UDS地址
	UICache     bool     `json:"ui_cache"`      // 是否开启缓存
	UICacheTime int      `json:"ui_cache_time"` // 缓存的失效时间
	UIRouter    []string `json:"ui_router"`     // 决定哪些url由前端路由处理
	AuthExpire  int      `json:"auth_expire"`   // 认证的失效时间
	AuthCode    string   `json:"auth_code"`     // 认证码
	Account     string   `json:"account"`       // 认证账户
	PassWd      string   `json:"passwd"`        // 认证密码
	OAuth       OAuth    `json:"oauth"`
}

// CI CI配置
type CI struct {
	DockerHost       string `json:"docker_host"`
	DockerTimeout    int    `json:"docker_timeout"`
	DockerAPIVersion string `json:"docker_api_version"`
}

// Module 动态模块
type Module struct {
	Enable bool `json:"enable"`
}

// Manager 管理外部数据
type Manager struct {
	ManagerNginx struct {
		NginxConf      string `json:"nginx_conf"`
		NginxConfd     string `json:"nginx_confd"`
		NginxMime      string `json:"nginx_mime"`
		NginxAccessLog string `json:"nginx_access_log"`
		NginxErrorLog  string `json:"nginx_error_log"`
		NginxCache     string `json:"nginx_cache"`
	} `json:"manager_nginx"`

	ManagerDatabase struct {
		DatabaseNode string `json:"database_node"`
		ReadOnly     bool   `json:"read_only"`
	} `json:"manager_database"`

	ManagerUds struct {
		UdsDir string `json:"uds_dir"`
	} `json:"manager_uds"`

	ManagerBackup struct {
		// 与appBackupDir重复 暂不实现此配置
	} `json:"manager_backup"`
}

// OAuth Oauth2.0
type OAuth struct {
	ClientID      string   `json:"client_id"`
	ClientSecret  string   `json:"client_secret"`
	AuthorizeList []string `json:"authorize_list"` // 允许的最高管理员github账户
	AllowApiList  []string `json:"allow_api_list"` // 允许操作的api列表采用前缀匹配
}

// Task 定时任务
// 单位s
type Task struct {
	CronJob struct {
		APPBackup string // 全局微服务备份
	} `json:"cron_job"`

	BackgroundJob struct {
		DBSave              int // 刷新运行时数据到数据库中
		DBPersist           int // 持久化数据库数据到bson格式的备份文件
		AppSync             int // 根据注册的服务名同步模型文件的更新数据到Apollo
		AppRuntimeSync      int // 运行时端口等数据同步到数据库
		AppCheck            int // 状态检查
		LogRotate           int // 日志转储
		NoEngineRuntimeSync int // 运行时生成的noengine端口信息同步到noengine.map
	} `json:"background_job"`

	AutoDiscover struct {
		App      int // 从模型文件中全量同步微服务
		NoEngine int // 从NoEngine模型中加载NoEngine服务
	} `json:"auto_discover"`
}

type GRPC struct {
	Host       string            `json:"host"`
	Port       int               `json:"port"`
	MaxAttempt int               `json:"max_attempt"`
	UdsAddr    map[string]string `json:"uds_addr"`
}

func (g *GRPC) GetAddr(name string) string {
	if g.UdsAddr == nil {
		return ""
	}
	val, ok := g.UdsAddr[name]
	if ok {
		return val
	}
	return ""
}

type AES struct {
	Key string `json:"key"`
}

// Experiment 实验特性
type Experiment struct {
	PortV2   bool `json:"port_v2"` // 端口管理器v2
	TaskData struct {
		Path     string `json:"path"`
		Duration int    `json:"duration"`
	} `json:"task_data"` // 任务持久化配置
}

// Sync 从配置文件中同步加载
func (d *DConfig) Sync() {
	if d.lock == nil {
		return
	}

	d.lock.Lock()
	defer d.lock.Unlock()

	_ = configen.ParseConfig(
		&ApolloConf,
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
	globals := reflect.ValueOf(&ApolloConf).Elem()

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
