package session

import (
	"context"
	"fgame/fgame/client/codec"
	"fgame/fgame/core/session"
	"fgame/fgame/pkg/timeutils"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type Player interface {
	Id() int64
}

type SessionState int

const (
	SessionStateInit SessionState = iota
	SessionStateAuth
	SessionStateClosed
)

type Session interface {
	Player() Player
	State() SessionState
	IsInit() bool
	IsAuth() bool
	IsClosed() bool
	Auth(p Player) bool
	Send(msg proto.Message)
	Close()
}

type SessionOptions struct {
	//队列发送超时
	SendTimeout int64
	//队列缓存
	SendMsgQueue int
}

type contextKey string

const (
	sessionKey = contextKey("fgame.client.session")
)

func SessionInContext(ctx context.Context) Session {
	s, ok := ctx.Value(sessionKey).(Session)
	if !ok {
		return nil
	}
	return s
}

func WithSession(ctx context.Context, s Session) context.Context {
	return context.WithValue(ctx, sessionKey, s)
}

//用户对话
type playerSession struct {
	//锁
	m sync.RWMutex
	//用户
	player Player
	//链接
	session session.Session
	//发送消息队列
	msgs chan proto.Message
	done chan struct{}
	//用户状态
	state SessionState
	//配置
	options *SessionOptions
}

func (s *playerSession) State() SessionState {
	return s.state
}

func (s *playerSession) IsInit() bool {
	return s.state == SessionStateInit
}
func (s *playerSession) IsAuth() bool {
	return s.state == SessionStateAuth
}

func (s *playerSession) IsClosed() bool {
	return s.state == SessionStateClosed
}

func (s *playerSession) Player() Player {
	return s.player
}

func (s *playerSession) Auth(pl Player) bool {
	s.m.Lock()
	defer s.m.Unlock()
	//认证成功
	s.player = pl

	return true
}

//可能并发
func (s *playerSession) Close() {
	s.m.Lock()
	defer s.m.Unlock()
	if s.state == SessionStateClosed {
		return
	}
	//关闭发送队列
	close(s.done)
	//关闭连接
	s.session.Close()
	//改变关闭状态
	s.state = SessionStateClosed

}

func (s *playerSession) Send(msg proto.Message) {
	s.m.RLock()
	defer s.m.RUnlock()
	if s.IsClosed() {
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
			log.WithFields(
				log.Fields{
					"player": s.player,
				}).Warn("发送客户端太慢超时")
			s.Close()
		}
	}

}

//启动消息队列
func (s *playerSession) start() {

	go func() {
		now := time.Now().UnixNano()
		count := 0
	Loop:
		for {
			select {
			case m, flag := <-s.msgs:
				{

					if !flag {
						s.Close()
						break Loop
					}
					log.WithFields(log.Fields{
						"msg": m,
					}).Debug("客户端发送消息")
					//TODO:超时重试
					bs, err := codec.Encode(m)
					if err != nil {
						log.WithFields(log.Fields{
							"error": err.Error(),
						}).Error("客户端解析消息错误")
						s.Close()
						break Loop
					}
					err = s.session.Send(bs)
					if err != nil {
						log.WithFields(log.Fields{
							"error": err.Error(),
						}).Error("客户端发送消息错误")
						s.Close()
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
			case <-s.done:
				{
					break Loop
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

func NewPlayerSession(s session.Session, op *SessionOptions) Session {
	ns := &playerSession{}
	if op == nil {
		op = &SessionOptions{
			SendTimeout:  30,
			SendMsgQueue: 10,
		}
	}
	//初始化消息队列结束
	ns.done = make(chan struct{})
	ns.msgs = make(chan proto.Message, op.SendMsgQueue)
	ns.session = s
	ns.options = op
	ns.start()
	return ns
}
