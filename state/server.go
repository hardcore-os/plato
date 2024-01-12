package state

import (
	"context"
	"fmt"

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
	// 启动时的全局上下文
	ctx := context.TODO()
	// 初始化全局配置
	config.Init(path)
	// 初始化RPC 客户端
	client.Init()
	// 启动时间轮
	InitTimer()
	// 启动远程cache状态机组件
	InitCacheState(ctx)
	// 启动 命令处理写协程
	go cmdHandler()
	// 注册rpc server
	s := prpc.NewPServer(
		prpc.WithServiceName(config.GetStateServiceName()),
		prpc.WithIP(config.GetSateServiceAddr()),
		prpc.WithPort(config.GetSateServerPort()), prpc.WithWeight(config.GetSateRPCWeight()))
	s.RegisterService(func(server *grpc.Server) {
		service.RegisterStateServer(server, cs.server)
	})
	// 启动 rpc server
	s.Start(ctx)
}

// 消费信令通道，识别gateway与state server之间的协议路由
func cmdHandler() {
	for cmdCtx := range cs.server.CmdChannel {
		switch cmdCtx.Cmd {
		case service.CancelConnCmd:
			fmt.Printf("cancel conn endpoint:%s, coonID:%d, data:%+v\n", cmdCtx.Endpoint, cmdCtx.ConnID, cmdCtx.Payload)
			cs.connLogOut(*cmdCtx.Ctx, cmdCtx.ConnID)
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

// 识别消息类型，识别客户端与state server之间的协议路由
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

// 实现登陆功能
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
	err = cs.connLogin(*cmdCtx.Ctx, loginMsg.Head.DeviceID, cmdCtx.ConnID)
	if err != nil {
		panic(err)
	}
	sendACKMsg(message.CmdType_Login, cmdCtx.ConnID, 0, 0, "login ok")
}

// 心跳消息的处理
func hearbeatMsgHandler(cmdCtx *service.CmdContext, msgCmd *message.MsgCmd) {
	heartMsg := &message.HeartbeatMsg{}
	err := proto.Unmarshal(msgCmd.Payload, heartMsg)
	if err != nil {
		fmt.Printf("hearbeatMsgHandler:err=%s\n", err.Error())
		return
	}
	cs.reSetHeartTimer(cmdCtx.ConnID)
	fmt.Printf("hearbeatMsgHandler connID=%d\n", cmdCtx.ConnID)
	// TODO未减少通信量，可以暂时不回复心跳的ack
}

// 重连逻辑处理
func reConnMsgHandler(cmdCtx *service.CmdContext, msgCmd *message.MsgCmd) {
	reConnMsg := &message.ReConnMsg{}
	err := proto.Unmarshal(msgCmd.Payload, reConnMsg)
	var code uint32
	msg := "reconn ok"
	if err != nil {
		fmt.Printf("reConnMsgHandler:err=%s\n", err.Error())
		return
	}
	// 重连的消息头中的connID才是上一次断开连接的connID
	if err := cs.reConn(*cmdCtx.Ctx, reConnMsg.Head.ConnID, cmdCtx.ConnID); err != nil {
		code, msg = 1, "reconn failed"
		panic(err)
	}
	sendACKMsg(message.CmdType_ReConn, cmdCtx.ConnID, 0, code, msg)
}

// 处理上行消息，并进行消息可靠性检查
func upMsgHandler(cmdCtx *service.CmdContext, msgCmd *message.MsgCmd) {
	upMsg := &message.UPMsg{}
	err := proto.Unmarshal(msgCmd.Payload, upMsg)
	if err != nil {
		fmt.Printf("upMsgHandler:err=%s\n", err.Error())
		return
	}
	if cs.compareAndIncrClientID(*cmdCtx.Ctx, cmdCtx.ConnID, upMsg.Head.ClientID, upMsg.Head.SessionId) {
		// 调用下游业务层rpc，只有当rpc回复成功后才能更新max_clientID
		sendACKMsg(message.CmdType_UP, cmdCtx.ConnID, upMsg.Head.ClientID, 0, "ok")
		// TODO 这里应该调用业务层的代码
		pushMsg(*cmdCtx.Ctx, cmdCtx.ConnID, cs.msgID, 0, upMsg.UPMsgBody)
	}
}

// 处理下行消息ack回复
func ackMsgHandler(cmdCtx *service.CmdContext, msgCmd *message.MsgCmd) {
	ackMsg := &message.ACKMsg{}
	err := proto.Unmarshal(msgCmd.Payload, ackMsg)
	if err != nil {
		fmt.Printf("ackMsgHandler:err=%s\n", err.Error())
		return
	}
	cs.ackLastMsg(*cmdCtx.Ctx, ackMsg.ConnID, ackMsg.SessionID, ackMsg.MsgID)
}

// 由业务层调用，下行消息处理
func pushMsg(ctx context.Context, connID, sessionID, msgID uint64, data []byte) {
	// TODO 先在这里push消息
	pushMsg := &message.PushMsg{
		Content: data,
		MsgID:   cs.msgID,
	}
	if data, err := proto.Marshal(pushMsg); err != nil {
		fmt.Printf("Marshal:err=%s\n", err.Error())
	} else {
		//TODO 这里就要涉及到 下行消息的下发了,不管成功与否，都要更新last msg
		sendMsg(connID, message.CmdType_Push, data)
		err = cs.appendLastMsg(ctx, connID, pushMsg)
		if err != nil {
			panic(err)
		}
	}
}

// 发送ack msg
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

// 发送msg
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

// 重新发送push msg
func rePush(connID uint64) {
	pushMsg, err := cs.getLastMsg(context.Background(), connID)
	if err != nil {
		panic(err)
	}
	if pushMsg == nil {
		return
	}
	msgData, err := proto.Marshal(pushMsg)
	if err != nil {
		panic(err)
	}
	sendMsg(connID, message.CmdType_Push, msgData)
	if state, ok := cs.loadConnIDState(connID); ok {
		state.reSetMsgTimer(connID, pushMsg.SessionID, pushMsg.MsgID)
	}
}
