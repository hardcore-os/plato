package state

import (
	"context"
	"sync"
	"time"

	"github.com/hardcore-os/plato/common/timingwheel"
	"github.com/hardcore-os/plato/state/rpc/client"
	"github.com/hardcore-os/plato/state/rpc/service"
)

var cmdChannel chan *service.CmdContext
var connToStateTable sync.Map

type connState struct {
	sync.RWMutex
	heartTimer  *timingwheel.Timer
	reConnTimer *timingwheel.Timer
	connID      uint64
}

func (c *connState) reSetHeartTimer() {
	c.Lock()
	defer c.Unlock()
	c.heartTimer.Stop()
	c.heartTimer = AfterFunc(5*time.Second, func() {
		clearState(c.connID)
	})
}

// 为了实现重连，这里不要立即释放连接的状态, 给予10s的延迟时间
func clearState(connID uint64) {
	if data, ok := connToStateTable.Load(connID); ok {
		state, _ := data.(*connState)
		state.Lock()
		defer state.Unlock()
		state.reConnTimer = AfterFunc(10*time.Second, func() {
			ctx := context.TODO()
			client.DelConn(&ctx, connID, nil)
			// 删除一些state的状态
			connToStateTable.Delete(connID)
		})
	}
}
