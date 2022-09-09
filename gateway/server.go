package gateway

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/hardcore-os/plato/common/config"
	"github.com/hardcore-os/plato/common/tcp"
)

// RunMain 启动网关服务
func RunMain(path string) {
	config.Init(path)
	ln, err := net.ListenTCP("tcp", &net.TCPAddr{Port: config.GetGatewayServerPort()})
	if err != nil {
		log.Fatalf("StartTCPEPollServer err:%s", err.Error())
		panic(err)
	}
	initWorkPoll()
	initEpoll(ln, runProc)
	fmt.Println("-------------im gateway stated------------")
	select {}
}

func runProc(c *connection, ep *epoller) {
	// step1: 读取一个完整的消息包
	dataBuf, err := tcp.ReadData(c.conn)
	if err != nil {
		// 如果读取conn时发现连接关闭，则直接端口连接
		if errors.Is(err, io.EOF) {
			ep.remove(c)
		}
		return
	}
	err = wPool.Submit(func() {
		// step2:交给 state server rpc 处理
		bytes := tcp.DataPgk{
			Len:  uint32(len(dataBuf)),
			Data: dataBuf,
		}
		tcp.SendData(c.conn, bytes.Marshal())
	})
	if err != nil {
		fmt.Errorf("runProc:err:%+v\n", err.Error())
	}
}
