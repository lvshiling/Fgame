package monitor

import (
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"fgame/fgame/gm/gamegm/session"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
)

type UserState int

const (
	UserStateInit UserState = iota
	UserStateAuth
	UserStateClosing
	UserStateClosed
)

type UserOptions struct {
	//认证超时秒级
	AuthTimeout int64
	//ping超时
	PingTimeout int64
	//队列发送超时
	SendTimeout int64
	//队列缓存
	SendMsgQueue int
}

//用户
type User struct {
	//锁
	m sync.RWMutex
	//用户id
	id int64
	//ip
	ip string
	//链接
	session session.Session
	//发送消息队列
	msgs chan []byte
	done chan struct{}
	//用户状态
	state UserState
	//配置
	options *UserOptions
	//认证超时
	authTimer *time.Timer
	//ping定时器
	pingTimer *time.Timer
}

func (u *User) State() UserState {
	return u.state
}

func (u *User) IsInit() bool {
	return u.state == UserStateInit
}
func (u *User) IsAuth() bool {
	return u.state == UserStateAuth
}

func (u *User) IsClosed() bool {
	return u.state == UserStateClosed
}

func (u *User) Id() int64 {
	return u.id
}

func (u *User) Ip() string {
	return u.ip
}

func (u *User) Session() session.Session {
	return u.session
}

func (u *User) initAuthTimer() {
	authTimeout := timeutils.SecondToDuration(u.options.AuthTimeout)
	u.authTimer = time.AfterFunc(authTimeout, u.authTimeout)
}

func (u *User) clearAuthTimer() bool {
	return u.authTimer.Stop()
}

func (u *User) authTimeout() {
	log.WithFields(log.Fields{
		"sessionId": u.session.Id(),
	}).Warn("用户认证超时")
	//断开连接
	u.Close()
}

func (u *User) initPingTimer() {
	pingTimeout := time.Duration(u.options.PingTimeout * int64(time.Second))
	u.pingTimer = time.AfterFunc(pingTimeout, u.pingTimeout)
}

func (u *User) clearPingTimer() bool {
	if u.pingTimer == nil {
		return false
	}
	return u.pingTimer.Stop()
}

func (u *User) pingTimeout() {
	log.WithFields(log.Fields{
		"sessionId": u.session.Id(),
	}).Warn("用户ping超时")
	//ping超时
	u.Close()
}

func (u *User) Auth(id int64, ip string) bool {
	u.m.Lock()
	defer u.m.Unlock()
	//认证定时器被停止
	if !u.clearAuthTimer() {
		return false
	}
	//热证成功
	u.id = id
	u.ip = ip
	u.start()
	return true
}

func (u *User) Ping() bool {
	u.m.Lock()
	defer u.m.Unlock()
	if u.state == UserStateClosed {
		return false
	}
	pingTimeout := time.Duration(u.options.PingTimeout * int64(time.Second))
	return u.pingTimer.Reset(pingTimeout)
}

func (u *User) End() {
	u.m.Lock()
	defer u.m.Unlock()
	if u.state == UserStateClosed {
		return
	}
	if u.state == UserStateClosing {
		return
	}
	u.state = UserStateClosing
	close(u.msgs)
}

//可能并发
func (u *User) Close() {
	u.m.Lock()
	defer u.m.Unlock()
	if u.state == UserStateClosed {
		return
	}
	//清楚定时器
	u.clearAuthTimer()
	u.clearPingTimer()
	//关闭发送队列
	close(u.done)
	//关闭连接
	u.session.Close()
	//改变关闭状态
	u.state = UserStateClosed

}

func (u *User) Send(msg []byte) {
	u.m.RLock()
	defer u.m.RUnlock()
	if u.state == UserStateClosed || u.state == UserStateClosing {
		return
	}
	sendTimeout := timeutils.SecondToDuration(u.options.SendTimeout)
	select {
	//TODO 设置发送消息超时
	case u.msgs <- msg:
		{
			//TODO 是否加个人统计
		}
	case <-time.After(sendTimeout):
		{
			log.WithFields(
				log.Fields{
					"userId": u.id,
				}).Warn("client handle too slow")
			u.Close()
		}
	}

}

//启动消息队列
func (u *User) start() {
	if !u.IsInit() {
		panic("user never reach here")
	}
	//TODO 验证用户只启动一次
	u.state = UserStateAuth
	//初始化ping计时器
	u.initPingTimer()

	go func() {
		now := time.Now().UnixNano()
		count := 0
	Loop:
		for {
			select {
			case m, flag := <-u.msgs:
				{
					if !flag {
						u.Close()
						break Loop
					}
					//TODO:超时重试
					err := u.session.Send(m)
					if err != nil {
						log.WithFields(log.Fields{
							"error": err.Error(),
						}).Error("user send message error")
						u.Close()
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
			case <-u.done:
				{
					break Loop
				}
			}

		}

		log.WithFields(log.Fields{
			"sessionId": u.session.Id(),
		}).Info("user stop")

	}()

	log.WithFields(log.Fields{
		"sessionId": u.session.Id(),
	}).Info("user start")
}

func NewUser(s session.Session, op *UserOptions) *User {
	u := &User{}
	if op == nil {
		op = &UserOptions{
			AuthTimeout:  10,
			PingTimeout:  30,
			SendTimeout:  30,
			SendMsgQueue: 100,
		}
	}
	//初始化消息队列结束
	u.done = make(chan struct{})
	u.msgs = make(chan []byte, op.SendMsgQueue)
	u.session = s
	u.options = op
	u.initAuthTimer()

	return u
}
