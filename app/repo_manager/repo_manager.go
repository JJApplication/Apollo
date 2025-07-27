package repo_manager

import (
	"context"
	"sync"

	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/grpc/twt"
	"github.com/JJApplication/Apollo/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 环境变量管理实例

const RepoManagerPrefix = "[Repo Manager]"

var manager *RepoManager

type RepoManager struct {
	init     bool // 是否初始化
	grpcConn *grpc.ClientConn
	client   twt.RepositoryServiceClient
}

func InitRepoManager() {
	go func() {
		manager = new(RepoManager)
		manager.Init()
	}()
}

func GetRepoManager() *RepoManager {
	if manager == nil || !manager.init {
		sync.OnceFunc(func() {
			manager = new(RepoManager)
			manager.Init()
		})
		return manager
	}
	return manager
}

func (m *RepoManager) Init() {
	addr := config.ApolloConf.GRPC.GetAddr("TwT")
	logger.LoggerSugar.Infof("%s start to init grpc client at %s", RepoManagerPrefix, addr)
	client, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithMaxCallAttempts(config.ApolloConf.GRPC.MaxAttempt))
	if err != nil {
		m.init = false
		logger.LoggerSugar.Errorf("%s grpc dial err %s", RepoManagerPrefix, err.Error())
		return
	}
	m.client = twt.NewRepositoryServiceClient(client)
	m.init = true
	logger.LoggerSugar.Infof("%s grpc client init success at: %s", RepoManagerPrefix, addr)
}

func (m *RepoManager) Close() {
	if m.grpcConn != nil {
		if err := m.grpcConn.Close(); err != nil {
			logger.LoggerSugar.Errorf("%s grpc conn close err %s", RepoManagerPrefix, err.Error())
		}
	}
}

func (m *RepoManager) IsInit() bool {
	return m.init
}

// GetRepos 获取全部仓库信息
func (m *RepoManager) GetRepos() *twt.GetRepositoriesResponse {
	ctx := context.Background()
	items, err := m.client.GetRepositories(ctx, &twt.GetRepositoriesRequest{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		logger.LoggerSugar.Errorf("%s list all repos err %s", RepoManagerPrefix, err.Error())
		return nil
	}
	return items
}

// GetRepo 列举某个仓库详情
func (m *RepoManager) GetRepo(fullName string) *twt.GetRepositoryResponse {
	if !m.IsInit() {
		return nil
	}
	ctx := context.Background()
	data, err := m.client.GetRepository(ctx, &twt.GetRepositoryRequest{
		FullName: fullName,
	})
	if err != nil {
		logger.LoggerSugar.Errorf("%s list repo of %s err %s", RepoManagerPrefix, fullName, err.Error())
		return nil
	}

	return data
}

// SyncRepos 同步所有仓库信息
func (m *RepoManager) SyncRepos() *twt.SyncRepositoriesResponse {
	if !m.IsInit() {
		return nil
	}
	ctx := context.Background()
	data, err := m.client.SyncRepositories(ctx, &twt.SyncRepositoriesRequest{})
	if err != nil {
		logger.LoggerSugar.Errorf("%s sync repo err %s", RepoManagerPrefix, err.Error())
		return nil
	}
	return data
}

// GetCommits 获取仓库的提交记录
func (m *RepoManager) GetCommits(fullName string) *twt.GetCommitsResponse {
	if !m.IsInit() {
		return nil
	}
	ctx := context.Background()
	data, err := m.client.GetCommits(ctx, &twt.GetCommitsRequest{
		RepositoryFullName: fullName,
		Limit:              100,
		Offset:             0,
	})
	if err != nil {
		logger.LoggerSugar.Errorf("%s list repo %s commits err %s", RepoManagerPrefix, fullName, err.Error())
		return nil
	}
	return data
}

// SyncCommits 同步某个仓库的commits
func (m *RepoManager) SyncCommits(fullName string) *twt.SyncCommitsResponse {
	if !m.IsInit() {
		return nil
	}
	ctx := context.Background()
	data, err := m.client.SyncCommits(ctx, &twt.SyncCommitsRequest{
		RepositoryFullName: fullName,
		Limit:              100,
	})
	if err != nil {
		logger.LoggerSugar.Errorf("%s sync repo %s commits err %s", RepoManagerPrefix, fullName, err.Error())
		return nil
	}
	return data
}
