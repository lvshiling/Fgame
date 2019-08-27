package session

import (
	"context"
	"fgame/fgame/common/codec"
	"fgame/fgame/core/session"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"reflect"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

//TODO 测试

type SessionState int

const (
	//初始化
	SessionStateInit SessionState = iota
	//验证
	SessionStateAuth
	//关闭
	SessionStateClosed
)

type Session interface {
	Context() context.Context
	State() SessionState
	Player() Player
	Ip() string
	Auth(p Player) bool
	IsAuth() bool
	Send(msg proto.Message)
	Ping() bool
	Close(active bool)
}

type SessionOptions struct {
	//认证超时秒级
	AuthTimeout int64
	//ping超时
	PingTimeout int64

	//队列缓存
	SendMsgQueue int
}

type contextKey string

const (
	sessionKey = contextKey("fgame.game.session")
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
	//ip
	ip string
	//链接
	session session.Session
	//发送消息队列
	msgs chan proto.Message
	//用户状态
	state SessionState
	//配置
	options *SessionOptions
	//认证超时
	authTimer *time.Timer
	//ping定时器
	pingTimer *time.Timer
	//结束
	done chan struct{}
	//编码器
	codec *codec.Codec
}

func (s *playerSession) State() SessionState {
	return s.state
}

func (s *playerSession) Context() context.Context {
	return s.session.Context()
}

func (s *playerSession) Ip() string {
	return s.session.Ip()
}

func (s *playerSession) Player() Player {
	return s.player
}

func (s *playerSession) IsAuth() bool {
	return s.state == SessionStateAuth
}

func (s *playerSession) initAuthTimer() {
	authTimeout := timeutils.SecondToDuration(s.options.AuthTimeout)
	s.authTimer = time.AfterFunc(authTimeout, s.authTimeout)
}

func (s *playerSession) clearAuthTimer() bool {
	return s.authTimer.Stop()
}

func (s *playerSession) authTimeout() {
	log.WithFields(log.Fields{
		"sessionId": s.session.Id(),
	}).Warn("session:客户端认证超时")
	//TODO 立即断开
	//断开连接
	s.closeImmediate(true)

}

func (s *playerSession) initPingTimer() {
	pingTimeout := time.Duration(s.options.PingTimeout * int64(time.Second))
	s.pingTimer = time.AfterFunc(pingTimeout, s.pingTimeout)
}

func (s *playerSession) clearPingTimer() bool {
	if s.pingTimer == nil {
		return false
	}
	return s.pingTimer.Stop()
}

func (s *playerSession) pingTimeout() {
	log.WithFields(log.Fields{
		"sessionId": s.session.Id(),
	}).Warn("客户端ping超时")
	//ping超时
	s.closeImmediate(true)
}

func (s *playerSession) Auth(pl Player) bool {
	s.m.Lock()
	defer s.m.Unlock()
	if s.state != SessionStateInit {
		return false
	}
	//认证定时器被停止
	if !s.clearAuthTimer() {
		return false
	}
	s.initPingTimer()
	//认证成功
	s.player = pl
	s.state = SessionStateAuth
	return true
}

func (s *playerSession) Ping() bool {
	s.m.Lock()
	defer s.m.Unlock()
	if s.state != SessionStateAuth {
		return false
	}
	pingTimeout := time.Duration(s.options.PingTimeout * int64(time.Second))
	return s.pingTimer.Reset(pingTimeout)
}

//可能并发
func (s *playerSession) Close(active bool) {
	s.m.Lock()
	defer s.m.Unlock()
	if s.state == SessionStateClosed {
		return
	}
	s.close(active, true)
}

//内部使用
func (s *playerSession) closeImmediate(active bool) {
	s.m.Lock()
	defer s.m.Unlock()
	if s.state == SessionStateClosed {
		return
	}
	s.close(active, false)
}

// Lock should be held
func (s *playerSession) close(active bool, graceful bool) {
	log.WithFields(log.Fields{
		"sessionId": s.session.Id(),
	}).Info("session:尝试关闭")
	//清楚定时器
	s.clearAuthTimer()
	s.clearPingTimer()
	//关闭发送队列
	close(s.msgs)
	//改变关闭状态
	s.state = SessionStateClosed
	//通过goroutine等候
	//等待所有消息处理完
	if graceful {
		go s.waitDone(active)
	} else {
		if active {
			s.session.Close()
		}
	}
}

func (s *playerSession) waitDone(active bool) {
	<-s.done
	if active {
		s.session.Close()
	}
}

func (s *playerSession) Send(msg proto.Message) {
	//TODO 优化无锁
	s.m.Lock()
	defer s.m.Unlock()
	if s.state == SessionStateClosed {
		return
	}

	if reflect.ValueOf(msg) == reflect.ValueOf(nil) {
		fmt.Println("asd")
	}
	select {
	//TODO 设置发送消息超时
	case s.msgs <- msg:
		{

			//TODO 是否加个人统计
		}
	default:
		{
			log.WithFields(log.Fields{
				"sessionId": s.session.Id(),
			}).Warn("session:发送客户端太慢超时")
			s.close(true, false)
		}
	}

}

//启动消息队列
func (s *playerSession) start() {

	go func() {
		closed := false
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
								"type":  reflect.TypeOf(m),
								"msg":   m.String(),
								"error": err.Error(),
							}).Error("session:客户端压缩消息错误")
						closed = true
						break Loop
					}

					//设置发送超时
					//TODO:超时重试
					err = s.session.Send(bs)
					if err != nil {
						log.WithFields(log.Fields{
							"msg":   m.String(),
							"error": err.Error(),
						}).Error("客户端发送消息错误")
						closed = true
						break Loop
					}
					count++
					curNow := time.Now().UnixNano()
					eclapse := curNow - now
					//TODO 限制每秒速度
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
		close(s.done)
		if closed {
			s.closeImmediate(true)
			//TODO 立即关闭
			// s.Close(true)
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
	defaultSessionOptions = &SessionOptions{
		AuthTimeout:  30,
		PingTimeout:  210,
		SendMsgQueue: 2000,
	}
)

func NewPlayerSession(s session.Session, c *codec.Codec, op *SessionOptions) Session {
	ns := &playerSession{}
	if op == nil {
		op = defaultSessionOptions
	}
	//初始化消息队列结束
	ns.msgs = make(chan proto.Message, op.SendMsgQueue)
	ns.done = make(chan struct{})
	ns.session = s
	ns.codec = c
	ns.options = op
	ns.initAuthTimer()
	ns.start()
	return ns
}
