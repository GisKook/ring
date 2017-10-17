package protocol

import (
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/pb/common"
	"github.com/golang/protobuf/proto"
)

type ReportUpresultPkg struct {
	Header *base.Header
	Imei   string
	Serial string
	Result string // now this will always be 1(fail)
}

func (p *ReportUpresultPkg) Serialize() []byte {
	report := &Carrier.Report{
		Header: &Carrier.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Protocol: Carrier.Report_UPRESULT,
		Upresult: &Carrier.ReportUpresult{
			Imei:   p.Imei,
			Serial: p.Serial,
			Result: p.Result,
		},
	}

	data, _ := proto.Marshal(report)

	return data
}

func ParseReportUpresult(p []string, h *base.Header) *ReportUpresultPkg {
	return &ReportUpresultPkg{
		Header: h,
		Imei:   p[2],
		Serial: p[3],
		Result: p[4],
	}
}
