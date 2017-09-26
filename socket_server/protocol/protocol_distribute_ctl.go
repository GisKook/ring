package protocol

import (
	"github.com/giskook/ring/pb"
)

type DistributeCtlPkg struct {
	Imei   string
	Serial string
	Ctl    string // 0: shutdown 1: reboot
}

func (d *DistributeCtlPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_CTL, d.Imei)
	cmd += d.Serial
	cmd += PROTOCOL_SEP
	cmd += d.Ctl
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeCtl(d *Carrier.Distribute) (string, *DistributeCtlPkg) {
	return d.Ctl.Imei, &DistributeCtlPkg{
		Imei:   d.Ctl.Imei,
		Serial: d.Ctl.Serial,
		Ctl:    d.Ctl.Ctl,
	}
}
