package protocol

import (
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/pb"
	"github.com/golang/protobuf/proto"
)

type ReportLoginPkg struct {
	Header     *base.Header
	InnerID    uint64
	ID         string
	DeviceType string
	Protocol   string
}

func (p *ReportLoginPkg) Serialize() []byte {
	report := &Carrier.Report{
		Header: &Carrier.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Protocol: Carrier.Report_LOGIN,
		Login: &Carrier.ReportLogin{
			Id:         p.ID,
			DeviceType: p.DeviceType,
			Protocol:   p.Protocol,
		},
	}

	data, _ := proto.Marshal(report)

	return data
}

func ParseReportLogin(p []string, h *base.Header) *ReportLoginPkg {
	return &ReportLoginPkg{
		Header:     h,
		InnerID:    GetInnerID(p[1]),
		ID:         p[1],
		DeviceType: p[2],
		Protocol:   p[3],
	}
}