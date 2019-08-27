package nats_test

import (
	"fmt"
	"testing"
	"time"

	"fgame/fgame/core/messaging"
	. "fgame/fgame/core/messaging/nats"

	"github.com/nats-io/nats"
)

var (
	p *NatsProducer
)

const (
	url         = "http://192.168.1.13:4222"
	testSubject = "log"
	testData    = "test"
)

func newConn() (conn *nats.Conn, err error) {
	conn, err = nats.Connect(url)
	return
}

// func BenchmarkProduce(b *testing.B) {
// 	conn, err := newConn()
// 	if err != nil {
// 		b.Fatalf("create conn failed %s", err.Error())
// 	}
// 	p := NewNatsProducer(conn)
// 	defer p.Close()
// 	for i := 0; i < b.N; i++ {
// 		err = p.Send(testSubject, []byte(testData))
// 		if err != nil {
// 			b.Fatalf("send msg failed %s", err.Error())
// 		}
// 	}
// }

// func BenchmarkConsume(b *testing.B) {
// 	conn, err := newConn()
// 	if err != nil {
// 		b.Fatalf("create conn failed %s", err.Error())
// 	}

// 	p := NewNatsProducer(conn)
// 	defer p.Close()

// 	cconn, err := newConn()
// 	if err != nil {
// 		b.Fatalf("create consume conn failed %s", err.Error())
// 	}
// 	c := NewNatsConsumer(cconn, testSubject)
// 	err = c.Start(messaging.HandlerFunc(consume))
// 	if err != nil {
// 		b.Fatalf("start consume failed %s", err.Error())
// 	}
// 	defer c.Stop()

// 	for i := 0; i < b.N; i++ {
// 		err = p.Send(testSubject, []byte(testData))
// 		if err != nil {
// 			b.Fatalf("send msg failed %s", err.Error())
// 		}
// 	}
// }

// func TestProduce(t *testing.T) {
// 	err := produce()
// 	if err != nil {
// 		t.Fatalf("send msg failed %s", err.Error())
// 	}
// }

func produce() error {
	conn, err := newConn()
	if err != nil {
		return err
	}
	p := NewNatsProducer(conn)
	defer p.Close()
	err = p.Send(testSubject, []byte(testData))
	return err
}

// func TestProducerClose(t *testing.T) {
// 	conn, err := newConn()
// 	if err != nil {
// 		t.Fatalf("create conn failed %s", err.Error())
// 	}
// 	p := NewNatsProducer(conn)
// 	p.Close()
// }

func TestConsumer(t *testing.T) {
	conn, err := newConn()
	if err != nil {
		t.Fatalf("create conn failed %s", err.Error())
	}
	p := NewNatsConsumer(conn, testSubject)
	defer p.Stop()
	err = p.Start(messaging.HandlerFunc(consume))
	if err != nil {
		t.Fatalf("start consume failed %s", err.Error())
	}
	err = produce()
	if err != nil {
		t.Fatalf("produce failed %s", err.Error())
	}
	time.Sleep(time.Second * 10)
}

func consume(msg []byte) error {

	fmt.Printf("receive msg %s", string(msg))
	return nil
}
