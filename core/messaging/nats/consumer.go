package nats

import (
	"fgame/fgame/core/messaging"

	log "github.com/Sirupsen/logrus"
	"github.com/nats-io/nats"
)

type NatsConsumer struct {
	conn      *nats.Conn
	subject   string
	subscribe *nats.Subscription
	h         messaging.Handler
}

func (nc *NatsConsumer) Start(h messaging.Handler) error {
	log.WithFields(log.Fields{
		"subject": nc.subject,
	}).Info("consumer start")
	nc.h = h
	subscribe, err := nc.conn.Subscribe(nc.subject, nc.handleMsg)
	if err != nil {
		return err
	}
	nc.subscribe = subscribe
	return nil
}

func (nc *NatsConsumer) handleMsg(msg *nats.Msg) {
	err := nc.h.Handle(msg.Data)
	if err != nil {
		log.WithFields(log.Fields{
			"subject": nc.subject,
			"error":   err.Error(),
		}).Warnf("handle msg with error")
	}
}

func (nc *NatsConsumer) Stop() error {
	log.WithFields(log.Fields{
		"subject": nc.subject,
	}).Info("consumer stop")
	err := nc.subscribe.Unsubscribe()
	if err != nil {
		return err
	}
	nc.conn.Close()
	return nil
}

func NewNatsConsumer(conn *nats.Conn, subject string) (nc *NatsConsumer) {
	nc = &NatsConsumer{}
	nc.conn = conn
	nc.subject = subject
	return
}
