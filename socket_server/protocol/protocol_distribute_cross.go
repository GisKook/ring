package protocol

import (
	//"encoding/base64"
	"github.com/giskook/ring/pb"
)

type DistributeCrossPkg struct {
	Imei   string
	Serial string
	Cross  string
}

func (d *DistributeCrossPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_CROSS, d.Imei)
	cmd += d.Serial
	cmd += PROTOCOL_SEP
	//cmd += base64.StdEncoding.EncodeToString([]byte(d.Cross))
	cmd += d.Cross
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeCross(d *Carrier.Distribute) (string, *DistributeCrossPkg) {
	return d.Cross.Imei, &DistributeCrossPkg{
		Imei:   d.Cross.Imei,
		Serial: d.Cross.Serial,
		Cross:  d.Cross.Cross,
	}
}
