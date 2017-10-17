package protocol

import (
	"github.com/giskook/ring/pb/common"
)

type DistributeReqpPkg struct {
	Imei   string
	Serial string
	Type   string
}

func (d *DistributeReqpPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_REQP, d.Imei)
	cmd += d.Serial
	cmd += PROTOCOL_SEP
	cmd += d.Type
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeReqp(d *Carrier.Distribute) (string, *DistributeReqpPkg) {
	return d.Reqp.Imei, &DistributeReqpPkg{
		Imei:   d.Reqp.Imei,
		Serial: d.Reqp.Serial,
		Type:   d.Reqp.Type,
	}
}
