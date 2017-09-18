package socket_server

import (
	"github.com/giskook/ring/pb"
	"github.com/giskook/ring/socket_server/protocol"
	"github.com/golang/protobuf/proto"
	"log"
)

func (ss *SocketServer) consumer_worker(c chan []byte) {
	go func() {
		for {
			select {
			case <-ss.exit:
				return
			case p := <-ss.SocketOut:
				distribute := &Carrier.Distribute{}
				err := proto.Unmarshal(p, distribute)
				log.Printf("<IN NSQ> %x %s \n", p, string(p))
				if err != nil {
					log.Println("<ERR> %x unmarshal error", p)
				} else {
					switch distribute.Protocol {
					case Carrier.Distribute_LOGRT:
						ss.Send(protocol.ParseDistributeLogRt(distribute))
					}
				}

			}
		}
	}()
}