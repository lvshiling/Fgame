package nsq

import (
	"github.com/nsqio/go-nsq"
)

type NSQProducer struct {
	producer *nsq.Producer
}

func (nsqp *NSQProducer) Send(topic string, msg []byte) error {
	return nsqp.producer.Publish(topic, msg)
}

func (nsqd *NSQProducer) Close() error {
	nsqd.producer.Stop()
	return nil
}

func NewProducer(p *nsq.Producer) (nsqp *NSQProducer) {
	nsqp = &NSQProducer{}
	nsqp.producer = p
	return
}
