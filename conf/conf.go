package conf

import (
	"encoding/json"
	"os"
)

type ProducerBase struct {
	Addr   string
	Topic  string
	Number int
}

type Producer struct {
	Manage   *ProducerBase
	Location *ProducerBase
	Control  *ProducerBase
}

type ConsumerBase struct {
	Addr    string
	Topic   string
	Channel string
	Number  int
}

type Consumer struct {
	Manage   *ConsumerBase
	Location *ConsumerBase
	Control  *ConsumerBase
}

type Nsq struct {
	Producer *Producer
	Consumer *Consumer
}

type TcpServer struct {
	BindPort          string
	ReadLimit         uint16
	WriteLimit        uint16
	ConnTimeout       uint16
	ConnCheckInterval uint16
}

type Conf struct {
	AppID     string
	UUID      string
	Nsq       *Nsq
	TcpServer *TcpServer
}

func ReadConfig(confpath string) (*Conf, error) {
	file, _ := os.Open(confpath)
	decoder := json.NewDecoder(file)
	config := Conf{}
	err := decoder.Decode(&config)

	return &config, err
}
