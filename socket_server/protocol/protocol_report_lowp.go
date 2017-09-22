package protocol

import (
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/pb"
	"github.com/golang/protobuf/proto"
)

type ReportLowpPkg struct {
	Header *base.Header
	Imei   string
	Time   string
}

func (p *ReportLowpPkg) Serialize() []byte {
	report := &Carrier.Report{
		Header: &Carrier.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Protocol: Carrier.Report_LOWP,
		Lowp: &Carrier.ReportLowp{
			Imei: p.Imei,
			Time: p.Time,
		},
	}

	data, _ := proto.Marshal(report)

	return data
}

func ParseReportLowp(p []string, h *base.Header) *ReportLowpPkg {
	return &ReportLowpPkg{
		Header: h,
		Imei:   p[2],
		Time:   p[3],
	}
}
