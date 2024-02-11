package user

import (
	"context"

	"github.com/hardcore-os/plato/common/config"
	"github.com/hardcore-os/plato/common/prpc"
	"github.com/hardcore-os/plato/domain/user/rpc/client"
	"github.com/hardcore-os/plato/domain/user/rpc/service"
	"google.golang.org/grpc"
)

// RunMain user 领域服务
func RunMain(path string) {
	// 启动时的全局上下文
	ctx := context.TODO()
	// 初始化全局配置
	config.Init(path)
	// 初始化RPC
	client.Init()
	// 初始化存储层
	service.Init(false)
	// 注册rpc server
	s := prpc.NewPServer(
		prpc.WithServiceName(config.GetDomainUserServerName()),
		prpc.WithIP(config.GetDomainUserServerAddr()),
		prpc.WithPort(config.GetDomainUserServerPoint()), prpc.WithWeight(config.GetDomainUserRPCWeight()))
	s.RegisterService(func(server *grpc.Server) {
		service.RegisterUserServer(server, &service.Service{})
	})
	// 启动 rpc server
	s.Start(ctx)
}
