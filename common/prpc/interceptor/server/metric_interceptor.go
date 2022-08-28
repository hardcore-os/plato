package server

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/hardcore-os/plato/common/prpc/prome"

	"github.com/hardcore-os/plato/common/prpc/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const nameSpace = "prpc_server"

var (
	serverHandleCounter = prome.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: nameSpace,
			Subsystem: "req",
			Name:      "client_handle_total",
		},
		[]string{"method", "server", "code", "ip"},
	)

	serverHandleHistogram = prome.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: nameSpace,
			Subsystem: "req",
			Name:      "client_handle_seconds",
		},
		[]string{"method", "server", "ip"},
	)
)

// MetricUnaryServerInterceptor ...
func MetricUnaryServerInterceptor(serverName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		beg := time.Now()
		resp, err = handler(ctx, req)

		code := status.Code(err)
		serverHandleCounter.WithLabelValues(info.FullMethod, serverName, code.String(), util.ExternalIP()).Inc()
		serverHandleHistogram.WithLabelValues(info.FullMethod, serverName, util.ExternalIP()).Observe(time.Since(beg).Seconds())

		return
	}
}
