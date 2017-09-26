package protocol

import (
	"encoding/base64"
	"github.com/giskook/ring/pb"
)

type DistributeUpgradePkg struct {
	Imei   string
	Serial string
	Ip     string
	Port   string
	User   string
	Passwd string
}

func (d *DistributeUpgradePkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_UPGRADE, d.Imei)
	cmd += d.Serial
	cmd += PROTOCOL_SEP
	cmd += d.Ip
	cmd += PROTOCOL_SEP
	cmd += d.Port
	cmd += PROTOCOL_SEP
	cmd += base64.StdEncoding.EncodeToString([]byte(d.User))
	cmd += PROTOCOL_SEP
	cmd += base64.StdEncoding.EncodeToString([]byte(d.Passwd))
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeUpgrade(d *Carrier.Distribute) (string, *DistributeUpgradePkg) {
	return d.Upgrade.Imei, &DistributeUpgradePkg{
		Imei:   d.Upgrade.Imei,
		Serial: d.Upgrade.Serial,
		Ip:     d.Upgrade.Ip,
		Port:   d.Upgrade.Port,
		User:   d.Upgrade.User,
		Passwd: d.Upgrade.Passwd,
	}
}
