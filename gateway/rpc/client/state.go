package client

import (
	"context"
	"fmt"
	"time"

	"github.com/hardcore-os/plato/common/config"
	"github.com/hardcore-os/plato/common/prpc"
	"github.com/hardcore-os/plato/state/rpc/service"
)

var stateClient service.StateClient

func initStateClient() {
	pCli, err := prpc.NewPClient(config.GetStateServiceName())
	if err != nil {
		panic(err)
	}
	cli, err := pCli.DialByEndPoint(config.GetGatewayStateServerEndPoint())
	if err != nil {
		panic(err)
	}
	stateClient = service.NewStateClient(cli)
}

func CancelConn(ctx *context.Context, endpoint string, connID uint64, Payload []byte) error {
	rpcCtx, _ := context.WithTimeout(*ctx, 100*time.Millisecond)
	stateClient.CancelConn(rpcCtx, &service.StateRequest{
		Endpoint: endpoint,
		ConnID:   connID,
		Data:     Payload,
	})
	return nil
}

func SendMsg(ctx *context.Context, endpoint string, connID uint64, Payload []byte) error {
	rpcCtx, _ := context.WithTimeout(*ctx, 100*time.Millisecond)
	fmt.Println("sendMsg", connID, string(Payload))
	_, err := stateClient.SendMsg(rpcCtx, &service.StateRequest{
		Endpoint: endpoint,
		ConnID:   connID,
		Data:     Payload,
	})
	if err != nil {
		panic(err)
	}
	return nil
}
