package client

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"google.golang.org/grpc"
)

func TestTimeoutUnaryClientInterceptor(t *testing.T) {
	cc := new(grpc.ClientConn)
	err := TimeoutUnaryClientInterceptor(3*time.Second, 3*time.Second)(context.TODO(), "/create", nil, nil, cc,
		func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
			opts ...grpc.CallOption) error {
			return nil
		},
	)

	assert.NoError(t, err)

	err = TimeoutUnaryClientInterceptor(3*time.Second, 3*time.Second)(context.TODO(), "/create", nil, nil, cc,
		func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
			opts ...grpc.CallOption) error {
			time.Sleep(4 * time.Second)
			return nil
		},
	)

	assert.NoError(t, err)
}
