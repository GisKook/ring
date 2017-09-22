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

func (ss *SocketServer) eh_report_location(p []string) {
	header := &base.Header{
		AppID: ss.conf.AppID,
		From:  ss.conf.UUID,
		To:    ss.conf.Nsq.TopicPLocation,
	}
	location := protocol.ParseReportLocation(p, header)
	// because of the 'TPOSUP' protocol trigger two different way
	// so if 'TPOSUP' is GPS go to the upper server directly
	// if 'TPOSUP' is wifi or cell infos the protocol goto location parser server
	// but in the protobuf the header's finnal dst will always be upper server
	if location.PosType != protocol.LOCATION_TYPE_GPS {
		header.To = ss.conf.Nsq.TopicPLocationParser
	}

	ss.SocketIn <- &base.SocketData{
		Header: header,
		Data:   location.Serialize(),
	}
}
