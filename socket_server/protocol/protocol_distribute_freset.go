package protocol

import (
	"github.com/giskook/ring/pb"
)

type DistributeFresetPkg struct {
	Imei   string
	Serial string
	Freset string // unit: minute
}

func (d *DistributeFresetPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_FRESET, d.Imei)
	cmd += d.Serial
	cmd += PROTOCOL_SEP
	cmd += d.Freset
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeFreset(d *Carrier.Distribute) (string, *DistributeFresetPkg) {
	return d.Freset.Imei, &DistributeFresetPkg{
		Imei:   d.Freset.Imei,
		Serial: d.Freset.Serial,
		Freset: d.Freset.Frequency,
	}
}
