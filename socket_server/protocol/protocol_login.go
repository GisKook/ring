package protocol

import (
	"bytes"
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
		Header:   p.Header,
		Protocol: Carrier.Report_LOGIN,
		ReportLogin: &Carrier.ReportLogin{
			Id:         p.ID,
			DeviceType: p.DeviceType,
			Protocol:   p.Protocol,
		},
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseReportLogin(p []string, h *base.Header) *ReqLoginPkg {
	return &ReportLoginPkg{
		Header:          h,
		InnerID:         GetInnerID(p[1]),
		ID:              p[1],
		HardwareVersion: p[2],
		SoftwareVersion: p[3],
	}
}
