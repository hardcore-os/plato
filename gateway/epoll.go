package gateway

import (
	"fmt"
	"log"
	"net"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/hardcore-os/plato/common/config"
	"golang.org/x/sys/unix"
)

// 全局对象
var ep *ePool    // epoll池
var tcpNum int32 // 当前服务允许接入的最大tcp连接数

type ePool struct {
	eChan  chan *connection
	tables sync.Map
	eSize  int
	done   chan struct{}

	ln *net.TCPListener
	f  func(c *connection, ep *epoller)
}

func initEpoll(ln *net.TCPListener, f func(c *connection, ep *epoller)) {
	setLimit()
	ep = newEPool(ln, f)
	ep.createAcceptProcess()
	ep.startEPool()
}

func newEPool(ln *net.TCPListener, cb func(c *connection, ep *epoller)) *ePool {
	return &ePool{
		eChan:  make(chan *connection, config.GetGatewayEpollerChanNum()),
		done:   make(chan struct{}),
		eSize:  config.GetGatewayEpollerNum(),
		tables: sync.Map{},
		ln:     ln,
		f:      cb,
	}
}

// 创建一个专门处理 accept 事件的协程，与当前cpu的核数对应，能够发挥最大功效
func (e *ePool) createAcceptProcess() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				conn, e := e.ln.AcceptTCP()
				// 限流熔断
				if !checkTcp() {
					_ = conn.Close()
					continue
				}
				setTcpConifg(conn)
				if e != nil {
					if ne, ok := e.(net.Error); ok && ne.Temporary() {
						fmt.Errorf("accept temp err: %v", ne)
						continue
					}
					fmt.Errorf("accept err: %v", e)
				}
				c := NewConnection(conn)
				ep.addTask(c)
			}
		}()
	}
}

func (e *ePool) startEPool() {
	for i := 0; i < e.eSize; i++ {
		go e.startEProc()
	}
}

// 轮询器池 处理器
func (e *ePool) startEProc() {
	ep, err := newEpoller()
	if err != nil {
		panic(err)
	}
	// 监听连接创建事件
	go func() {
		for {
			select {
			case <-e.done:
				return
			case conn := <-e.eChan:
				addTcpNum()
				fmt.Printf("tcpNum:%d\n", tcpNum)
				if err := ep.add(conn); err != nil {
					fmt.Printf("failed to add connection %v\n", err)
					conn.Close() //登录未成功直接关闭连接
					continue
				}
				fmt.Printf("EpollerPool new connection[%v] tcpSize:%d\n", conn.RemoteAddr(), tcpNum)
			}
		}
	}()
	// 轮询器在这里轮询等待, 当有wait发生时则调用回调函数去处理
	for {
		select {
		case <-e.done:
			return
		default:
			connections, err := ep.wait(200) // 200ms 一次轮询避免 忙轮询
			if err != nil && err != syscall.EINTR {
				fmt.Printf("failed to epoll wait %v\n", err)
				continue
			}
			for _, conn := range connections {
				if conn == nil {
					break
				}
				e.f(conn, ep)
			}
		}
	}
}

func (e *ePool) addTask(c *connection) {
	e.eChan <- c
}

// epoller 对象 轮询器
type epoller struct {
	fd            int
	fdToConnTable sync.Map
}

func newEpoller() (*epoller, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &epoller{
		fd: fd,
	}, nil
}

// TODO: 默认水平触发模式,可采用非阻塞FD,优化边沿触发模式
func (e *epoller) add(conn *connection) error {
	// Extract file descriptor associated with the connection
	fd := conn.fd
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.EPOLLIN | unix.EPOLLHUP, Fd: int32(fd)})
	if err != nil {
		return err
	}
	e.fdToConnTable.Store(conn.fd, conn)
	ep.tables.Store(conn.id, conn)
	conn.BindEpoller(e)
	return nil
}
func (e *epoller) remove(c *connection) error {
	subTcpNum()
	fd := c.fd
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}
	ep.tables.Delete(c.id)
	e.fdToConnTable.Delete(c.fd)
	return nil
}
func (e *epoller) wait(msec int) ([]*connection, error) {
	events := make([]unix.EpollEvent, config.GetGatewayEpollWaitQueueSize())
	n, err := unix.EpollWait(e.fd, events, msec)
	if err != nil {
		return nil, err
	}
	var connections []*connection
	for i := 0; i < n; i++ {
		if conn, ok := e.fdToConnTable.Load(int(events[i].Fd)); ok {
			connections = append(connections, conn.(*connection))
		}
	}
	return connections, nil
}
func socketFD(conn *net.TCPConn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(*conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
	return int(pfdVal.FieldByName("Sysfd").Int())
}

// 设置go 进程打开文件数的限制
func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	log.Printf("set cur limit: %d", rLimit.Cur)
}

func addTcpNum() {
	atomic.AddInt32(&tcpNum, 1)
}

func getTcpNum() int32 {
	return atomic.LoadInt32(&tcpNum)
}
func subTcpNum() {
	atomic.AddInt32(&tcpNum, -1)
}

func checkTcp() bool {
	num := getTcpNum()
	maxTcpNum := config.GetGatewayMaxTcpNum()
	return num <= maxTcpNum
}

func setTcpConifg(c *net.TCPConn) {
	_ = c.SetKeepAlive(true)
}
