package main

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/hardcore-os/plato/common/config"
	"github.com/hardcore-os/plato/common/prpc"
	"github.com/hardcore-os/plato/common/prpc/example/helloservice"
	ptrace "github.com/hardcore-os/plato/common/prpc/trace"
)

func main() {
	config.Init(currentFileDir() + "/prpc_client.yaml")

	ptrace.StartAgent()
	defer ptrace.StopAgent()

	pCli, _ := prpc.NewPClient("prpc_server")

	ctx, _ := context.WithTimeout(context.TODO(), 100*time.Second)
	cli := helloservice.NewGreeterClient(pCli.Conn())
	resp, err := cli.SayHello(ctx, &helloservice.HelloRequest{
		Name: "xxxxxx",
	})
	fmt.Println(resp, err)
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
