package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/hardcore-os/plato/common/config"
	"github.com/hardcore-os/plato/common/idl/domain/user"
	"github.com/hardcore-os/plato/common/prpc"
	"github.com/hardcore-os/plato/domain/user/rpc/service"
	"google.golang.org/grpc"
)

const (
	testIp      = "127.0.0.1"
	testPort    = 8869
	serviceName = "plato.domain.user"
)

var userCli service.UserClient

// 在测试之前启动模拟的 gRPC 服务器
// go test -timeout 30s -run ^TestUserDomain$ -v github.com/hardcore-os/plato/domain/user
func TestUserDomain(t *testing.T) {
	mockRPC()
	ctx := context.Background()
	resp, err := userCli.CreateUsers(ctx, mockCreateUsersRequest())
	if err != nil {
		panic(err)
	}
	fmt.Printf("CreateUsers reps=%+v\n", *resp)

	queryResp, err := userCli.QueryUsers(ctx, mockQueryUsersRequest())
	if err != nil {
		panic(err)
	}
	fmt.Printf("QueryUsers1 reps=%+v\n", *queryResp)

	updateResp, err := userCli.UpdateUsers(ctx, mockUpdateUsersRequest())
	if err != nil {
		panic(err)
	}
	fmt.Printf("UpdateUsers reps=%+v\n", *updateResp)

	queryResp, err = userCli.QueryUsers(ctx, mockQueryUsersRequest())
	if err != nil {
		panic(err)
	}
	fmt.Printf("QueryUsers2 reps=%+v\n", *queryResp)
}

func mockRPC() {
	// 只初始化一次
	if userCli != nil {
		return
	}
	// 初始化配置
	config.Init("../../plato.yaml")
	// 启动server
	go func() {
		mokcServer()
	}()
	cli, err := prpc.NewPClient(serviceName)
	if err != nil {
		panic(err)
	}
	conn, err := cli.DialByEndPoint(fmt.Sprintf("%s:%d", testIp, testPort))
	if err != nil {
		panic(err)
	}
	userCli = service.NewUserClient(conn)
}

func mokcServer() {
	service.Init(true)
	s := prpc.NewPServer(prpc.WithServiceName(serviceName), prpc.WithIP(testIp), prpc.WithPort(testPort), prpc.WithWeight(100))
	s.RegisterService(func(server *grpc.Server) {
		service.RegisterUserServer(server, &service.Service{})
	})
	s.Start(context.TODO())
}

func mockQueryUsersRequest() *service.QueryUsersRequest {
	return &service.QueryUsersRequest{
		Opts: map[uint64]*service.QueryUserOption{
			1: {
				AllDevice:    false,
				ActiveDevice: true,
			},
			2: {
				AllDevice:    true,
				ActiveDevice: false,
			},
		},
	}
}

func mockCreateUsersRequest() *service.CreateUsersRequest {
	return &service.CreateUsersRequest{
		Users: []*user.UserDTO{
			{
				UserID:      1,
				Information: &user.InformationDTO{Nickname: "test1"},
			},
			{
				UserID:      2,
				Information: &user.InformationDTO{Nickname: "test2"},
			},
		},
	}
}

func mockUpdateUsersRequest() *service.UpdateUsersRequest {
	return &service.UpdateUsersRequest{
		Users: []*user.UserDTO{
			{
				UserID:      1,
				Information: &user.InformationDTO{Nickname: "test3"},
			},
			{
				UserID:      2,
				Information: &user.InformationDTO{Nickname: "test4"},
			},
		},
	}
}
