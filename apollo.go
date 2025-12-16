package main

import (
	"runtime"

	"github.com/JJApplication/Apollo/cmd"

	"github.com/JJApplication/Apollo/logger"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

// @title Apollo ServiceGroup
// @version 1.0
// @description Apollo服务接口文档
// @termsOfService http://service.renj.io/terms
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath http://service.renj.io/api
func main() {
	// 监听信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)

	// 如果是子进程启动，继承文件描述符（如果需要）
	if os.Getenv("APOLLO_GRACEFUL_RESTART") == "true" {
		logger.LoggerSugar.Info("Apollo started gracefully")
	}

	cmd.ApolloCmd()
	initGlobalConfig()
	initDataDir()
	logger.Check()
	err := logger.InitLogger()
	if err != nil {
		logger.Logger.Error(logger.LoggerInitFailed)
		return
	}
	logger.CoreDump()
	// init recover
	defer func() {
		if r := recover(); r != nil {
			logger.LoggerSugar.Error(r)
			// add stack
			var stackBuf [1 << 16]byte
			stackLen := runtime.Stack(stackBuf[:], false)
			logger.LoggerSugar.Errorf("%s", stackBuf[:stackLen])
		}
	}()

	// init database
	initMongo()
	// init KV
	initKV()
	// init app manager
	initAPPManager()
	// init app discover
	initDiscoverManager()
	// init log manager
	initLogManager()
	// init task manager
	initTaskManager()
	// init background ticker
	initBackgroundJobs()
	// init background cron job
	initCronJobs()
	// init docker
	initDockerClient()
	// init NoEngine
	initNoEngineApps()
	//init env manager etcd client
	initEnvManager()
	// init repo manager
	initRepoManager()
	// init port manager
	initPortManager()
	// init indicator manager
	initIndicatorManager()
	// init database manager
	initDatabaseManager()
	// init process manager
	initProcessManager()
	// init Secure manager
	initSecureManager()
	// init gRPC server
	initGRPCServer()
	// only all manager init, can uds server be active
	initUDS()
	// init engine
	initEngine()

	// 阻塞并等待信号
	handleSignals(signals)
}

func handleSignals(c chan os.Signal) {
	for {
		s := <-c
		switch s {
		case syscall.SIGHUP:
			logger.LoggerSugar.Info("Received SIGHUP, reloading...")
			err := forkAndExec()
			if err != nil {
				logger.LoggerSugar.Errorf("ForkExec failed: %v", err)
			}
		case syscall.SIGTERM, syscall.SIGINT:
			logger.LoggerSugar.Info("Received stop signal, shutting down...")
			// 执行清理操作
			os.Exit(0)
		}
	}
}

func forkAndExec() error {
	argv0, err := exec.LookPath(os.Args[0])
	if err != nil {
		return err
	}
	
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	files := []*os.File{
		os.Stdin,
		os.Stdout,
		os.Stderr,
	}

	// 设置环境变量标记
	env := append(os.Environ(), "APOLLO_GRACEFUL_RESTART=true")

	p, err := os.StartProcess(argv0, os.Args, &os.ProcAttr{
		Dir:   wd,
		Env:   env,
		Files: files,
		Sys:   &syscall.SysProcAttr{},
	})
	
	if err != nil {
		return err
	}
	
	logger.LoggerSugar.Infof("Forked new process PID: %d", p.Pid)
	
	// 父进程退出
	os.Exit(0)
	return nil
}
