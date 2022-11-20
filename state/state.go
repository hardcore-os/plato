package state

import (
	"context"
	"sync"
	"time"

	"github.com/hardcore-os/plato/common/idl/message"
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
	msgTimer    *timingwheel.Timer
	connID      uint64
	maxClientID uint64
	msgID       uint64 // test专用
}

// 这里可以考虑设计成无锁的并发模式
func (c *connState) checkUPMsg(clientID uint64) bool {
	c.Lock()
	defer c.Unlock()
	if c.maxClientID+1 == clientID {
		return true
	}
	return false
}

// 这里可以考虑设计成无锁的并发模式
func (c *connState) addMaxClientID() {
	c.Lock()         // 不要迷恋原子操作，如果锁的临界区很小，性能与原子操作相差无己并发性能瓶颈，保持简单可靠即可。
	defer c.Unlock() // go的读写锁本身就有自旋等无锁优化
	c.maxClientID++
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

func rePush(connID uint64, msgData []byte) {
	sendMsg(connID, message.CmdType_Push, msgData)
	if data, ok := connToStateTable.Load(connID); ok {
		state, _ := data.(*connState)
		state.Lock()
		defer state.Unlock()
		if state.msgTimer != nil {
			state.msgTimer.Stop()
		}
		state.msgTimer = AfterFunc(100*time.Millisecond, func() {
			rePush(connID, msgData)
		})
	}
}
