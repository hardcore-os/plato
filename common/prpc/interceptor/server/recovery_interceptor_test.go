package server

import (
	"context"
	"testing"

	"google.golang.org/grpc"
)

func TestRecoveryUnaryServerInterceptor(t *testing.T) {
	RecoveryUnaryServerInterceptor()(context.Background(), nil, &grpc.UnaryServerInfo{
		FullMethod: "/helloworld.Greeter/SayHello",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		panic("xxxxx")
		return nil, nil
	})
}
