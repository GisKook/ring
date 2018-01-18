package protocol

import (
	"encoding/base64"
	iconv "github.com/djimenez/iconv-go"
	"github.com/giskook/ring/pb/common"
)

type DistributeMessagePkg struct {
	Imei    string
	Serial  string
	Message string
}

func (d *DistributeMessagePkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_MESSAGE, d.Imei)
	cmd += d.Serial
	cmd += PROTOCOL_SEP
	msg, _ := iconv.ConvertString(d.Message, "UTF-8", "GB18030")
	cmd += base64.StdEncoding.EncodeToString([]byte(msg))
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeMessage(d *Carrier.Distribute) (string, *DistributeMessagePkg) {
	return d.Message.Imei, &DistributeMessagePkg{
		Imei:    d.Message.Imei,
		Serial:  d.Message.Serial,
		Message: d.Message.Message,
	}
}
