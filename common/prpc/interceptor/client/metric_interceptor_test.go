package client

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/hardcore-os/plato/common/prpc/prome"
)

func TestMetricUnaryClientInterceptor(t *testing.T) {
	prome.StartAgent("0.0.0.0", 8927)

	cc := new(grpc.ClientConn)
	for {
		MetricUnaryClientInterceptor()(context.TODO(), "/create", nil, nil, cc,
			func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
				opts ...grpc.CallOption) error {
				return nil
			})
		time.Sleep(1 * time.Second)
	}

}
