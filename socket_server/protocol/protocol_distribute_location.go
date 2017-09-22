package protocol

import (
	"github.com/giskook/ring/pb"
	"github.com/golang/protobuf/proto"
)

type DistributeLocationPkg struct {
	Imei        string
	Time        string
	PosReason   string
	ParseResult string
}

func (d *DistributeLocationPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_LOCATION, d.Imei)
	cmd += d.Time + PROTOCOL_SEP
	cmd += d.PosReason + PROTOCOL_SEP
	cmd += d.ParseResult
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeLocation(d *Carrier.Distribute) (string, *DistributeLocationPkg) {
	extra := &Carrier.LocationExtra{}
	proto.Unmarshal(d.LocationResult.Extra, extra)
	return extra.Imei, &DistributeLocationPkg{
		Imei:        extra.Imei,
		Time:        extra.Time,
		PosReason:   extra.PosReason,
		ParseResult: d.LocationResult.ParseResult,
	}
}
