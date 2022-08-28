package prpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hardcore-os/plato/common/prpc/discov/plugin"

	"github.com/bytedance/gopkg/util/logger"

	"github.com/hardcore-os/plato/common/prpc/discov"
	serverinterceptor "github.com/hardcore-os/plato/common/prpc/interceptor/server"
	"google.golang.org/grpc"
)

type RegisterFn func(*grpc.Server)

type PServer struct {
	serverOptions
	registers    []RegisterFn
	interceptors []grpc.UnaryServerInterceptor
}

type serverOptions struct {
	serviceName string
	ip          string
	port        int
	weight      int
	health      bool
	d           discov.Discovery
}

type ServerOption func(opts *serverOptions)

// WithServiceName set serviceName
func WithServiceName(serviceName string) ServerOption {
	return func(opts *serverOptions) {
		opts.serviceName = serviceName
	}
}

// WithIP set ip
func WithIP(ip string) ServerOption {
	return func(opts *serverOptions) {
		opts.ip = ip
	}
}

// WithPort set port
func WithPort(port int) ServerOption {
	return func(opts *serverOptions) {
		opts.port = port
	}
}

// WithWeight set weight
func WithWeight(weight int) ServerOption {
	return func(opts *serverOptions) {
		opts.weight = weight
	}
}

// WithHealth set health
func WithHealth(health bool) ServerOption {
	return func(opts *serverOptions) {
		opts.health = health
	}
}

func NewPServer(opts ...ServerOption) *PServer {
	opt := serverOptions{}
	for _, o := range opts {
		o(&opt)
	}

	if opt.d == nil {
		dis, err := plugin.GetDiscovInstance()
		if err != nil {
			panic(err)
		}

		opt.d = dis
	}

	return &PServer{
		opt,
		make([]RegisterFn, 0),
		make([]grpc.UnaryServerInterceptor, 0),
	}
}

// RegisterService ...
// eg :
// p.RegisterService(func(server *grpc.Server) {
//     test.RegisterGreeterServer(server, &Server{})
// })
func (p *PServer) RegisterService(register ...RegisterFn) {
	p.registers = append(p.registers, register...)
}

// RegisterUnaryServerInterceptor 注册自定义拦截器，例如限流拦截器或者自己的一些业务自定义拦截器
func (p *PServer) RegisterUnaryServerInterceptor(i grpc.UnaryServerInterceptor) {
	p.interceptors = append(p.interceptors, i)
}

// Start 开启server
func (p *PServer) Start(ctx context.Context) {
	service := discov.Service{
		Name: p.serviceName,
		Endpoints: []*discov.Endpoint{
			{
				ServerName: p.serviceName,
				IP:         p.ip,
				Port:       p.port,
				Weight:     p.weight,
				Enable:     true,
			},
		},
	}

	// 加载中间件
	interceptors := []grpc.UnaryServerInterceptor{
		serverinterceptor.RecoveryUnaryServerInterceptor(),
		serverinterceptor.TraceUnaryServerInterceptor(),
		serverinterceptor.MetricUnaryServerInterceptor(p.serviceName),
	}
	interceptors = append(interceptors, p.interceptors...)

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors...))

	// 注册服务
	for _, register := range p.registers {
		register(s)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", p.ip, p.port))
	if err != nil {
		panic(err)
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	// 服务注册
	p.d.Register(ctx, &service)

	logger.Info("start PRCP success")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-c
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			s.Stop()
			p.d.UnRegister(ctx, &service)
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
