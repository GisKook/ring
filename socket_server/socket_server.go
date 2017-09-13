package socket_server

import (
	"github.com/gansidui/gotcp"
	"github.com/giskook/bed2/base"
	"github.com/giskook/bed2/conf"
	"log"
	"net"
	"time"
)

type SocketServer struct {
	conf      *conf.Conf
	srv       *gotcp.Server
	cm        *ConnMgr
	SocketIn  chan *base.SocketData
	SocketOut chan *base.SocketData
}

func NewSocketServer(conf *conf.Conf) *SocketServer {
	return &SocketServer{
		conf:      conf,
		cm:        NewConnMgr(),
		SocketIn:  make(chan *base.SocketData),
		SocketOut: make(chan *base.SocketData),
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
	log.Println("socket Listening:", listener.Addr())

	return nil
}

func (ss *SocketServer) Send(id uint64, p gotcp.Packet) error {
	c := ss.cm.Get(id)
	if c != nil {
		return c.Send(p)
	}

	return base.ErrBedOffline

}

func (ss *SocketServer) Close() {
	close(ss.SocketOutResult)
	ss.srv.Stop()
}
