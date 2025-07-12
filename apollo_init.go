/*
Project: Apollo apollo_init.go
Created: 2021/11/28 by Landers
*/

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JJApplication/Apollo/app/env_manager"
	"github.com/JJApplication/Apollo/app/log_manager"
	"github.com/JJApplication/Apollo/app/noengine_manager"
	"github.com/JJApplication/Apollo/router/router_auth"
	"github.com/JJApplication/Apollo/router/router_env"
	"github.com/JJApplication/Apollo/router/router_log"
	"github.com/JJApplication/Apollo/router/router_noengine"
	"github.com/JJApplication/Apollo/router/router_oauth"
	"github.com/JJApplication/Apollo/router/router_script"
	"github.com/JJApplication/Apollo/router/router_status"
	"github.com/JJApplication/Apollo/router/router_system"

	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/app/discover_manager"
	"github.com/JJApplication/Apollo/app/docker_manager"
	"github.com/JJApplication/Apollo/app/task_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/cron"
	"github.com/JJApplication/Apollo/database"
	"github.com/JJApplication/Apollo/engine"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/router/router_alarm"
	"github.com/JJApplication/Apollo/router/router_app"
	"github.com/JJApplication/Apollo/router/router_container"
	"github.com/JJApplication/Apollo/router/router_modules"
	"github.com/JJApplication/Apollo/router/router_tasks"
	"github.com/JJApplication/Apollo/router/router_web"
	"github.com/JJApplication/Apollo/uds"
	"github.com/JJApplication/Apollo/utils"
)

const (
	APPName = "[Apollo]"
)

// 初始化运行时数据
func initGlobalConfig() {
	t := time.Now()
	err := config.InitGlobalConfig()
	if err != nil {
		fmt.Printf("%s init config failed: %s\n", APPName, err)
		return
	}
	fmt.Printf("%s load config in %dms\n", APPName, utils.TimeCalcUnix(t))
}

// 初始化数据库
func initMongo() {
	err := database.InitDBMongo()
	if err != nil {
		fmt.Printf("%s init mongo failed: %s\n", APPName, err.Error())
		database.MongoPing = false
		return
	}
	if err = database.Ping(); err != nil {
		fmt.Printf("%s ping mongo failed: %s\n", APPName, err.Error())
		database.MongoPing = false
		return
	}
	fmt.Println(APPName + " init mongo success")
}

func initAPPManager() {
	app_manager.InitAPPManager()
	app_manager.SaveToDB()
	app_manager.FirstLoad()
	logger.LoggerSugar.Infof("%s init APPManager done", APPName)
}

func initDiscoverManager() {
	discover_manager.InitDiscoverManager()
	logger.LoggerSugar.Infof("%s init APPDiscover done", APPName)
}

// 任务管理初始化
func initTaskManager() {
	task_manager.InitTaskManager()
}

func initBackgroundJobs() {
	cron.InitBackgroundJobs()
	logger.LoggerSugar.Infof("%s init BackgroundJobs done", APPName)
}

func initCronJobs() {
	cron.InitCronJobs()
	logger.LoggerSugar.Infof("%s init CronJobs done", APPName)
}

func initNoEngineApps() {
	noengine_manager.InitNoEngineManager()
}

func initEngine() {
	apolloEngine := engine.NewEngine(&engine.EngineConfig{
		Host: config.ApolloConf.Server.Host,
		Port: config.ApolloConf.Server.Port,
	})
	apolloEngine.Init()

	// load router
	r := apolloEngine.GetEngine()
	router_app.Init(r)
	router_web.Init(r)
	router_tasks.Init(r)
	router_container.Init(r)
	router_alarm.Init(r)
	router_modules.Init(r)
	router_status.Init(r)
	router_system.Init(r)
	router_auth.Init(r)
	router_script.Init(r)
	router_log.Init(r)
	router_noengine.Init(r)
	router_oauth.Init(r)
	router_env.Init(r)

	// hooks engine
	engine.Hooks(apolloEngine)

	logger.LoggerSugar.Infof("%s init Engine done, start at %s", APPName, fmt.Sprintf("%s:%d", config.ApolloConf.Server.Host, config.ApolloConf.Server.Port))
	err := apolloEngine.RunServer()
	if err != nil && err != http.ErrServerClosed {
		logger.LoggerSugar.Errorf("%s server start failed", APPName)
		logger.LoggerSugar.Errorf("%s %s", APPName, err.Error())
		cron.InsureTickerExit()
		fmt.Println(APPName + " server start failed ⚠️")
		env_manager.GetEnvManager().Close()
		return
	}

	if err == http.ErrServerClosed {
		env_manager.GetEnvManager().Close()
		fmt.Println(APPName + " thanks for using 💕️")
	}
}

func initUDS() {
	uds.Listen()
}

func initDockerClient() {
	docker_manager.InitDockerClient()
}

func initLogManager() {
	log_manager.InitLogManager()
}

func initEnvManager() {
	env_manager.InitEnvManager()
}
