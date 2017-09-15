package base

import (
	"errors"
)

var (
	ErrTerminalOffline             = errors.New("terminal is offline")
	ErrNsqConsumerSocketCreateFail = errors.New("Nsq consumer socket create fail")
	ErrNsqProducerSocketCreateFail = errors.New("Nsq producer socket create fail")
)
