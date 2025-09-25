package port_manager

import (
	"context"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/grpc/themis"
	"github.com/JJApplication/Apollo/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
)

const (
	PortManagerPrefix = "[Port Manager]"
)

var manager *PortManager

type PortManager struct {
	init     bool // 是否初始化
	grpcConn *grpc.ClientConn
	client   themis.PortServiceClient
}

func InitPortManager() {
	go func() {
		manager = new(PortManager)
		manager.Init()
	}()
}

func GetPortManager() *PortManager {
	if manager == nil || !manager.init {
		sync.OnceFunc(func() {
			manager = new(PortManager)
			manager.Init()
		})
		return manager
	}
	return manager
}

func (m *PortManager) Init() {
	// 未开启特性时跳过
	if !config.ApolloConf.Experiment.PortV2 {
		return
	}
	logger.LoggerSugar.Infof("%s start to init grpc client at %s", PortManagerPrefix, config.ApolloConf.GRPC.GetAddr("Nidavellir"))
	client, err := grpc.NewClient(
		config.ApolloConf.GRPC.GetAddr("Themis"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithMaxCallAttempts(config.ApolloConf.GRPC.MaxAttempt))
	if err != nil {
		m.init = false
		logger.LoggerSugar.Errorf("%s grpc dial err %s", PortManagerPrefix, err.Error())
		return
	}
	m.client = themis.NewPortServiceClient(client)
	m.grpcConn = client
	m.init = true
	logger.LoggerSugar.Infof("%s grpc client init success at: %s", PortManagerPrefix, config.ApolloConf.GRPC.UdsAddr)
}

func (m *PortManager) Close() {
	if m.grpcConn != nil {
		if err := m.grpcConn.Close(); err != nil {
			logger.LoggerSugar.Errorf("%s grpc conn close err %s", PortManagerPrefix, err.Error())
		}
	}
}

func (m *PortManager) IsInit() bool {
	if !config.ApolloConf.Experiment.PortV2 {
		return false
	}
	return m.init
}

// GetAll TODO 获取所有注册的端口
func (m *PortManager) GetAll() {
	if !m.init {
		return
	}
	return
}

// NewPort 获取新的随机端口
func (m *PortManager) NewPort() int {
	response, err := m.client.GetRandomPort(context.Background(), &themis.GetRandomPortRequest{})
	if err != nil {
		logger.LoggerSugar.Errorf("%s get random port err %s", PortManagerPrefix, err.Error())
		return 0
	}
	if response.Error != "" {
		logger.LoggerSugar.Errorf("%s new port err %s", PortManagerPrefix, response.Error)
		return 0
	}
	return int(response.Port)
}

// GetAppPort 获取应用的端口
func (m *PortManager) GetAppPort(app string) int {
	response, err := m.client.GetAppPort(context.Background(), &themis.GetAppPortRequest{AppName: app})
	if err != nil {
		logger.LoggerSugar.Errorf("%s get app: %s port err %s", PortManagerPrefix, app, err.Error())
		return 0
	}
	if response.Error != "" {
		logger.LoggerSugar.Errorf("%s get app: %s port err %s", PortManagerPrefix, app, response.Error)
		return 0
	}
	return int(response.Port)
}

// SetAppPort 设置应用的端口
func (m *PortManager) SetAppPort(app string, port int) error {
	response, err := m.client.SetAppPort(context.Background(), &themis.SetAppPortRequest{
		AppName: app,
		Port:    int32(port),
	})
	if err != nil {
		logger.LoggerSugar.Errorf("%s set app: %s port err %s", PortManagerPrefix, app, err.Error())
		return err
	}
	if response.Error != "" {
		logger.LoggerSugar.Errorf("%s set app: %s port err %s", PortManagerPrefix, app, response.Error)
		return err
	}
	return nil
}
