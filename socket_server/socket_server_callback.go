package socket_server

import (
	"github.com/gansidui/gotcp"
	"github.com/giskook/ring/socket_server/protocol"
	"log"
	//"runtime/debug"
	"sync/atomic"
)

func (ss *SocketServer) OnConnect(c *gotcp.Conn) bool {
	connection := NewConnection(c, &ConnConf{
		read_limit:  ss.conf.TcpServer.ReadLimit,
		write_limit: ss.conf.TcpServer.WriteLimit,
		uuid:        atomic.AddUint32(&ss.conn_uuid, 1),
	})

	c.PutExtraData(connection)
	//go connection.Check()
	log.Printf("<CNT> %x \n", c.GetRawConn())

	return true
}

func (ss *SocketServer) OnClose(c *gotcp.Conn) {
	connection := c.GetExtraData().(*Connection)
	ss.cm.Del(connection)
	connection.Close()
	log.Printf("<DIS> %x\n", c.GetRawConn())
	//debug.PrintStack()
}

func (ss *SocketServer) OnMessage(c *gotcp.Conn, p gotcp.Packet) bool {
	connection := c.GetExtraData().(*Connection)
	connection.SetReadDeadline()
	connection.RecvBuffer.Write(p.Serialize())
	for {
		protocol_id, values := protocol.CheckProtocol(connection.RecvBuffer)
		if protocol_id != protocol.PROTOCOL_REPORT_LOGIN && connection != nil && connection.status != USER_STATUS_NORMAL {
			log.Printf("<SWALLOW> %d ", connection.ID)
			return true
		}
		switch protocol_id {
		case protocol.PROTOCOL_HALF_PACK:
			return true
		case protocol.PROTOCOL_ILLEGAL:
			return true
		case protocol.PROTOCOL_REPORT_LOGIN:
			ss.eh_report_login(values, connection)
		case protocol.PROTOCOL_REPORT_LOCATION:
			ss.eh_report_location(values)
		case protocol.PROTOCOL_REPORT_LOWP:
			ss.eh_report_lowp(values)
		case protocol.PROTOCOL_REPORT_RECEIPT:
			ss.eh_report_receipt(values)
		case protocol.PROTOCOL_REPORT_ACK:
			ss.eh_report_ack(values)
		case protocol.PROTOCOL_REPORT_UPRESULT:
			ss.eh_report_upresult(values)
		}
	}
}
