package protocol

import (
	"bytes"
	"strconv"
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

	PROTOCOL_LOGIN Protocol = 1
)

var PROTOCOL = map[string]Protocol{
	"LOGIN": PROTOCOL_LOGIN,
}

func Parse(buffer string) []string {
	return strings.Split(buffer, PROTOCOL_SEP)
}

func GetInnerID(id string) uint64 {
	_id, _ := strconv.ParseUint(id, 10, 64)
	return _id
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
