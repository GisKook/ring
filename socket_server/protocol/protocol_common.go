package protocol

import (
	"bytes"
	"strings"
)

type Protocol uint32

const (
	PROTOCOL_MIN_LENGTH int    = 26
	PROTOCOL_START_FLAG string = "$"
	PROTOCOL_END_FLAG   string = "\r\n"
	PROTOCOL_SEP        string = ":"

	PROTOCOL_ILLEGAL   Protocol = 0xffffffff
	PROTOCOL_HALF_PACK Protocol = 0xfffffffe
	PROTOCOL_UNKNOWN   Protocol = 0

	PROTOCOL_REPORT_LOGIN    Protocol = 1
	PROTOCOL_REPORT_LOCATION Protocol = 2
	PROTOCOL_REPORT_LOWP     Protocol = 4
	PROTOCOL_REPORT_RECEIPT  Protocol = 5
	PROTOCOL_REPORT_ACK      Protocol = 6
	PROTOCOL_REPORT_UPRESULT Protocol = 7

	PROTOCOL_DISTRIBUTE_LOGRT    string = "PLOGRT"
	PROTOCOL_DISTRIBUTE_LOCATION string = "PPOSOK"
	PROTOCOL_DISTRIBUTE_ACK      string = "PACK"
	PROTOCOL_DISTRIBUTE_REQP     string = "PREQP"
	PROTOCOL_DISTRIBUTE_TEARCHK  string = "PTEARCHK"
	PROTOCOL_DISTRIBUTE_MESSAGE  string = "PSNDM"
	PROTOCOL_DISTRIBUTE_CROSS    string = "PCROSS"
	PROTOCOL_DISTRIBUTE_FRESET   string = "PFRESET"
	PROTOCOL_DISTRIBUTE_LOWPSET  string = "PLOWPSET"
	PROTOCOL_DISTRIBUTE_CTL      string = "PCTL"
	PROTOCOL_DISTRIBUTE_SRVSET   string = "PSRVSET"
	PROTOCOL_DISTRIBUTE_UPGRADE  string = "PUPGRADE"

	PROTOCOL_REPORT_LOWP_FLAG string = "TLOWP"
)

var PROTOCOL = map[string]Protocol{
	"TLOGIN":                  PROTOCOL_REPORT_LOGIN,
	"TPOSUP":                  PROTOCOL_REPORT_LOCATION,
	PROTOCOL_REPORT_LOWP_FLAG: PROTOCOL_REPORT_LOWP,
	"TMSREAD":                 PROTOCOL_REPORT_RECEIPT,
	"TACK":                    PROTOCOL_REPORT_ACK,
	"TUPRESULT":               PROTOCOL_REPORT_UPRESULT,
}

func Parse(buffer string) []string {
	return strings.Split(buffer, PROTOCOL_SEP)
}

func write_header(protocol_id string, imei string) string {
	cmd := PROTOCOL_START_FLAG + PROTOCOL_SEP
	cmd += protocol_id
	cmd += PROTOCOL_SEP
	cmd += imei
	cmd += PROTOCOL_SEP

	return cmd
}

func CheckProtocol(buffer *bytes.Buffer) (Protocol, []string) {
	cmd := PROTOCOL_ILLEGAL
	var values []string
	bufferlen := buffer.Len()
	if bufferlen == 0 {
		return PROTOCOL_ILLEGAL, nil
	}
	p := string(buffer.Bytes())
	if string(p[0]) != PROTOCOL_START_FLAG {
		buffer.ReadByte()
		cmd, values = CheckProtocol(buffer)
	} else if bufferlen >= PROTOCOL_MIN_LENGTH {
		end_idx := strings.Index(p, PROTOCOL_END_FLAG)
		if end_idx == -1 {
			return PROTOCOL_HALF_PACK, nil
		} else {
			buf, _ := buffer.ReadString('\n')
			values = strings.Split(buf, PROTOCOL_SEP)
			return PROTOCOL[values[1]], values
		}
	} else {
		return PROTOCOL_HALF_PACK, nil
	}

	return cmd, values
}
