package main

import (
	"context"
	"runtime"
	"strings"

	"github.com/hardcore-os/plato/common/config"

	"github.com/hardcore-os/plato/common/prpc"
	"github.com/hardcore-os/plato/common/prpc/example/helloservice"
	ptrace "github.com/hardcore-os/plato/common/prpc/trace"
	"google.golang.org/grpc"
)

const (
	testIp   = "127.0.0.1"
	testPort = 8867
)

func main() {
	config.Init(currentFileDir() + "/prpc_server.yaml")

	ptrace.StartAgent()
	defer ptrace.StopAgent()

	s := prpc.NewPServer(prpc.WithServiceName("prpc_server"), prpc.WithIP(testIp), prpc.WithPort(testPort), prpc.WithWeight(100))
	s.RegisterService(func(server *grpc.Server) {
		helloservice.RegisterGreeterServer(server, helloservice.HelloServer{})
	})
	s.Start(context.TODO())
}

func currentFileDir() string {
	_, file, _, ok := runtime.Caller(1)
	parts := strings.Split(file, "/")

	if !ok {
		return ""
	}

	dir := ""
	for i := 0; i < len(parts)-1; i++ {
		dir += "/" + parts[i]
	}

	return dir[1:]
}
