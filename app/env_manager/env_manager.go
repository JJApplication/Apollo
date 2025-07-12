package env_manager

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/grpc/nidavellir"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 环境变量管理实例

const EnvManagerPrefix = "[Env Manager]"

var manager *EnvManager

type EnvManager struct {
	init     bool // 是否初始化
	grpcConn *grpc.ClientConn
	client   nidavellir.ConfigServiceClient
}

func InitEnvManager() {
	go func() {
		manager = new(EnvManager)
		manager.Init()
	}()
}

func GetEnvManager() *EnvManager {
	if manager == nil || !manager.init {
		sync.OnceFunc(func() {
			manager = new(EnvManager)
			manager.Init()
		})
		return manager
	}
	return manager
}

func (m *EnvManager) Init() {
	logger.LoggerSugar.Infof("%s start to init grpc client at %s", EnvManagerPrefix, config.ApolloConf.GRPC.UdsAddr)
	client, err := grpc.NewClient(
		config.ApolloConf.GRPC.UdsAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithMaxCallAttempts(5))
	if err != nil {
		m.init = false
		logger.LoggerSugar.Errorf("%s grpc dial err %s", EnvManagerPrefix, err.Error())
		return
	}
	m.client = nidavellir.NewConfigServiceClient(client)
	m.grpcConn = client
	m.init = true
	logger.LoggerSugar.Infof("%s grpc client init success at: %s", EnvManagerPrefix, config.ApolloConf.GRPC.UdsAddr)
}

func (m *EnvManager) Close() {
	if m.grpcConn != nil {
		if err := m.grpcConn.Close(); err != nil {
			logger.LoggerSugar.Errorf("%s grpc conn close err %s", EnvManagerPrefix, err.Error())
		}
	}
}

func (m *EnvManager) IsInit() bool {
	return m.init
}

// GetEnvs 获取用于设置的环境变量数组
func (m *EnvManager) GetEnvs(service string) []string {
	items, err := m.GetServiceConfigs(service)
	if err != nil {
		return nil
	}
	data := make([]string, 0, len(items))
	for _, item := range items {
		if item.GetEncrypt() {
			decryptVal, err := utils.DecryptAES256(item.Value, config.ApolloConf.AES.Key)
			if err != nil {
				logger.LoggerSugar.Errorf("%s decrypt [%s]-[%s] err %s", EnvManagerPrefix, service, item.Key, err.Error())
				continue
			}
			data = append(data, fmt.Sprintf("%s=%s", item.Key, decryptVal))
		} else {
			data = append(data, fmt.Sprintf("%s=%s", item.Key, item.Value))
		}
	}
	return data
}

// ListAllServices 列出全部微服务
func (m *EnvManager) ListAllServices() []string {
	if !m.IsInit() {
		return nil
	}
	ctx := context.Background()
	data, err := m.client.ListServices(ctx, nil)
	if err != nil {
		logger.LoggerSugar.Errorf("%s list services err %s", EnvManagerPrefix, err.Error())
		return nil
	}

	return data.Services
}

func (m *EnvManager) SetConfig(service, key, val string, encrypt bool) error {
	if !m.IsInit() {
		return errors.New("not init")
	}
	if service == "" || key == "" {
		return errors.New("invalid param")
	}
	if encrypt {
		encryptData, err := utils.EncryptAES256(val, config.ApolloConf.AES.Key)
		if err != nil {
			return err
		}
		ctx := context.Background()
		_, err = m.client.SetConfig(ctx, &nidavellir.SetConfigRequest{
			ServiceName: service,
			Key:         key,
			Value:       encryptData,
			Encrypt:     encrypt,
		})
		return err
	}
	ctx := context.Background()
	_, err := m.client.SetConfig(ctx, &nidavellir.SetConfigRequest{
		ServiceName: service,
		Key:         key,
		Value:       val,
		Encrypt:     encrypt,
	})
	return err
}

func (m *EnvManager) GetConfig(service, key string) (*nidavellir.ConfigItem, error) {
	if !m.IsInit() {
		return nil, errors.New("not init")
	}
	if service == "" || key == "" {
		return nil, errors.New("invalid param")
	}
	ctx := context.Background()
	data, err := m.client.GetConfig(ctx, &nidavellir.GetConfigRequest{
		ServiceName: service,
		Key:         key,
	})
	if err != nil {
		return nil, err
	}
	return data.GetConfig(), nil
}

func (m *EnvManager) GetServiceConfigs(service string) (map[string]*nidavellir.ConfigItem, error) {
	if !m.IsInit() {
		return nil, errors.New("not init")
	}
	ctx := context.Background()
	data, err := m.client.GetServiceConfigs(ctx, &nidavellir.GetServiceConfigsRequest{
		ServiceName: service,
	})
	if err != nil {
		return nil, err
	}
	return data.GetConfigs(), nil
}

func (m *EnvManager) DeleteConfig(service, key string) error {
	if !m.IsInit() {
		return errors.New("not init")
	}
	if service == "" || key == "" {
		return errors.New("invalid param")
	}
	ctx := context.Background()
	_, err := m.client.DeleteConfig(ctx, &nidavellir.DeleteConfigRequest{
		ServiceName: service,
		Key:         key,
	})
	return err
}

func (m *EnvManager) DeleteServiceConfigs(service string) (*nidavellir.DeleteServiceConfigsResponse, error) {
	if !m.IsInit() {
		return nil, errors.New("not init")
	}
	ctx := context.Background()
	_, err := m.client.DeleteServiceConfigs(ctx, &nidavellir.DeleteServiceConfigsRequest{
		ServiceName: service,
	})
	return nil, err
}
