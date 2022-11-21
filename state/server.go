package state

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hardcore-os/plato/common/config"
	"github.com/hardcore-os/plato/common/prpc"
	"google.golang.org/protobuf/proto"

	"github.com/hardcore-os/plato/common/idl/message"
	"github.com/hardcore-os/plato/state/rpc/client"
	"github.com/hardcore-os/plato/state/rpc/service"
	"google.golang.org/grpc"
)

// RunMain 启动网关服务
func RunMain(path string) {
	config.Init(path)
	cmdChannel = make(chan *service.CmdContext, config.GetSateCmdChannelNum())
	connToStateTable = sync.Map{}
	s := prpc.NewPServer(
		prpc.WithServiceName(config.GetStateServiceName()),
		prpc.WithIP(config.GetSateServiceAddr()),
		prpc.WithPort(config.GetSateServerPort()), prpc.WithWeight(config.GetSateRPCWeight()))

	s.RegisterService(func(server *grpc.Server) {
		service.RegisterStateServer(server, &service.Service{CmdChannel: cmdChannel})
	})
	// 初始化RPC 客户端
	client.Init()
	// 启动时间轮
	InitTimer()
	// 启动 命令处理写协程
	go cmdHandler()
	// 启动 rpc server
	s.Start(context.TODO())
}

func cmdHandler() {
	for cmdCtx := range cmdChannel {
		switch cmdCtx.Cmd {
		case service.CancelConnCmd:
			fmt.Printf("cancelconn endpoint:%s, fd:%d, data:%+v", cmdCtx.Endpoint, cmdCtx.ConnID, cmdCtx.Payload)
		case service.SendMsgCmd:
			msgCmd := &message.MsgCmd{}
			err := proto.Unmarshal(cmdCtx.Payload, msgCmd)
			if err != nil {
				fmt.Printf("SendMsgCmd:err=%s\n", err.Error())
			}
			msgCmdHandler(cmdCtx, msgCmd)
		}
	}
}

func msgCmdHandler(cmdCtx *service.CmdContext, msgCmd *message.MsgCmd) {
	switch msgCmd.Type {
	case message.CmdType_Login:
		loginMsgHandler(cmdCtx, msgCmd)
	case message.CmdType_Heartbeat:
		hearbeatMsgHandler(cmdCtx, msgCmd)
	case message.CmdType_ReConn:
		reConnMsgHandler(cmdCtx, msgCmd)
	case message.CmdType_UP:
		upMsgHandler(cmdCtx, msgCmd)
	case message.CmdType_ACK:
		ackMsgHandler(cmdCtx, msgCmd)
	}
}

// 处理下行消息
func ackMsgHandler(cmdCtx *service.CmdContext, msgCmd *message.MsgCmd) {
	ackMsg := &message.ACKMsg{}
	err := proto.Unmarshal(msgCmd.Payload, ackMsg)
	if err != nil {
		fmt.Printf("ackMsgHandler:err=%s\n", err.Error())
		return
	}
	if data, ok := connToStateTable.Load(ackMsg.ConnID); ok {
		state, _ := data.(*connState)
		state.Lock()
		defer state.Unlock()
		if state.msgTimer != nil {
			state.msgTimer.Stop()
			state.msgTimer = nil
		}
	}
}

// 处理上行消息，并进行消息可靠性检查
func upMsgHandler(cmdCtx *service.CmdContext, msgCmd *message.MsgCmd) {
	upMsg := &message.UPMsg{}
	err := proto.Unmarshal(msgCmd.Payload, upMsg)
	if err != nil {
		fmt.Printf("upMsgHandler:err=%s\n", err.Error())
		return
	}
	if data, ok := connToStateTable.Load(upMsg.Head.ConnID); ok {
		state, _ := data.(*connState)
		if state.checkUPMsg(upMsg.Head.ClientID) {
			// 调用下游业务层rpc，只有当rpc回复成功后才能更新max_clientID
			// 这里先假设成功
			state.addMaxClientID()
			// TODO 这里构建下行消息并发送过去，msg_id先在state中自增
			state.msgID++
			sendACKMsg(message.CmdType_UP, cmdCtx.ConnID, upMsg.Head.ClientID, 0, "ok")

			// TODO 先在这里push消息
			pushMsg := &message.PushMsg{
				MsgID:   state.msgID,
				Content: upMsg.UPMsgBody, // 直接ping-pong
			}
			if data, err := proto.Marshal(pushMsg); err == nil {
				sendMsg(state.connID, message.CmdType_Push, data)
				if state.msgTimer != nil {
					state.msgTimer.Stop()
				}
				// 创建定时器
				t := AfterFunc(100*time.Millisecond, func() {
					rePush(cmdCtx.ConnID, data)
				})
				state.msgTimer = t
			} else {
				fmt.Printf("Marshal:err=%s\n", err.Error())
			}
		}
		// TODO 如果没有通过检查，当作先直接忽略即可
	}
}

func reConnMsgHandler(cmdCtx *service.CmdContext, msgCmd *message.MsgCmd) {
	reConnMsg := &message.ReConnMsg{}
	err := proto.Unmarshal(msgCmd.Payload, reConnMsg)
	if err != nil {
		fmt.Printf("reConnMsgHandler:err=%s\n", err.Error())
		return
	}
	// 重连的消息头中的connID才是上一次断开连接的connID
	if data, ok := connToStateTable.Load(reConnMsg.Head.ConnID); ok {
		state, _ := data.(*connState)
		state.Lock()
		defer state.Unlock()
		// 先停止定时任务的回调
		if state.reConnTimer != nil {
			state.reConnTimer.Stop()
			state.reConnTimer = nil // 重连定时器被清除
		}
		// 从索引中删除 旧的connID
		connToStateTable.Delete(reConnMsg.Head.ConnID)
		// 变更connID, cmdCtx中的connID才是 gateway重连的新连接
		state.connID = cmdCtx.ConnID
		connToStateTable.Store(cmdCtx.ConnID, state)
		sendACKMsg(message.CmdType_ReConn, cmdCtx.ConnID, 0, 0, "reconn ok")
	} else {
		sendACKMsg(message.CmdType_ReConn, cmdCtx.ConnID, 0, 1, "reconn failed")
	}
}

func hearbeatMsgHandler(cmdCtx *service.CmdContext, msgCmd *message.MsgCmd) {
	heartMsg := &message.HeartbeatMsg{}
	err := proto.Unmarshal(msgCmd.Payload, heartMsg)
	if err != nil {
		fmt.Printf("hearbeatMsgHandler:err=%s\n", err.Error())
		return
	}
	if data, ok := connToStateTable.Load(cmdCtx.ConnID); ok {
		sate, _ := data.(*connState)
		sate.reSetHeartTimer()
	}
	// 未减少通信量，可以暂时不回复心跳的ack
}

func loginMsgHandler(cmdCtx *service.CmdContext, msgCmd *message.MsgCmd) {
	loginMsg := &message.LoginMsg{}
	err := proto.Unmarshal(msgCmd.Payload, loginMsg)
	if err != nil {
		fmt.Printf("loginMsgHandler:err=%s\n", err.Error())
		return
	}
	if loginMsg.Head != nil {
		// 这里会把 login msg 传送给业务层做处理
		fmt.Println("loginMsgHandler", loginMsg.Head.DeviceID)
	}
	// 创建定时器
	t := AfterFunc(300*time.Second, func() {
		clearState(cmdCtx.ConnID)
	})
	// 初始化连接的状态
	connToStateTable.Store(cmdCtx.ConnID, &connState{heartTimer: t, connID: cmdCtx.ConnID})
	sendACKMsg(message.CmdType_Login, cmdCtx.ConnID, 0, 0, "login ok")
}
func sendACKMsg(ackType message.CmdType, connID, clientID uint64, code uint32, msg string) {
	ackMsg := &message.ACKMsg{}
	ackMsg.Code = code
	ackMsg.Msg = msg
	ackMsg.ConnID = connID
	ackMsg.Type = ackType
	ackMsg.ClientID = clientID
	downLoad, err := proto.Marshal(ackMsg)
	if err != nil {
		fmt.Println("sendACKMsg", err)
	}
	sendMsg(connID, message.CmdType_ACK, downLoad)
}

func sendMsg(connID uint64, ty message.CmdType, downLoad []byte) {
	mc := &message.MsgCmd{}
	mc.Type = ty
	mc.Payload = downLoad
	data, err := proto.Marshal(mc)
	ctx := context.TODO()
	if err != nil {
		fmt.Println("sendMsg", ty, err)
	}
	client.Push(&ctx, connID, data)
}
