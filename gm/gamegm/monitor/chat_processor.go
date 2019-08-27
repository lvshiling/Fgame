package monitor

import (
	"fgame/fgame/gm/gamegm/basic/pb"
	"runtime"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type IChatMsg interface {
	ToProto() *pb.Message
	ServerID() int32
}

type ChatProcessor struct {
	userServer IUserServerManager
	pm         *PlayerManager
	messages   chan IChatMsg
	tick       time.Duration
	closed     bool
	m          sync.Mutex
	done       chan struct{}
}

func (p *ChatProcessor) ReceiveChat(p_chat IChatMsg) {
	p.messages <- p_chat
}

func (p *ChatProcessor) Start() error {
	go func() {
		defer func() {
			//TODO 做补偿
			if rerr := recover(); rerr != nil {
				stackBuffer := make([]byte, 10240)
				tempPos := runtime.Stack(stackBuffer, true)

				log.WithFields(
					log.Fields{
						"error": rerr,
						"stack": string(stackBuffer[:tempPos]),
					}).Error("处理错误")
			}
			p.Stop()
			log.WithFields(
				log.Fields{}).Info("处理器结束")
		}()

	Loop:
		for {
			select {
			case msg, flag := <-p.messages:
				{
					if !flag {
						break Loop
					}
					p.send(msg)
				}
			case <-time.After(p.tick):
				{
				}
			case <-p.done:
				break Loop
			}
		}
	}()
	return nil
}

func (p *ChatProcessor) send(p_chat IChatMsg) error {
	msg := p_chat.ToProto()
	databyte, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	serverId := p_chat.ServerID()
	playerlist := p.userServer.GetServerUserList(serverId)
	for _, value := range playerlist {
		pl := p.pm.GetPlayerById(value)
		if pl == nil {
			continue
		}
		pl.Send(databyte)
	}
	return nil
}

func (p *ChatProcessor) Stop() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.done)
}

func NewChatProcessor(p_userSer IUserServerManager, pm *PlayerManager, queueSize int, tick time.Duration) *ChatProcessor {
	p := &ChatProcessor{}
	p.messages = make(chan IChatMsg, queueSize)
	p.tick = tick
	p.done = make(chan struct{})
	p.userServer = p_userSer
	p.pm = pm
	return p
}
