package server

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ptrace "github.com/hardcore-os/plato/common/prpc/trace"
	"google.golang.org/grpc"
)

func TestTraceUnaryServerInterceptor(t *testing.T) {
	ptrace.StartAgent()
	defer ptrace.StopAgent()

	//cc := new(grpc.ClientConn)
	TraceUnaryServerInterceptor()(context.Background(), nil, &grpc.UnaryServerInfo{
		FullMethod: "/helloworld.Greeter/SayHello",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	})

	TraceUnaryServerInterceptor()(context.Background(), nil, &grpc.UnaryServerInfo{
		FullMethod: "/helloworld.Greeter/SayBye",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, status.Error(codes.DataLoss, "dummy")
	})
}
