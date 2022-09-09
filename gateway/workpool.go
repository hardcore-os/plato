package gateway

import (
	"fmt"

	"github.com/hardcore-os/plato/common/config"
	"github.com/panjf2000/ants"
)

var wPool *ants.Pool

func initWorkPoll() {
	var err error
	if wPool, err = ants.NewPool(config.GetGatewayWorkerPoolNum()); err != nil {
		fmt.Printf("InitWorkPoll.err :%s num:%d\n", err.Error(), config.GetGatewayWorkerPoolNum())
	}
}
