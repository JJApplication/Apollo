/*
   Create: 2024/1/10
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package noengine_manager

import (
	"github.com/JJApplication/Apollo/app/docker_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"io/fs"
	"path/filepath"
	"sync"
)

var (
	NoEngineManager      = "[NoEngineManager]"
	NoEngine             = "NoEngine"
	NoEngineAPP          = ""              // NoEngine根目录
	NoEngineAPPConf      = "noengine.conf" // 通用的NoNngine http模板
	NoEngineTemplateConf = "noengine.cf.json"
	NoEngineAPPMap       = "noengine.app.json" // 存储服务端口映射的文件
	NoEngineInitFlag     bool                  // 是否为初次加载
)

var (
	NoEngineMap   *sync.Map // 基于配置文件加载的初始配置（用于操作容器|只读）
	NoEngineMapRt *sync.Map // 运行过程中产生的配置（在每次操作容器时被重新赋值）
)

func InitNoEngineManager() {
	NoEngineMap = new(sync.Map)
	NoEngineMapRt = new(sync.Map)
	NoEngineAPP = filepath.Join(config.ApolloConf.APPRoot, NoEngine)                         // NoEngine根目录
	NoEngineAPPConf = filepath.Join(config.ApolloConf.APPRoot, NoEngine, "noengine.conf")    // 通用的NoNngine http模板
	NoEngineAPPMap = filepath.Join(config.ApolloConf.APPRoot, NoEngine, "noengine.app.json") // 存储服务端口映射的文件
	LoadAllNoEngineAPPs()
	NoEngineInitFlag = true
}

// GetNoEngineAPPDir 获取微服务NoEngine根目录
func GetNoEngineAPPDir(app string) string {
	return filepath.Join(NoEngineAPP, app)
}

// GetNoEngineAPPTempCf 获取微服务NoEngine配置文件
func GetNoEngineAPPTempCf(app string) string {
	return filepath.Join(NoEngineAPP, app, NoEngineTemplateConf)
}

// LoadAllNoEngineAPPs 加载所有NoEngine服务到缓存中
// 需要实时刷新的是RT缓存，初始配置仅在首次加载时加载
func LoadAllNoEngineAPPs() {
	engineMap, rterr := ReloadNoEngineMap()
	if !NoEngineInitFlag {
		// 从$NoEngineAPP下的全部目录寻找$NoEngineTemplateConf文件解析
		if err := filepath.Walk(NoEngineAPP, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// 为文件时是否为$NoEngineTemplateConf
			// 为Dir时进入
			if info.Name() == NoEngineTemplateConf {
				var data NoEngineTemplate
				if err = utils.ParseJsonFile(path, &data); err != nil {
					logger.LoggerSugar.Errorf("%s load NoEngineAPP -> %s error: %s", NoEngineManager, path, err.Error())
				} else {
					NoEngineMap.Store(data.ServerName, data)
				}
			}
			return err
		}); err != nil {
			logger.LoggerSugar.Errorf("%s init NoEngineAPPs error: %s", NoEngineManager, err.Error())
		}
	}
	// 不存在缓存即为初始加载
	if rterr != nil {
		logger.LoggerSugar.Warnf("%s reload NoEngineAPPs from cache error: %s", NoEngineManager, rterr.Error())
	} else {
		// 更新运行时配置
		for app, val := range engineMap {
			NoEngineMapRt.Store(app, val)
		}
		logger.LoggerSugar.Infof("%s reload NoEngineAPPs from cache success", NoEngineManager)
	}

	// 从本地目录下更新初始配置（自动发现场景)
	if err := filepath.Walk(NoEngineAPP, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == NoEngineTemplateConf {
			var data NoEngineTemplate
			if err = utils.ParseJsonFile(path, &data); err != nil {
				logger.LoggerSugar.Errorf("%s load NoEngineAPP -> %s error: %s", NoEngineManager, path, err.Error())
			} else {
				if _, ok := NoEngineMap.Load(data.ServerName); !ok {
					NoEngineMap.Store(data.ServerName, data)
				}
			}
		}
		return err
	}); err != nil {
		logger.LoggerSugar.Errorf("%s init NoEngineAPPs error: %s", NoEngineManager, err.Error())
	}
	logger.LoggerSugar.Infof("%s init NoEngineAPPs success", NoEngineManager)
}

// GetAllNoEngineAPPs 获取全部NoEngineApp(初始配置)
func GetAllNoEngineAPPs() map[string]NoEngineTemplate {
	var res = make(map[string]NoEngineTemplate)
	NoEngineMap.Range(func(key, value any) bool {
		if key.(string) == "" {
			return false
		}
		res[key.(string)] = value.(NoEngineTemplate)
		return true
	})

	return res
}

// GetAllNoEngineAPPsRt 获取当前运行时配置文件(运行时配置)
func GetAllNoEngineAPPsRt() map[string]NoEngineTemplate {
	var res = make(map[string]NoEngineTemplate)
	NoEngineMapRt.Range(func(key, value any) bool {
		if key.(string) == "" {
			return false
		}
		res[key.(string)] = value.(NoEngineTemplate)
		return true
	})

	return res
}

// GetNoEngineAPP 获取指定NoEngineApp(初始配置)
func GetNoEngineAPP(app string) NoEngineTemplate {
	if !HasNoEngineApp(app) {
		return NoEngineTemplate{}
	}
	val, _ := NoEngineMap.Load(app)
	return val.(NoEngineTemplate)
}

// GetNoEngineAPPRt 获取指定NoEngineApp(运行时)
func GetNoEngineAPPRt(app string) NoEngineTemplate {
	if !HasNoEngineAppRt(app) {
		return NoEngineTemplate{}
	}
	val, _ := NoEngineMapRt.Load(app)
	return val.(NoEngineTemplate)
}

// HasNoEngineApp 是否存在此APP
func HasNoEngineApp(app string) bool {
	_, ok := NoEngineMap.Load(app)
	return ok
}

// HasNoEngineAppRt 是否存在此APP(运行时)
func HasNoEngineAppRt(app string) bool {
	_, ok := NoEngineMapRt.Load(app)
	return ok
}

// StartNoEngineApp 每次启动都是全新启动，会清除容器内置的缓存 重置随机端口
func StartNoEngineApp(app string) error {
	if app == "" {
		return nil
	}
	// 已经启动的容器跳过
	if NoEngineAPPID(app) != "" {
		return docker_manager.ContainerStart(NoEngineAPPID(app))
	}
	temp := GetNoEngineAPP(app)
	if temp.ServerName == "" {
		return nil
	}
	err, tempInit := createContainer(temp)
	if err != nil {
		logger.LoggerSugar.Errorf("%s start NoEngineAPP -> %s error: %s", NoEngineManager, app, err.Error())
	}
	// 更新运行时配置
	NoEngineMapRt.Store(app, tempInit)
	syncMap()
	go RefreshNoEngineMap()
	return err
}

// StopNoEngineApp 停止NoEngine服务 不会删除容器
// 不使用--rm参数启动容器 否则停止会自动删除
func StopNoEngineApp(app string) error {
	// 不存在的容器跳过
	if NoEngineAPPID(app) == "" {
		return nil
	}
	return docker_manager.ContainerStop(NoEngineAPPID(app))
}

// RemoveNoEngineApp 删除容器 同时清理Map
func RemoveNoEngineApp(app string) error {
	// 优先清理Map
	// 初始配置为模板 无需删除
	NoEngineMapRt.Delete(app)
	syncMap()
	// 不存在的容器跳过
	if NoEngineAPPID(app) == "" {
		return nil
	}
	if err := docker_manager.ContainerRemove(NoEngineAPPID(app)); err != nil {
		return err
	}
	go RefreshNoEngineMap()
	return nil
}

// PauseNoEngineApp 暂停NoEngine服务
func PauseNoEngineApp(app string) error {
	// 不存在的容器跳过
	if NoEngineAPPID(app) == "" {
		return nil
	}
	return docker_manager.ContainerPause(NoEngineAPPID(app))
}

// ResumeNoEngineApp 恢复NoEngine服务
func ResumeNoEngineApp(app string) error {
	// 不存在的容器跳过
	if NoEngineAPPID(app) == "" {
		return nil
	}
	return docker_manager.ContainerResume(NoEngineAPPID(app))
}

// StatusNoEngineApp 查看NoEngine服务状态
// 根据容器的状态只有Running Ready Pause Exit NoExist UnKnown
func StatusNoEngineApp(app string) string {
	info, err := docker_manager.ContainerInfo(NoEngineAPPID(app))
	if err != nil {
		return "unknown"
	}
	return info.State.Status
}

func NoEngineAPPID(app string) string {
	return docker_manager.GetContainerIDByName(app)
}

// IsNoEngineAppCreated 是否存在此容器
func IsNoEngineAppCreated(app string) bool {
	return NoEngineAPPID(app) != ""
}
