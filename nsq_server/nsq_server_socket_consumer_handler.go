package nsq_server

import (
	"github.com/bitly/go-nsq"
)

func (n *NsqServer) HandleMessage(message *nsq.Message) error {
	n.NsqDataIn <- message.Body
	return nil
}
