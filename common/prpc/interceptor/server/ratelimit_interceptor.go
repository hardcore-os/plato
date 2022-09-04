package server

import (
	"context"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	pcode "github.com/hardcore-os/plato/common/prpc/code"
	"github.com/juju/ratelimit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type MethodName string

type RateLimitConfig struct {
	Cap             int64         `json:"cap"`
	Rate            float64       `json:"rate"`
	WaitMaxDuration time.Duration `json:"wait_max_duration"`
}

// RateLimitUnaryServerInterceptor ...
func RateLimitUnaryServerInterceptor(configs map[MethodName]RateLimitConfig) grpc.UnaryServerInterceptor {
	buckets := make(map[MethodName]*ratelimit.Bucket)
	for name, config := range configs {
		buckets[name] = ratelimit.NewBucketWithRate(config.Rate, config.Cap)
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if bucket, ok := buckets[MethodName(info.FullMethod)]; ok {
			if _, ok := bucket.TakeMaxDuration(1, configs[MethodName(info.FullMethod)].WaitMaxDuration); !ok {
				logger.CtxErrorf(ctx, "too many request")
				return nil, status.Errorf(pcode.CodeTooManyRequest, "too many request")
			}

			return handler(ctx, req)
		}

		return handler(ctx, req)
	}
}
