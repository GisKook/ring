package protocol

import (
	"github.com/giskook/ring/pb/common"
)

type DistributeTearchkPkg struct {
	Imei   string
	Serial string
}

func (d *DistributeTearchkPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_TEARCHK, d.Imei)
	cmd += d.Serial
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeTearchk(d *Carrier.Distribute) (string, *DistributeTearchkPkg) {
	return d.Tearchk.Imei, &DistributeTearchkPkg{
		Imei:   d.Tearchk.Imei,
		Serial: d.Tearchk.Serial,
	}
}
