package socket_server

import (
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/socket_server/protocol"
	"strconv"
)

func (ss *SocketServer) eh_report_login(p []string, c *Connection) {
	header := &base.Header{
		AppID: ss.conf.AppID,
		From:  ss.conf.UUID,
		To:    ss.conf.Nsq.TopicPManage,
	}
	login := protocol.ParseReportLogin(p, header)
	id, _ := strconv.ParseUint(login.Imei, 10, 64)
	c.ID = id
	ss.cm.Put(id, c)
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

func (ss *SocketServer) eh_report_lowp(p []string) {
	header := &base.Header{
		AppID: ss.conf.AppID,
		From:  ss.conf.UUID,
		To:    ss.conf.Nsq.TopicPControl,
	}
	lowp := protocol.ParseReportLowp(p, header)

	ss.SocketIn <- &base.SocketData{
		Header: header,
		Data:   lowp.Serialize(),
	}
	ss.Send(protocol.GenPackPkg(lowp.Imei, lowp.Time, protocol.PROTOCOL_REPORT_LOWP_FLAG))
}

func (ss *SocketServer) eh_report_receipt(p []string) {
	header := &base.Header{
		AppID: ss.conf.AppID,
		From:  ss.conf.UUID,
		To:    ss.conf.Nsq.TopicPControl,
	}
	receipt := protocol.ParseReportReceipt(p, header)

	ss.SocketIn <- &base.SocketData{
		Header: header,
		Data:   receipt.Serialize(),
	}
}

func (ss *SocketServer) eh_report_ack(p []string) {
	header := &base.Header{
		AppID: ss.conf.AppID,
		From:  ss.conf.UUID,
		To:    ss.conf.Nsq.TopicPControl,
	}
	ack := protocol.ParseReportAck(p, header)

	ss.SocketIn <- &base.SocketData{
		Header: header,
		Data:   ack.Serialize(),
	}
}

func (ss *SocketServer) eh_report_upresult(p []string) {
	header := &base.Header{
		AppID: ss.conf.AppID,
		From:  ss.conf.UUID,
		To:    ss.conf.Nsq.TopicPControl,
	}
	upresult := protocol.ParseReportUpresult(p, header)

	ss.SocketIn <- &base.SocketData{
		Header: header,
		Data:   upresult.Serialize(),
	}
}
