package service

import (
	context "context"
)

const (
	DelConnCmd = 1 // DelConn
	PushCmd    = 2 // push
)

type CmdContext struct {
	Ctx     *context.Context
	Cmd     int32
	ConnID  uint64
	Payload []byte
}

type Service struct {
	CmdChannel chan *CmdContext
}

func (s *Service) DelConn(ctx context.Context, gr *GatewayRequest) (*GatewayResponse, error) {
	c := context.TODO()
	s.CmdChannel <- &CmdContext{
		Ctx:    &c,
		Cmd:    DelConnCmd,
		ConnID: gr.ConnID,
	}
	return &GatewayResponse{
		Code: 0,
		Msg:  "success",
	}, nil
}

func (s *Service) Push(ctx context.Context, gr *GatewayRequest) (*GatewayResponse, error) {
	c := context.TODO()
	s.CmdChannel <- &CmdContext{
		Ctx:     &c,
		Cmd:     PushCmd,
		ConnID:  gr.ConnID,
		Payload: gr.GetData(),
	}
	return &GatewayResponse{
		Code: 0,
		Msg:  "success",
	}, nil
}
