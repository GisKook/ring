package socket_server

import (
	"github.com/gansidui/gotcp"
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/conf"
	"log"
	"net"
	"sync"
	"time"
)

type SocketServer struct {
	conf      *conf.Conf
	srv       *gotcp.Server
	cm        *ConnMgr
	SocketIn  chan *base.SocketData
	SocketOut chan []byte
	exit      chan struct{}
	wait_exit *sync.WaitGroup
}

func NewSocketServer(conf *conf.Conf) *SocketServer {
	return &SocketServer{
		conf:      conf,
		cm:        NewConnMgr(),
		SocketIn:  make(chan *base.SocketData),
		SocketOut: make(chan []byte),
		exit:      make(chan struct{}),
		wait_exit: new(sync.WaitGroup),
	}
}

func (ss *SocketServer) Start() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+ss.conf.TcpServer.BindPort)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	config := &gotcp.Config{
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}

	ss.srv = gotcp.NewServer(config, ss, ss)

	go ss.srv.Start(listener, time.Second)
	log.Println("<INFO> socket listening:", listener.Addr())

	for i := 0; i < ss.conf.TcpServer.WorkerNum; i++ {
		ss.consumer_worker(ss.SocketOut)
	}

	return nil
}

func (ss *SocketServer) Send(id uint64, p gotcp.Packet) error {
	c := ss.cm.Get(id)
	if c != nil {
		return c.Send(p)
	}

	return base.ErrTerminalOffline

}

func (ss *SocketServer) Stop() {
	close(ss.exit)
	ss.wait_exit.Wait()
	close(ss.SocketOut)
	close(ss.SocketIn)

	ss.srv.Stop()
}
