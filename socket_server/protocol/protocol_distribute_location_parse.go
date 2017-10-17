package protocol

import (
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/pb/common"
	"github.com/giskook/ring/pb/lbs_parser"
	"github.com/golang/protobuf/proto"
)

type DistributeLocationParsedPkg struct {
	Header *base.Header

	Imei         string
	Time         string
	PosReason    string
	ParsedResult string

	Longitude string
	Latitude  string
	Extra     []byte
}

func (d *DistributeLocationParsedPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_LOCATION, d.Imei)
	cmd += d.Time + PROTOCOL_SEP
	cmd += d.PosReason + PROTOCOL_SEP
	cmd += d.ParsedResult
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func (p *DistributeLocationParsedPkg) SerializeToUpper() []byte {
	report := &Carrier.Report{
		Header: &Carrier.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Protocol: Carrier.Report_LOCATION,
		Location: &Carrier.ReportLocation{
			Extra:     p.Extra,
			Longitude: p.Longitude,
			Latitude:  p.Latitude,
		},
	}
	data, _ := proto.Marshal(report)

	return data
}

func ParsedDistributeLocationParsed(d *Lbs.Distribute, header *base.Header) (string, *DistributeLocationParsedPkg) {
	extra := &Carrier.LocationExtra{}
	proto.Unmarshal(d.Extra, extra)
	return extra.Imei, &DistributeLocationParsedPkg{
		Header:       header,
		Imei:         extra.Imei,
		Time:         extra.Time,
		PosReason:    extra.PosReason,
		ParsedResult: d.ParseResult,
		Longitude:    d.Longitude,
		Latitude:     d.Latitude,
		Extra:        d.Extra,
	}
}
