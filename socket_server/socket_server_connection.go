package socket_server

import (
	"bytes"
	"github.com/gansidui/gotcp"
	"github.com/giskook/ring/base"
	"time"
)

const (
	USER_STATUS_INIT    uint8 = 0
	USER_STATUS_NORMAL  uint8 = 1
	USER_STATUS_ILLEGAL uint8 = 2
)

type ConnConf struct {
	read_limit  int
	write_limit int
	uuid        uint32
}

type Connection struct {
	conf            *ConnConf
	c               *gotcp.Conn
	ID              uint64
	RecvBuffer      *bytes.Buffer
	read_timestamp  int64
	write_timestamp int64
	exit            chan struct{}
	status          uint8

	ticker *time.Ticker
}

func NewConnection(c *gotcp.Conn, conf *ConnConf) *Connection {
	tcp_c := c.GetRawConn()
	tcp_c.SetReadDeadline(time.Now().Add(time.Duration(conf.read_limit) * time.Second))
	tcp_c.SetWriteDeadline(time.Now().Add(time.Duration(conf.write_limit) * time.Second))
	return &Connection{
		conf:            conf,
		c:               c,
		read_timestamp:  time.Now().Unix(),
		write_timestamp: time.Now().Unix(),
		RecvBuffer:      bytes.NewBuffer([]byte{}),
		ticker:          time.NewTicker(10 * time.Second),
		exit:            make(chan struct{}),
	}
}

func (c *Connection) SetReadDeadline() {
	c.c.GetRawConn().SetReadDeadline(time.Now().Add(time.Duration(c.conf.read_limit) * time.Second))
}

func (c *Connection) SetWriteDeadline() {
	c.c.GetRawConn().SetWriteDeadline(time.Now().Add(time.Duration(c.conf.write_limit) * time.Second))
}

func (c *Connection) Close() {
	close(c.exit)
	c.RecvBuffer.Reset()
	c.ticker.Stop()
}

func (c *Connection) Equal(cc *Connection) bool {
	return c.conf.uuid == cc.conf.uuid
}

func (c *Connection) Send(p gotcp.Packet) error {
	if c != nil && c.c != nil {
		return c.c.AsyncWritePacket(p, 0)
	}

	return base.ErrSocketAlreadyNotExist
}
