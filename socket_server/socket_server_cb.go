package socket_server

import (
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/socket_server/protocol"
)

func (ss *SocketServer) eh_report_login(p []string, c *Connection) {
	header := &base.Header{
		AppID: ss.conf.AppID,
		From:  ss.conf.UUID,
		To:    ss.conf.Nsq.TopicPManage,
	}
	login := protocol.ParseReportLogin(p, header)
	c.ID = login.InnerID
	ss.cm.Put(login.InnerID, c)
	ss.SocketIn <- &base.SocketData{
		Header: header,
		Data:   login.Serialize(),
	}
}
