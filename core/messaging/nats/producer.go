package nats

import "github.com/nats-io/nats"

type NatsProducer struct {
	conn *nats.Conn
}

func (np *NatsProducer) Send(subject string, msg []byte) error {
	return np.conn.Publish(subject, msg)
}

func (np *NatsProducer) Close() error {
	np.conn.Close()
	return nil
}

func NewNatsProducer(conn *nats.Conn) *NatsProducer {
	np := &NatsProducer{}
	np.conn = conn

	return np
}
