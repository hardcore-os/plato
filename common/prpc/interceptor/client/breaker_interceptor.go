package client

import (
	"context"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BreakerUnaryClientInterceptor 熔断器，配置其实都可以考虑用option选项模式实现，等待有人缘人优化吧
func BreakerUnaryClientInterceptor(name string, maxRequest uint32, interval, timeout time.Duration, readyToTrip func(counts gobreaker.Counts) bool) grpc.UnaryClientInterceptor {
	var cb *gobreaker.CircuitBreaker
	cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        name,
		MaxRequests: maxRequest,
		Interval:    interval,
		ReadyToTrip: readyToTrip,
		IsSuccessful: func(err error) bool {
			switch status.Code(err) {
			case codes.DeadlineExceeded, codes.Internal, codes.Unavailable, codes.DataLoss, codes.Unimplemented:
				return false
			default:
				return true
			}
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			logger.Errorf("name:%s,old state:%s,new state:%s", name, from, to)
		},
	})

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		_, err := cb.Execute(func() (interface{}, error) {
			err := invoker(ctx, method, req, reply, cc, opts...)
			return nil, err
		})

		return err
	}
}
