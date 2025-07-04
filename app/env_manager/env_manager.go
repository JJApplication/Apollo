package env_manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/grpc/nidavellir"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
)

// 环境变量管理实例

var manager *EnvManager

type EnvManager struct {
	init     bool // 是否初始化
	grpcConn *grpc.ClientConn
	client   nidavellir.ConfigServiceClient
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
	client, err := grpc.Dial(
		config.ApolloConf.GRPC.UdsAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		m.init = false
		logger.LoggerSugar.Errorf("grpc dial err %s", err.Error())
		return
	}
	m.client = nidavellir.NewConfigServiceClient(client)
	m.init = true
}

func (m *EnvManager) Close() {
	if m.grpcConn != nil {
		if err := m.grpcConn.Close(); err != nil {
			logger.LoggerSugar.Errorf("grpc conn close err %s", err.Error())
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
				logger.LoggerSugar.Errorf("decrypt [%s]-[%s] err %s", service, item.Key, err.Error())
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
		logger.LoggerSugar.Errorf("list services err %s", err.Error())
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
