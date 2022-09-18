package service

import (
	context "context"
	"fmt"
)

const (
	CancelConnCmd = 1
	SendMsgCmd    = 2
)

type CmdContext struct {
	Ctx      *context.Context
	Cmd      int32
	Endpoint string
	FD       int
	Playload []byte
}

type Service struct {
	CmdChannel chan *CmdContext
}

func (s *Service) CancelConn(ctx context.Context, sr *StateRequest) (*StateResponse, error) {
	c := context.TODO()
	s.CmdChannel <- &CmdContext{
		Ctx:      &c,
		Cmd:      CancelConnCmd,
		FD:       int(sr.GetFd()),
		Endpoint: sr.GetEndpoint(),
	}
	return &StateResponse{
		Code: 0,
		Msg:  "success",
	}, nil
}

func (s *Service) SendMsg(ctx context.Context, sr *StateRequest) (*StateResponse, error) {
	fmt.Println("state.SendMsg.ok")
	c := context.TODO()
	s.CmdChannel <- &CmdContext{
		Ctx:      &c,
		Cmd:      SendMsgCmd,
		FD:       int(sr.GetFd()),
		Endpoint: sr.GetEndpoint(),
		Playload: sr.GetData(),
	}
	fmt.Println("state.SendMsg.okk")
	return &StateResponse{
		Code: 0,
		Msg:  "success",
	}, nil
}
