package protocol

import (
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/pb"
	"github.com/golang/protobuf/proto"
)

type ReportReceiptPkg struct {
	Header *base.Header
	Imei   string
	Serial string
	Time   string
}

func (p *ReportReceiptPkg) Serialize() []byte {
	report := &Carrier.Report{
		Header: &Carrier.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Protocol: Carrier.Report_RECEIPT,
		Receipt: &Carrier.ReportReceipt{
			Imei:   p.Imei,
			Serial: p.Serial,
			Time:   p.Time,
		},
	}

	data, _ := proto.Marshal(report)

	return data
}

func ParseReportReceipt(p []string, h *base.Header) *ReportReceiptPkg {
	return &ReportReceiptPkg{
		Header: h,
		Imei:   p[2],
		Serial: p[3],
		Time:   p[4],
	}
}
