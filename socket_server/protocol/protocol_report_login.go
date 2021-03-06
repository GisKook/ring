package protocol

import (
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/pb/common"
	"github.com/golang/protobuf/proto"
	"time"
)

type ReportLoginPkg struct {
	Header     *base.Header
	Imei       string
	Imsi       string
	DeviceType string
	Protocol   string
	Time       string
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
			Imei:       p.Imei,
			Imsi:       p.Imsi,
			DeviceType: p.DeviceType,
			Protocol:   p.Protocol,
			Time:       p.Time,
		},
	}

	data, _ := proto.Marshal(report)

	return data
}

func ParseReportLogin(p []string, h *base.Header) *ReportLoginPkg {
	return &ReportLoginPkg{
		Header:     h,
		Imei:       p[2],
		Imsi:       p[3],
		DeviceType: p[4],
		Protocol:   p[5],
		Time:       time.Now().Format("060102-150405"),
	}
}
