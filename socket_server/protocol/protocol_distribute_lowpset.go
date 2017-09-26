package protocol

import (
	"github.com/giskook/ring/pb"
)

type DistributeLowpsetPkg struct {
	Imei    string
	Serial  string
	Lowpset string // unit 1% eg 30 means 30%
}

func (d *DistributeLowpsetPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_LOWPSET, d.Imei)
	cmd += d.Serial
	cmd += PROTOCOL_SEP
	cmd += d.Lowpset
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeLowpset(d *Carrier.Distribute) (string, *DistributeLowpsetPkg) {
	return d.Lowpset.Imei, &DistributeLowpsetPkg{
		Imei:    d.Lowpset.Imei,
		Serial:  d.Lowpset.Serial,
		Lowpset: d.Lowpset.Lowpset,
	}
}
