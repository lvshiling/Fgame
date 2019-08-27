package server

import (
	"fgame/fgame/core/messaging"
	messagingnats "fgame/fgame/core/messaging/nats"
	coremongo "fgame/fgame/core/mongo"

	logserverlog "fgame/fgame/logserver/log"
	_ "fgame/fgame/logserver/model"
	logserverpb "fgame/fgame/logserver/pb"
	"fgame/fgame/logserver/pbutil"
	"fmt"
	"runtime/debug"
	"sync"

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

//mongo db
//日志服配置
type Options struct {
	Nats  *NatsOptions           `json:"nats"`
	Mongo *coremongo.MongoConfig `json:"mongo"`
}

type global struct {
	options     *Options
	consumer    messaging.Consumer
	mongService coremongo.MongoService
	wg          sync.WaitGroup
}

func (g *global) Init() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("日志数据初始化")
	//初始化mongo
	g.mongService, err = coremongo.NewMongoService(g.options.Mongo)
	if err != nil {
		return errors.Wrap(err, "mongo")
	}
	//初始化日志
	err = g.initLogService()
	if err != nil {
		return errors.Wrap(err, "log")
	}
	url := fmt.Sprintf("%s:%d", g.options.Nats.Host, g.options.Nats.Port)
	conn, err := nats.Connect(url)
	if err != nil {
		return errors.Wrap(err, "nats")
	}
	g.consumer = messagingnats.NewNatsConsumer(conn, g.options.Nats.Topic)

	return
}

func (g *global) initLogService() (err error) {
	err = logserverlog.Init(g.mongService)
	if err != nil {
		return
	}
	return
}

func (g *global) Start() (err error) {
	g.consumer.Start(messaging.HandlerFunc(handleMessage))
	g.wg.Add(1)
	g.wg.Wait()
	return
}

func (g *global) Stop() (err error) {
	g.wg.Done()
	return
}

func handleMessage(msgBytes []byte) (err error) {
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

	err = handleLog(msg)
	if err != nil {
		return
	}
	return
}

func handleLog(msg *logserverpb.LogMessage) (err error) {
	m, err := pbutil.ConvertFromLogMessage(msg)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("handler:玩家日志转换,失败")
		return
	}
	err = logserverlog.GetLogService().AddLog(m)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("handler:玩家添加日志,失败")
		return
	}
	return nil
}

func newGlobal(options *Options) *global {
	g := &global{
		options: options,
	}
	return g
}
