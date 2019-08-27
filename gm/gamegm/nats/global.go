package nats

import (
	"fgame/fgame/core/messaging"
	messagingnats "fgame/fgame/core/messaging/nats"
	logserverpb "fgame/fgame/logserver/pb"
	"fgame/fgame/logserver/pbutil"
	"fmt"
	"runtime/debug"
	"sync"

	logserverlog "fgame/fgame/logserver/log"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats"
	"github.com/pkg/errors"
)

type NatsOptions struct {
	Host  string `json:"host"`
	Port  int32  `json:"port"`
	Topic string `json:"topic"`
}

type INatsReceiveEvent interface {
	HandleLog(msg logserverlog.LogMsg) (err error)
}

type NatsGlobal struct {
	options    *NatsOptions
	consumer   messaging.Consumer
	wg         sync.WaitGroup
	eventArray []INatsReceiveEvent
}

func (g *NatsGlobal) Init() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("日志数据初始化")
	//初始化mongo

	url := fmt.Sprintf("%s:%d", g.options.Host, g.options.Port)
	conn, err := nats.Connect(url)
	if err != nil {
		return errors.Wrap(err, "nats")
	}
	g.consumer = messagingnats.NewNatsConsumer(conn, g.options.Topic)

	return
}

func (g *NatsGlobal) Start() (err error) {
	go func() {
		g.consumer.Start(messaging.HandlerFunc(g.handleMessage))
		g.wg.Add(1)
		g.wg.Wait()
		return
	}()
	return nil
}

func (g *NatsGlobal) Stop() (err error) {
	g.wg.Done()
	return
}

func (g *NatsGlobal) OnReceiveLog(p_event INatsReceiveEvent) {
	g.eventArray = append(g.eventArray, p_event)
}

func (g *NatsGlobal) handleMessage(msgBytes []byte) (err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			debug.PrintStack()
			log.WithFields(
				log.Fields{
					"error": rerr,
					"stack": string(debug.Stack()),
				}).Error("log:处理异常")
		}
	}()
	msg := &logserverpb.LogMessage{}
	err = proto.Unmarshal(msgBytes, msg)
	if err != nil {
		return
	}

	err = g.handleLog(msg)
	if err != nil {
		return
	}
	return
}

func (g *NatsGlobal) handleLog(msg *logserverpb.LogMessage) (err error) {
	m, err := pbutil.ConvertFromLogMessage(msg)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("handler:玩家日志转换,失败")
		return
	}
	if m != nil && len(g.eventArray) > 0 {
		for _, value := range g.eventArray {
			err = value.HandleLog(m)
			if err != nil {
				return
			}
		}
	}
	return nil
}

func NewNatsGlobal(options *NatsOptions) *NatsGlobal {
	g := &NatsGlobal{
		options: options,
	}
	g.eventArray = make([]INatsReceiveEvent, 0)
	return g
}
