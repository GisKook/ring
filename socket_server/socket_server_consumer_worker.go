package socket_server

import (
	"github.com/giskook/ring/pb"
	"github.com/giskook/ring/socket_server/protocol"
	"github.com/golang/protobuf/proto"
	"log"
)

func (ss *SocketServer) consumer_worker(c chan []byte) {
	ss.wait_exit.Add(1)
	go func() {
		for {
			select {
			case <-ss.exit:
				ss.wait_exit.Done()
				return
			case p := <-ss.SocketOut:
				distribute := &Carrier.Distribute{}
				err := proto.Unmarshal(p, distribute)
				log.Printf("<IN NSQ> %x %s \n", p, string(p))
				if err != nil {
					log.Println("<ERR> %x unmarshal error\n", p)
				} else {
					var err error
					switch distribute.Protocol {
					case Carrier.Distribute_LOGRT:
						err = ss.Send(protocol.ParseDistributeLogRt(distribute))
					case Carrier.Distribute_LOCATION_RESULT:
						err = ss.Send(protocol.ParseDistributeLocation(distribute))
					case Carrier.Distribute_REQP:
						err = ss.Send(protocol.ParseDistributeReqp(distribute))
					case Carrier.Distribute_TEARCHK:
						err = ss.Send(protocol.ParseDistributeTearchk(distribute))
					case Carrier.Distribute_MESSAGE:
						err = ss.Send(protocol.ParseDistributeMessage(distribute))
					case Carrier.Distribute_CROSS:
						err = ss.Send(protocol.ParseDistributeCross(distribute))
					case Carrier.Distribute_FRESET:
						err = ss.Send(protocol.ParseDistributeFreset(distribute))
					case Carrier.Distribute_LOWPSET:
						err = ss.Send(protocol.ParseDistributeLowpset(distribute))
					case Carrier.Distribute_CTL:
						err = ss.Send(protocol.ParseDistributeCtl(distribute))
					case Carrier.Distribute_SRVSET:
						err = ss.Send(protocol.ParseDistributeSrvset(distribute))
					case Carrier.Distribute_UPGRADE:
						err = ss.Send(protocol.ParseDistributeUpgrade(distribute))
					}

					if err != nil {
						log.Printf("%x %s\n", p, err.Error())
					}
				}

			}
		}
	}()
}
