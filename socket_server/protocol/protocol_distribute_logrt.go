package protocol

import (
	"github.com/giskook/ring/pb/common"
	"strconv"
)

type DistributeLogRtPkg struct {
	Imei     string
	Time     string
	Result   string
	Settings map[string]string
}

func (d *DistributeLogRtPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_LOGRT, d.Imei)
	cmd += d.Time
	cmd += PROTOCOL_SEP

	setting_len := len(d.Settings)
	var cur_inx int
	for k, v := range d.Settings {
		cmd += k + "," + v
		cur_inx, _ = strconv.Atoi(k)
		if cur_inx < setting_len-1 {
			cmd += ";"
		}
	}
	cmd += PROTOCOL_SEP
	cmd += d.Result
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeLogRt(d *Carrier.Distribute) (string, *DistributeLogRtPkg) {
	dlogrt := &DistributeLogRtPkg{
		Imei:     d.Logrt.Imei,
		Time:     d.Logrt.Time,
		Result:   d.Logrt.Result,
		Settings: make(map[string]string),
	}
	for k, v := range d.Logrt.Settings {
		dlogrt.Settings[k] = v
	}

	return d.Logrt.Imei, dlogrt
}
