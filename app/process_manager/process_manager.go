package process_manager

import (
	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/kv"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"github.com/JJApplication/Apollo/utils/json"
	"go.etcd.io/bbolt"
	"sync"
	"time"
)

var (
	processManager       *ProcessManager
	ProcessManagerPrefix = "[Process Manager]"
)

type ProcessManager struct {
	pmLock sync.RWMutex
	pm     map[string]utils.SysProc // 上一次缓存的数据
}

// 时序数据

func Get() *ProcessManager {
	if processManager == nil {
		return &ProcessManager{pm: make(map[string]utils.SysProc)}
	}
	return processManager
}

func InitProcessManager() {
	processManager = &ProcessManager{
		pm:     make(map[string]utils.SysProc, 10),
		pmLock: sync.RWMutex{},
	}

	go func() {
		time.Sleep(time.Second * 5)
		processManager.Start()
	}()
}

// Add 刷新缓存 同时增加时序数据
func (p *ProcessManager) Add(app string, proc utils.SysProc) {
	kv.PreHookAdd(kv.KV, app)
	p.pm[app] = proc

	if err := kv.AddIfNotExist(app, proc.ToBytes()); err != nil {
		logger.LoggerSugar.Errorf("%s add app %s to kv failed, error: %s", ProcessManagerPrefix, app, err.Error())
	} else {
		logger.LoggerSugar.Infof("%s add app %s to kv success", ProcessManagerPrefix, app)
	}
}

func (p *ProcessManager) GetProc(app string) utils.SysProc {
	p.pmLock.RLock()
	if proc, ok := p.pm[app]; ok {
		p.pmLock.RUnlock()
		return proc
	}
	p.pmLock.RUnlock()
	// 不存在则实时获取并存储
	proc := utils.FilterProcess(app)
	if proc == nil {
		return utils.SysProc{}
	}
	data := utils.SysProc{
		PID:            utils.GetProcessPID(proc),
		CPUPercent:     utils.CalcProcessCpu(proc),
		ProcessMemInfo: utils.CalcProcessMem(proc),
		ProcessIO:      utils.CalcProcessIO(proc),
		NetConnections: utils.CalcProcessNet(proc),
		Threads:        utils.GetProcessThreads(proc),
	}
	p.pmLock.Lock()
	p.pm[app] = data
	p.pmLock.Unlock()
	return data
}

// Start 开启定时获取时序数据
func (p *ProcessManager) Start() {
	apps, _ := app_manager.GetAllAppName()
	if len(apps) == 0 {
		return
	}
	go func() {
		ticker := time.NewTicker(time.Hour * 1)
		for {
			select {
			case <-ticker.C:
				logger.LoggerSugar.Infof("%s start to fetch app process: %d", ProcessManagerPrefix, len(apps))
				for _, app := range apps {
					proc := utils.FilterProcess(app)
					if proc == nil {
						continue
					}
					logger.LoggerSugar.Infof("%s start to fetch app: %s", ProcessManagerPrefix, app)
					data := utils.SysProc{
						PID:            utils.GetProcessPID(proc),
						CPUPercent:     utils.CalcProcessCpu(proc),
						ProcessMemInfo: utils.CalcProcessMem(proc),
						ProcessIO:      utils.CalcProcessIO(proc),
						NetConnections: utils.CalcProcessNet(proc),
						Threads:        utils.GetProcessThreads(proc),
					}
					p.pmLock.Lock()
					p.pm[app] = data
					p.Add(app, data)
					p.pmLock.Unlock()
				}
			}
		}
	}()
}

// GetProcHistory 获取某个进程的历史数据
func (p *ProcessManager) GetProcHistory(app string) []utils.SysProc {
	var data []utils.SysProc
	err := kv.KV.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(app)).ForEach(func(k, v []byte) error {
			var s utils.SysProc
			err := json.Unmarshal(v, &s)
			if err != nil {
				return err
			}
			data = append(data, s)
			return nil
		})
	})
	if err != nil {
		return nil
	}

	return data
}

// GetProcList 获取当前进程的信息
func (p *ProcessManager) GetProcList() map[string]utils.SysProc {
	return p.pm
}
