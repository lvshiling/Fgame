package proxy

import (
	"fgame/fgame/chatproxy/sdk"
	"fgame/fgame/core/messaging"
	messagingnats "fgame/fgame/core/messaging/nats"

	_ "fgame/fgame/logserver/model"
	logserverpb "fgame/fgame/logserver/pb"
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
	Nats *NatsOptions `json:"nats"`
	Sdk  string       `json:"sdk"`
}

type proxy struct {
	options  *Options
	consumer messaging.Consumer
	wg       sync.WaitGroup
}

func (g *proxy) Init() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("聊天代理数据初始化")
	//初始化sdk
	err = sdk.Init(g.options.Sdk)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s:%d", g.options.Nats.Host, g.options.Nats.Port)
	conn, err := nats.Connect(url)
	if err != nil {
		return errors.Wrap(err, "nats")
	}
	g.consumer = messagingnats.NewNatsConsumer(conn, g.options.Nats.Topic)

	return
}

func (g *proxy) Start() (err error) {
	g.consumer.Start(messaging.HandlerFunc(handleMessage))
	g.wg.Add(1)
	g.wg.Wait()
	return
}

func (g *proxy) Stop() (err error) {
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

func newProxy(options *Options) *proxy {
	g := &proxy{
		options: options,
	}
	return g
}

//单例
var (
	p *proxy
)

//启动服务器
func Init(options *Options) (err error) {
	log.Infoln("正在初始化代理")

	if p != nil {
		panic("重复初始化代理")
	}

	p = newProxy(options)

	err = p.Init()
	if err != nil {
		return
	}

	p.Start()
	log.Infoln("代理初始化成功")

	return
}

func stop() {
	p.Stop()
}
