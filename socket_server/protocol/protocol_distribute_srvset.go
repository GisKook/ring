package protocol

import (
	"github.com/giskook/ring/pb"
)

type DistributeSrvsetPkg struct {
	Imei   string
	Serial string
	Srvset string // 0: shutdown 1: reboot
}

func (d *DistributeSrvsetPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_CTL, d.Imei)
	cmd += d.Serial
	cmd += PROTOCOL_SEP
	cmd += d.Srvset
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeSrvset(d *Carrier.Distribute) (string, *DistributeSrvsetPkg) {
	return d.Srvset.Imei, &DistributeSrvsetPkg{
		Imei:   d.Srvset.Imei,
		Serial: d.Srvset.Serial,
		Srvset: d.Srvset.Srvset,
	}
}
