package server

import (
	"context"
	"testing"

	pcode "github.com/hardcore-os/plato/common/prpc/code"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func TestRateLimitUnaryServerInterceptor(t *testing.T) {
	tests := []struct {
		name   string
		config map[MethodName]RateLimitConfig
		count  int
		res    error
	}{
		{
			"suc",
			map[MethodName]RateLimitConfig{
				"/helloworld.Greeter/SayHello": {
					Cap:  3,
					Rate: 1,
				},
			},
			1,
			nil,
		},
		{
			"fail",
			map[MethodName]RateLimitConfig{
				"/helloworld.Greeter/SayHello": {
					Cap:  3,
					Rate: 1,
				},
			},
			4,
			status.Errorf(pcode.CodeTooManyRequest, "too many request"),
		},
	}

	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			var err error
			r := RateLimitUnaryServerInterceptor(item.config)
			for i := 0; i < item.count; i++ {
				_, err = r(context.Background(), nil, &grpc.UnaryServerInfo{
					FullMethod: "/helloworld.Greeter/SayHello",
				}, func(ctx context.Context, req interface{}) (interface{}, error) {
					return nil, nil
				})
			}

			assert.Equal(t, item.res, err)
		})
	}
}
