package protocol

import (
	"github.com/giskook/ring/pb"
)

type DistributeLogRtPkg struct {
	Imei   string
	Time   string
	Result string
}

func (d *DistributeLogRtPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_LOGRT, d.Imei)
	cmd += d.Time
	cmd += PROTOCOL_SEP
	cmd += d.Result
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeLogRt(d *Carrier.Distribute) (string, *DistributeLogRtPkg) {
	return d.Logrt.Id, &DistributeLogRtPkg{
		Imei:   d.Logrt.Id,
		Time:   d.Logrt.Time,
		Result: d.Logrt.Result,
	}
}
