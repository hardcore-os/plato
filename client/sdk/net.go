package sdk

type connect struct {
	serverAddr         string
	sendChan, recvChan chan *Message
}

func newConnet(serverAddr string) *connect {
	return &connect{
		serverAddr: serverAddr,
		sendChan:   make(chan *Message),
		recvChan:   make(chan *Message),
	}
}

func (c *connect) send(data *Message) {
	// 直接发送给接收方
	c.recvChan <- data
}

func (c *connect) recv() <- chan *Message {
	return c.recvChan
}

func (c *connect) close() {
	// 目前没啥值得回收的
}
