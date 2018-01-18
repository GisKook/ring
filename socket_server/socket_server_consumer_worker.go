package socket_server

import (
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/pb/common"
	"github.com/giskook/ring/pb/lbs_parser"
	"github.com/giskook/ring/socket_server/protocol"
	"github.com/golang/protobuf/proto"
	"log"
)

func (ss *SocketServer) consumer_lbs(l []byte) {
	lbs := &Lbs.Distribute{}
	err := proto.Unmarshal(l, lbs)
	log.Printf("<IN NSQ LBS> %s \n", l)
	if err != nil {
		log.Printf("<ERR> %s unmarshal error\n", l)
	} else {
		header := &base.Header{
			AppID: ss.conf.AppID,
			From:  ss.conf.UUID,
			To:    ss.conf.Nsq.TopicPLocation,
		}
		imei, d := protocol.ParsedDistributeLocationParsed(lbs, header)
		// no need to send to upper any more
		//		if d.ParsedResult == "0" {
		//			ss.SocketIn <- &base.SocketData{
		//				Header: header,
		//				Data:   d.SerializeToUpper(),
		//			}
		//		}
		ss.Send(imei, d)
	}
}

func (ss *SocketServer) consumer_worker() {
	ss.wait_exit.Add(1)
	go func() {
		for {
			select {
			case <-ss.exit:
				ss.wait_exit.Done()
				return
			case l := <-ss.SocketLbsOut:
				ss.consumer_lbs(l)
			case p := <-ss.SocketOut:
				distribute := &Carrier.Distribute{}
				err := proto.Unmarshal(p, distribute)
				log.Printf("<IN NSQ> %s \n", p)
				if err != nil {
					log.Println("<ERR> %s unmarshal error\n", p)
				} else {
					var err error
					switch distribute.Protocol {
					case Carrier.Distribute_LOGRT:
						err = ss.Send(protocol.ParseDistributeLogRt(distribute))
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
