package sdk

import (
	"net"
	"sync"
	"time"

	"github.com/goccy/go-json"
	"github.com/hardcore-os/plato/common/idl/message"
	"github.com/hardcore-os/plato/common/tcp"
	"google.golang.org/protobuf/proto"
)

const (
	MsgTypeText      = "text"
	MsgTypeAck       = "ack"
	MsgTypeReConn    = "reConn"
	MsgTypeHeartbeat = "heartbeat"
	MsgLogin         = "loginMsg"
)

type Chat struct {
	Nick             string
	UserID           string
	SessionID        string
	conn             *connect
	closeChan        chan struct{}
	MsgClientIDTable map[string]uint64
	sync.RWMutex
}

type Message struct {
	Type       string
	Name       string
	FormUserID string
	ToUserID   string
	Content    string
	Session    string
}

func NewChat(ip net.IP, port int, nick, userID, sessionID string) *Chat {
	chat := &Chat{
		Nick:             nick,
		UserID:           userID,
		SessionID:        sessionID,
		conn:             newConnet(ip, port),
		closeChan:        make(chan struct{}),
		MsgClientIDTable: make(map[string]uint64),
	}
	go chat.loop()
	chat.login()
	go chat.heartbeat()
	return chat
}
func (chat *Chat) Send(msg *Message) {
	data, _ := json.Marshal(msg)
	//key := fmt.Sprintf("%d", chat.conn.connID)
	upMsg := &message.UPMsg{
		Head: &message.UPMsgHead{
			ClientID:  chat.getClientID(chat.SessionID),
			ConnID:    chat.conn.connID,
			SessionId: chat.SessionID,
		},
		UPMsgBody: data,
	}
	palyload, _ := proto.Marshal(upMsg)
	chat.conn.send(message.CmdType_UP, palyload)
}

func (chat *Chat) GetCurClientID() uint64 {
	if id, ok := chat.MsgClientIDTable[chat.SessionID]; ok {
		return id
	}
	return 0
}

//Close close chat
func (chat *Chat) Close() {
	chat.conn.close()
	close(chat.closeChan)
	close(chat.conn.recvChan)
	close(chat.conn.sendChan)
}

func (chat *Chat) ReConn() {
	chat.Lock()
	defer chat.Unlock()
	// 需要重置clientId
	chat.MsgClientIDTable = make(map[string]uint64)
	chat.conn.reConn()
	chat.reConn()
}

//Recv receive message
func (chat *Chat) Recv() <-chan *Message {
	return chat.conn.recv()
}

func (chat *Chat) loop() {
Loop:
	for {
		select {
		case <-chat.closeChan:
			return
		default:
			mc := &message.MsgCmd{}
			data, err := tcp.ReadData(chat.conn.conn)
			if err != nil {
				goto Loop
			}
			err = proto.Unmarshal(data, mc)
			if err != nil {
				panic(err)
			}
			var msg *Message
			switch mc.Type {
			case message.CmdType_ACK:
				msg = handAckMsg(chat.conn, mc.Payload)
			case message.CmdType_Push:
				msg = handPushMsg(chat.conn, mc.Payload)

			}
			chat.conn.recvChan <- msg
		}
	}
}

func (chat *Chat) getClientID(sessionID string) uint64 {
	chat.Lock()
	defer chat.Unlock()
	var res uint64
	if id, ok := chat.MsgClientIDTable[sessionID]; ok {
		res = id
	}
	chat.MsgClientIDTable[sessionID] = res + 1
	return res
}

func (chat *Chat) login() {
	loginMsg := message.LoginMsg{
		Head: &message.LoginMsgHead{
			DeviceID: 123,
		},
	}
	palyload, err := proto.Marshal(&loginMsg)
	if err != nil {
		panic(err)
	}
	chat.conn.send(message.CmdType_Login, palyload)
}

func (chat *Chat) reConn() {
	reConn := message.ReConnMsg{
		Head: &message.ReConnMsgHead{
			ConnID: chat.conn.connID,
		},
	}
	palyload, err := proto.Marshal(&reConn)
	if err != nil {
		panic(err)
	}
	chat.conn.send(message.CmdType_ReConn, palyload)
}

func (chat *Chat) heartbeat() {
	tc := time.NewTicker(1 * time.Second)
	defer func() {
		chat.heartbeat()
	}()
loop:
	for {
		select {
		case <-chat.closeChan:
			return
		case <-tc.C:
			hearbeat := message.HeartbeatMsg{
				Head: &message.HeartbeatMsgHead{},
			}
			palyload, err := proto.Marshal(&hearbeat)
			if err != nil {
				panic(err)
			}
			err = chat.conn.send(message.CmdType_Heartbeat, palyload)
			if err != nil {
				goto loop
			}
		}
	}
}
