package prpc

import (
	"context"
	"fmt"
	"time"

	"github.com/hardcore-os/plato/common/prpc/discov/plugin"

	"google.golang.org/grpc/resolver"

	"github.com/hardcore-os/plato/common/prpc/discov"
	clientinterceptor "github.com/hardcore-os/plato/common/prpc/interceptor/client"
	presolver "github.com/hardcore-os/plato/common/prpc/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

const (
	dialTimeout = 5 * time.Second
)

type PClient struct {
	serviceName  string
	d            discov.Discovery
	interceptors []grpc.UnaryClientInterceptor
	conn         *grpc.ClientConn
}

// NewPClient ...
func NewPClient(serviceName string, interceptors ...grpc.UnaryClientInterceptor) (*PClient, error) {
	p := &PClient{
		serviceName:  serviceName,
		interceptors: interceptors,
	}

	if p.d == nil {
		dis, err := plugin.GetDiscovInstance()
		if err != nil {
			panic(err)
		}

		p.d = dis
	}

	resolver.Register(presolver.NewDiscovBuilder(p.d))

	conn, err := p.dial()
	p.conn = conn

	return p, err
}

// Conn return *grpc.ClientConn
func (p *PClient) Conn() *grpc.ClientConn {
	return p.conn
}

func (p *PClient) dial() (*grpc.ClientConn, error) {
	svcCfg := fmt.Sprintf(`{"loadBalancingPolicy":"%s"}`, roundrobin.Name)
	balancerOpt := grpc.WithDefaultServiceConfig(svcCfg)

	interceptors := []grpc.UnaryClientInterceptor{
		clientinterceptor.TraceUnaryClientInterceptor(),
		clientinterceptor.MetricUnaryClientInterceptor(),
	}
	interceptors = append(interceptors, p.interceptors...)

	options := []grpc.DialOption{
		balancerOpt,
		grpc.WithChainUnaryInterceptor(interceptors...),
		grpc.WithInsecure(),
	}

	ctx, _ := context.WithTimeout(context.Background(), dialTimeout)

	return grpc.DialContext(ctx, fmt.Sprintf("discov:///%v", p.serviceName), options...)
}
