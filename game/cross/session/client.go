package session

import (
	"context"
	"fgame/fgame/common/codec"
	"fgame/fgame/core/session"
	"fgame/fgame/pkg/timeutils"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type SessionState int

type SendSession interface {
	Send(msg proto.Message)
	Close(active bool)
}

type SendSessionOptions struct {
	//队列发送超时
	SendTimeout  int64
	SendMsgQueue int32
}

type contextKey string

const (
	sessionKey = contextKey("fgame.game.cross.SendSession")
)

func SendSessionInContext(ctx context.Context) SendSession {
	s, ok := ctx.Value(sessionKey).(SendSession)
	if !ok {
		return nil
	}
	return s
}

func WithSendSession(ctx context.Context, s SendSession) context.Context {
	return context.WithValue(ctx, sessionKey, s)
}

//客户端
type clientSession struct {
	//锁
	m sync.RWMutex
	//链接
	session session.Session
	//发送消息队列
	msgs    chan proto.Message
	options *SendSessionOptions
	//结束
	done     chan struct{}
	isClosed bool
	codec    *codec.Codec
}

//可能并发
func (s *clientSession) Close(active bool) {
	s.m.Lock()
	defer s.m.Unlock()
	if s.isClosed {
		return
	}

	//关闭发送队列
	close(s.msgs)
	//改变关闭状态
	s.isClosed = true

	//关闭连接
	if active {
		s.session.Close()
	}
}

func (s *clientSession) Send(msg proto.Message) {
	s.m.RLock()
	defer s.m.RUnlock()
	if s.isClosed {
		return
	}
	//TODO 缓存
	sendTimeout := timeutils.SecondToDuration(s.options.SendTimeout)
	select {
	//TODO 设置发送消息超时
	case s.msgs <- msg:
		{
			//TODO 是否加个人统计
		}
	case <-time.After(sendTimeout):
		{
			log.WithFields(log.Fields{
				"sessionId": s.session.Id(),
			}).Warn("session:发送客户端太慢超时")
			s.Close(true)
		}
	}

}

//启动消息队列
func (s *clientSession) start() {

	go func() {
		now := time.Now().UnixNano()
		count := 0
	Loop:
		for {
			select {
			case m, flag := <-s.msgs:
				{
					if !flag {
						break Loop
					}

					bs, err := s.codec.Encode(m)
					if err != nil {
						log.WithFields(
							log.Fields{
								"error": err.Error(),
							}).Error("session:客户端压缩消息错误")
						s.Close(true)
						break Loop
					}
					//设置发送超时
					//TODO:超时重试
					err = s.session.Send(bs)
					if err != nil {
						log.WithFields(log.Fields{
							"error": err.Error(),
						}).Error("客户端发送消息错误")
						s.Close(true)
						break Loop
					}
					count++
					curNow := time.Now().UnixNano()
					eclapse := curNow - now
					//限制每秒速度
					if eclapse < int64(time.Second) {
						if count > 1000 {
							time.Sleep(time.Duration(int64(time.Second) - eclapse))
						}
					} else {
						//重置
						now = curNow
						count = 0
					}
				}
			}

		}

		log.WithFields(log.Fields{
			"sessionId": s.session.Id(),
		}).Info("客户端停止")

	}()

	log.WithFields(log.Fields{
		"sessionId": s.session.Id(),
	}).Info("客户端开始")
}

//TODO 临时调整
var (
	defaultSessionOptions = &SendSessionOptions{
		SendTimeout:  30,
		SendMsgQueue: 2000,
	}
)

func NewSendSession(s session.Session, c *codec.Codec, op *SendSessionOptions) SendSession {
	ns := &clientSession{}
	if op == nil {
		op = defaultSessionOptions
	}
	//初始化消息队列结束
	ns.msgs = make(chan proto.Message, op.SendMsgQueue)
	ns.done = make(chan struct{})
	ns.codec = c
	ns.session = s
	ns.options = op

	ns.start()
	return ns
}
