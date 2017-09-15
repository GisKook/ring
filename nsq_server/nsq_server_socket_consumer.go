package nsq_server

import (
	"github.com/bitly/go-nsq"
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/conf"
)

type nsq_server_socket_c struct {
	conf    *conf.Nsq
	c       *nsq.Consumer
	handler nsq.Handler
}

func new_nsq_server_socket_c(conf *conf.Nsq, handler nsq.Handler) *nsq_server_socket_c {
	nsq_conf := nsq.NewConfig()

	c, err := nsq.NewConsumer(conf.TopicConsumer.Topic, conf.TopicConsumer.Channel, nsq_conf)
	if err != nil {
		return nil
	}

	return &nsq_server_socket_c{
		conf:    conf,
		c:       c,
		handler: handler,
	}
}

func (n *nsq_server_socket_c) connect_to_server() error {
	n.c.AddHandler(n.handler)

	err := n.c.ConnectToNSQD(n.conf.Addr)
	if err != nil {
		return err
	}

	return nil
}

func (n *nsq_server_socket_c) stop() {
	n.c.Stop()
}

func (n *NsqServer) create_consumer_socket() error {
	for i := 0; i < n.conf.TopicConsumer.Number; i++ {
		n._consumers = append(n._consumers, new_nsq_server_socket_c(n.conf, n))
	}
	for _, c := range n._consumers {
		if c != nil {
			err := c.connect_to_server()
			if err != nil {
				return err
			}
		} else {
			return base.ErrNsqConsumerSocketCreateFail
		}
	}

	return nil
}
