package server

import (
	"context"
	"runtime"

	"github.com/bytedance/gopkg/util/logger"
	"google.golang.org/grpc"
)

// RecoveryUnaryServerInterceptor recovery中间件最好放在第一个去执行
func RecoveryUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if err := recover(); err != nil {
				stack := make([]byte, 4096)
				stack = stack[:runtime.Stack(stack, false)]
				logger.CtxErrorf(ctx, "err:%v\nstack:%s", err, stack)
			}

		}()

		return handler(ctx, req)
	}
}
