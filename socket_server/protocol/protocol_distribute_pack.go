package protocol

type DistributePackPkg struct {
	Imei    string
	Serial  string // sometimes the time stand for serial
	CmdType string
}

func (d *DistributePackPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_ACK, d.Imei)
	cmd += d.Serial
	cmd += PROTOCOL_SEP
	cmd += d.CmdType
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func GenPackPkg(imei string, serial string, cmdtype string) (string, *DistributePackPkg) {
	return imei, &DistributePackPkg{
		Imei:    imei,
		Serial:  serial,
		CmdType: cmdtype,
	}
}
