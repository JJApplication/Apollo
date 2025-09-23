package svr

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AppManagerServiceServerImpl gRPC服务实现，调用app_manager提供的服务管理功能
// 所有方法均通过manager.go暴露的接口实现
type AppManagerServiceServerImpl struct {
	UnimplementedAppManagerServiceServer
}

// GetAllAppName 获取所有应用名称
func (s *AppManagerServiceServerImpl) GetAllAppName(ctx context.Context, req *emptypb.Empty) (*ListReply, error) {
	appNames, err := app_manager.GetAllAppName()
	if err != nil {
		return &ListReply{Items: nil, Err: err.Error()}, nil
	}
	return &ListReply{Items: appNames, Err: ""}, nil
}

// Status 获取指定应用状态
func (s *AppManagerServiceServerImpl) Status(ctx context.Context, req *AppName) (*StatusReply, error) {
	status, err := app_manager.Status(req.Name)
	if err != nil {
		return &StatusReply{Status: status, Err: err.Error()}, nil
	}
	return &StatusReply{Status: status, Err: ""}, nil
}

// Start 启动指定应用
func (s *AppManagerServiceServerImpl) Start(ctx context.Context, req *AppName) (*BoolReply, error) {
	success, err := app_manager.Start(req.Name)
	message := "启动成功"
	if err != nil {
		message = err.Error()
	}
	return &BoolReply{Ok: success, Msg: message}, nil
}

// Stop 停止指定应用
func (s *AppManagerServiceServerImpl) Stop(ctx context.Context, req *AppName) (*BoolReply, error) {
	success, err := app_manager.Stop(req.Name)
	message := "停止成功"
	if err != nil {
		message = err.Error()
	}
	return &BoolReply{Ok: success, Msg: message}, nil
}

// ReStart 重启指定应用
func (s *AppManagerServiceServerImpl) ReStart(ctx context.Context, req *AppName) (*BoolReply, error) {
	success, err := app_manager.ReStart(req.Name)
	message := "重启成功"
	if err != nil {
		message = err.Error()
	}
	return &BoolReply{Ok: success, Msg: message}, nil
}

// StartAll 启动所有应用
func (s *AppManagerServiceServerImpl) StartAll(ctx context.Context, req *emptypb.Empty) (*ListReply, error) {
	results, err := app_manager.StartAll()
	if err != nil {
		return &ListReply{Items: results, Err: err.Error()}, nil
	}
	return &ListReply{Items: results, Err: ""}, nil
}

// StopAll 停止所有应用
func (s *AppManagerServiceServerImpl) StopAll(ctx context.Context, req *emptypb.Empty) (*ListReply, error) {
	results, err := app_manager.StopAll()
	if err != nil {
		return &ListReply{Items: results, Err: err.Error()}, nil
	}
	return &ListReply{Items: results, Err: ""}, nil
}

// StatusAll 获取所有应用状态
func (s *AppManagerServiceServerImpl) StatusAll(ctx context.Context, req *emptypb.Empty) (*ListReply, error) {
	results, err := app_manager.StatusAll()
	if err != nil {
		return &ListReply{Items: results, Err: err.Error()}, nil
	}
	return &ListReply{Items: results, Err: ""}, nil
}

// SyncApp 同步指定应用
func (s *AppManagerServiceServerImpl) SyncApp(ctx context.Context, req *AppName) (*BoolReply, error) {
	success, err := app_manager.SyncApp(req.Name)
	message := "同步应用成功"
	if err != nil {
		message = err.Error()
	}
	return &BoolReply{Ok: success, Msg: message}, nil
}

// SyncAll 同步所有应用
func (s *AppManagerServiceServerImpl) SyncAll(ctx context.Context, req *emptypb.Empty) (*BoolReply, error) {
	err := app_manager.SyncAll()
	success := err == nil
	message := "同步所有应用成功"
	if err != nil {
		message = err.Error()
	}
	return &BoolReply{Ok: success, Msg: message}, nil
}

// NewGRPCServer 创建并启动gRPC服务器
// 监听配置文件中指定的UDS地址，注册AppManagerService服务
func NewGRPCServer() error {
	// 获取配置中的gRPC UDS地址
	grpcAddr := config.ApolloConf.Server.GRPC
	if grpcAddr == "" {
		logger.LoggerSugar.Error("gRPC uds address is not configured")
		return fmt.Errorf("gRPC uds address is not configured")
	}

	// 删除已存在的socket文件
	if err := os.RemoveAll(grpcAddr); err != nil {
		logger.LoggerSugar.Errorf("delete unix sock errror: %v", err)
		return err
	}

	// 监听UDS地址
	listener, err := net.Listen("unix", grpcAddr)
	if err != nil {
		logger.LoggerSugar.Errorf("listen gRPC address error: %v", err)
		return err
	}

	// 创建gRPC服务器
	grpcServer := grpc.NewServer()

	// 注册AppManagerService服务
	RegisterAppManagerServiceServer(grpcServer, &AppManagerServiceServerImpl{})

	logger.LoggerSugar.Infof("gRPC server started, listen at: %s", grpcAddr)

	// 启动服务器（阻塞）
	if err = grpcServer.Serve(listener); err != nil {
		logger.LoggerSugar.Errorf("gRPC start failed: %v", err)
		return err
	}

	return nil
}
