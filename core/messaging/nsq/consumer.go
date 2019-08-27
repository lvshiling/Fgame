package nsq

import (
	"busybird/messaging"

	"runtime"

	"github.com/nsqio/go-nsq"
)

type NSQConsumer struct {
	consumer *nsq.Consumer
	addr     string
	handler  messaging.Handler
}

func (nsqc *NSQConsumer) Start(h messaging.Handler) error {
	err := nsqc.consumer.ConnectToNSQLookupd(nsqc.addr)
	if err != nil {
		return err
	}
	nsqc.handler = h
	return nil
}

func (nsqc *NSQConsumer) Stop() error {
	nsqc.consumer.Stop()
	return nil
}

func (nsqc *NSQConsumer) HandleMessage(msg *nsq.Message) error {
	err := nsqc.handler.Handle(msg.Body)
	if err != nil {
		return err
	}
	return nil
}

func NewConsumer(c *nsq.Consumer, addr string) (nsqc *NSQConsumer) {
	nsqc = &NSQConsumer{}
	nsqc.consumer = c
	nsqc.addr = addr
	n := runtime.GOMAXPROCS(0)
	nsqc.consumer.AddConcurrentHandlers(nsqc, n)
	return
}
