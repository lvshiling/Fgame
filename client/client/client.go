package client

import (
	"context"
	clientcodec "fgame/fgame/client/codec"
	"fgame/fgame/client/login/login"
	"fgame/fgame/client/processor"
	clientsession "fgame/fgame/client/session"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	tcpsession "fgame/fgame/core/session/tcp"
	"fmt"
	"net"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
)

// //登陆服务器配置
// type LoginServerOptions struct {
// 	Host string
// 	Port int
// }

// //游戏服务器配置
// type GameServerOptions struct {
// 	Host string
// 	Port int
// }

// type ServerOptions struct {
// 	LoginServer *LoginServerOptions `json:"loginServer"`
// 	GameServer  *GameServerOptions  `json:"gameServer"`
// }

// func NewClient(options *ServerOptions, userName string) *Client {
// 	c := &Client{
// 		options:  options,
// 		userName: userName,
// 		done:     make(chan struct{}, 0),
// 	}
// 	return c
// }

// type Client struct {
// 	options     *ServerOptions
// 	accountConn *tcpsession.TCPConnection
// 	conn        *tcpsession.TCPConnection
// 	userName    string
// 	token       string
// 	expiredTime int64
// 	done        chan struct{}
// }

// func (c *Client) Close() {
// 	c.conn.Close()
// }

// func (c *Client) Done() <-chan struct{} {
// 	return c.done
// }

// func (c *Client) Connect() (err error) {
// 	defer func() {
// 		if err != nil {
// 			close(c.done)
// 			return
// 		}
// 	}()
// 	err = c.connectAccountServer()
// 	return
// }

// func (c *Client) connectAccountServer() (err error) {
// 	log.WithFields(
// 		log.Fields{
// 			"userName": c.userName,
// 		}).Infoln("game:客户端正在连接登陆服务器")
// 	loginServerHost := c.options.LoginServer.Host
// 	loginServerPort := c.options.LoginServer.Port
// 	addr := fmt.Sprintf("%s:%d", loginServerHost, loginServerPort)
// 	conn, err := net.Dial("tcp", addr)
// 	if err != nil {
// 		return
// 	}
// 	//TODO启动超时机制

// 	c.accountConn = tcpsession.NewTCPConnection(conn)
// 	sessionOpener := session.SessionHandlerFunc(c.accountSessionOpen)
// 	sessionCloser := session.SessionHandlerFunc(c.accountSessionClose)
// 	sessionReceiver := session.HandlerFunc(c.accountSessionReceive)
// 	sessionSender := session.HandlerFunc(c.accountSessionSend)
// 	handler := tcpsession.NewTCPHandler(sessionOpener, sessionCloser, sessionReceiver, sessionSender)
// 	go func() {
// 		handler.Handle(c.accountConn)
// 	}()
// 	log.WithFields(
// 		log.Fields{
// 			"userName": c.userName,
// 		}).Infoln("game:客户端连接登陆服务器成功")

// 	return
// }

// func (c *Client) accountSessionOpen(s session.Session) error {
// 	ps := clientsession.NewPlayerSession(s, nil)
// 	nctx := clientsession.WithSession(s.Context(), ps)
// 	nctx = WithClient(nctx, c)
// 	s.SetContext(nctx)
// 	c.login(ps)
// 	return nil
// }

// func (c *Client) login(ps clientsession.Session) {
// 	login.AuthLoginServer(ps, c.userName)
// }

// func (c *Client) AccountLogin(token string, expiredTime int64) (err error) {
// 	defer func() {
// 		if err != nil {
// 			close(c.done)
// 		}
// 	}()
// 	c.accountConn.Close()
// 	c.token = token
// 	c.expiredTime = expiredTime
// 	err = c.connectGameServer()
// 	return
// }

// func (c *Client) accountSessionReceive(s session.Session, msg []byte) error {
// 	// handleSessionRecvStats(s, msg)
// 	log.WithFields(
// 		log.Fields{
// 			"sessionId": s.Id(),
// 		}).Debug("对话处理器,接收消息")
// 	//解析
// 	m, err := clientcodec.Decode(msg)
// 	if err != nil {
// 		return nil
// 	}
// 	//处理消息
// 	err = processor.GetDispatch().Handle(s, m)
// 	if err != nil {
// 		switch err.(type) {
// 		case dispatch.HandleError:
// 			// log.WithFields(
// 			// 	log.Fields{
// 			// 		"sessionId": s.Id(),
// 			// 		"err":       err.Error(),
// 			// 	}).Warn("client:处理消息错误")
// 			return nil
// 		}
// 		return err
// 	}
// 	return nil
// }
// func (c *Client) accountSessionClose(s session.Session) error {
// 	ps := clientsession.SessionInContext(s.Context())
// 	if ps == nil {
// 		panic("SessionClose: never reach here")
// 	}

// 	//TODO 移除别的东西
// 	return nil
// }
// func (c *Client) accountSessionSend(s session.Session, msg []byte) error {
// 	return nil
// }

// func (c *Client) connectGameServer() (err error) {
// 	log.WithFields(
// 		log.Fields{
// 			"userName": c.userName,
// 		}).Infoln("game:客户端正在连接游戏服务器")
// 	gameServerHost := c.options.GameServer.Host
// 	gameServerPort := c.options.GameServer.Port
// 	addr := fmt.Sprintf("%s:%d", gameServerHost, gameServerPort)
// 	conn, err := net.Dial("tcp", addr)
// 	if err != nil {
// 		return
// 	}
// 	c.conn = tcpsession.NewTCPConnection(conn)
// 	sessionOpener := session.SessionHandlerFunc(c.SessionOpen)
// 	sessionCloser := session.SessionHandlerFunc(c.SessionClose)
// 	sessionReceiver := session.HandlerFunc(c.SessionReceive)
// 	sessionSender := session.HandlerFunc(c.SessionSend)
// 	handler := tcpsession.NewTCPHandler(sessionOpener, sessionCloser, sessionReceiver, sessionSender)
// 	go func() {
// 		handler.Handle(c.conn)
// 		close(c.done)
// 	}()
// 	log.WithFields(
// 		log.Fields{
// 			"userName": c.userName,
// 		}).Infoln("game:客户端连接游戏服务器成功")

// 	return
// }

// //对话开启
// func (c *Client) SessionOpen(s session.Session) error {
// 	//设置玩家
// 	//TODO 设置session配置
// 	ps := clientsession.NewPlayerSession(s, nil)
// 	nctx := clientsession.WithSession(s.Context(), ps)
// 	s.SetContext(nctx)
// 	c.auth(ps)
// 	return nil
// }

// func (c *Client) auth(ps clientsession.Session) {
// 	login.AuthGameServer(ps, c.userName)
// }

// //对话关闭
// func (c *Client) SessionClose(s session.Session) error {

// 	ps := clientsession.SessionInContext(s.Context())
// 	if ps == nil {
// 		panic("SessionClose: never reach here")
// 	}

// 	//TODO 移除别的东西
// 	return nil
// }

// func (c *Client) SessionReceive(s session.Session, msg []byte) error {

// 	// handleSessionRecvStats(s, msg)
// 	log.WithFields(
// 		log.Fields{
// 			"sessionId": s.Id(),
// 		}).Debug("对话处理器,接收消息")
// 	//解析
// 	m, err := clientcodec.Decode(msg)
// 	if err != nil {
// 		// log.WithFields(
// 		// 	log.Fields{
// 		// 		"sessionId": s.Id(),
// 		// 		"err":       err.Error(),
// 		// 	}).Warn("client:处理消息错误")
// 		return nil
// 	}
// 	//处理消息
// 	err = processor.GetDispatch().Handle(s, m)
// 	if err != nil {
// 		switch err.(type) {
// 		case dispatch.HandleError:
// 			// log.WithFields(
// 			// 	log.Fields{
// 			// 		"sessionId": s.Id(),
// 			// 		"err":       err.Error(),
// 			// 	}).Warn("client:处理消息错误")
// 			return nil
// 		}
// 		return err
// 	}
// 	return nil
// }

// func (c *Client) SessionSend(s session.Session, msg []byte) error {
// 	return nil
// }

// type contextKey string

// var (
// 	clientContextKey contextKey = contextKey("fgame.client")
// )

// func ClientInContext(ctx context.Context) *Client {
// 	val := ctx.Value(clientContextKey)
// 	if val == nil {
// 		return nil
// 	}
// 	c, ok := val.(*Client)
// 	if !ok {
// 		return nil
// 	}
// 	return c
// }

// func WithClient(parent context.Context, c *Client) context.Context {
// 	ctx := context.WithValue(parent, clientContextKey, c)
// 	return ctx
// }

//登陆服务器配置
type LoginServerOptions struct {
	Host string
	Port int
}

//游戏服务器配置
type GameServerOptions struct {
	ServerId int32
	Host     string
	Port     int
}

//客户端选项
type ClientOptions struct {
	LoginServer *LoginServerOptions `json:"loginServer"`
	GameServer  *GameServerOptions  `json:"gameServer"`
}

func NewClient(options *ClientOptions, userName string) (c *Client, err error) {
	c = &Client{
		options:  options,
		userName: userName,
		done:     make(chan struct{}, 0),
	}
	err = c.connect()
	return
}

type Client struct {
	m           sync.Mutex
	options     *ClientOptions
	accountConn *tcpsession.TCPConnection
	conn        *tcpsession.TCPConnection
	userName    string
	token       string
	expiredTime int64
	ctx         context.Context
	cancelFunc  context.CancelFunc
	loginErr    chan error
	loginDone   chan struct{}
	done        chan struct{}
}

func (c *Client) Close() {
	c.cancelFunc()
	c.m.Lock()
	defer c.m.Unlock()
	if c.accountConn != nil {
		c.accountConn.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
	if c.done != nil {
		close(c.done)
		c.done = nil
	}
}

func (c *Client) Done() <-chan struct{} {
	return c.done
}

func (c *Client) connect() (err error) {
	startTime := time.Now().UnixNano()
	c.ctx, c.cancelFunc = context.WithCancel(context.Background())
	c.loginErr = make(chan error, 1)
	c.loginDone = make(chan struct{}, 1)
	err = c.connectAccountServer()
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{
			"costTime": (time.Now().UnixNano() - startTime) / int64(time.Millisecond),
		}).Infoln("game:登陆时间")
	return
}

var (
	loginTimeout = 10 * time.Second
)

func (c *Client) connectAccountServer() (err error) {
	log.WithFields(
		log.Fields{
			"userName": c.userName,
		}).Infoln("game:客户端正在连接登陆服务器")
	loginServerHost := c.options.LoginServer.Host
	loginServerPort := c.options.LoginServer.Port
	addr := fmt.Sprintf("%s:%d", loginServerHost, loginServerPort)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		c.Close()
		return
	}

	c.accountConn = tcpsession.NewTCPConnection(conn)
	sessionOpener := session.SessionHandlerFunc(c.accountSessionOpen)
	sessionCloser := session.SessionHandlerFunc(c.accountSessionClose)
	sessionReceiver := session.HandlerFunc(c.accountSessionReceive)
	sessionSender := session.HandlerFunc(c.accountSessionSend)
	handler := tcpsession.NewTCPHandler(sessionOpener, sessionCloser, sessionReceiver, sessionSender)
	go func() {
		handler.Handle(c.accountConn)
	}()

	loginSucc := false
	waitc := time.After(loginTimeout)
	select {
	case <-c.loginDone:
		loginSucc = true
	case <-c.ctx.Done():
	case <-waitc:
	}
	if !loginSucc {
		err := context.DeadlineExceeded
		select {
		case terr := <-c.loginErr:
			err = terr
		default:
		}
		c.Close()
		return err
	}
	defer func() {
		if err != nil {
			c.Close()
		}
	}()
	log.WithFields(
		log.Fields{
			"userName": c.userName,
		}).Infoln("game:客户端连接登陆服务器成功")

	err = c.connectGameServer()

	return
}

func (c *Client) accountSessionOpen(s session.Session) error {
	ps := clientsession.NewPlayerSession(s, nil)
	nctx := clientsession.WithSession(s.Context(), ps)
	nctx = WithClient(nctx, c)
	s.SetContext(nctx)
	c.login(ps)
	return nil
}

func (c *Client) login(ps clientsession.Session) {
	login.AuthLoginServer(ps, c.userName)
}

func (c *Client) AccountLogin(token string, expiredTime int64) (err error) {

	if c.accountConn != nil {
		c.accountConn.Close()
	}

	c.token = token
	c.expiredTime = expiredTime
	c.loginDone <- struct{}{}
	return
}

func (c *Client) accountSessionReceive(s session.Session, msg []byte) error {
	// handleSessionRecvStats(s, msg)
	log.WithFields(
		log.Fields{
			"sessionId": s.Id(),
		}).Debug("对话处理器,接收消息")
	//解析
	m, err := clientcodec.Decode(msg)
	if err != nil {
		return nil
	}
	//处理消息
	err = processor.GetDispatch().Handle(s, m)
	if err != nil {
		switch err.(type) {
		case dispatch.HandleError:
			// log.WithFields(
			// 	log.Fields{
			// 		"sessionId": s.Id(),
			// 		"err":       err.Error(),
			// 	}).Warn("client:处理消息错误")
			return nil
		}
		return err
	}
	return nil
}
func (c *Client) accountSessionClose(s session.Session) error {
	ps := clientsession.SessionInContext(s.Context())
	if ps == nil {
		panic("SessionClose: never reach here")
	}
	c.accountConn = nil
	//TODO 移除别的东西
	return nil
}
func (c *Client) accountSessionSend(s session.Session, msg []byte) error {
	return nil
}

func (c *Client) connectGameServer() (err error) {
	log.WithFields(
		log.Fields{
			"userName": c.userName,
		}).Infoln("game:客户端正在连接游戏服务器")
	gameServerHost := c.options.GameServer.Host
	gameServerPort := c.options.GameServer.Port
	addr := fmt.Sprintf("%s:%d", gameServerHost, gameServerPort)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	c.conn = tcpsession.NewTCPConnection(conn)
	sessionOpener := session.SessionHandlerFunc(c.sessionOpen)
	sessionCloser := session.SessionHandlerFunc(c.sessionClose)
	sessionReceiver := session.HandlerFunc(c.sessionReceive)
	sessionSender := session.HandlerFunc(c.sessionSend)
	handler := tcpsession.NewTCPHandler(sessionOpener, sessionCloser, sessionReceiver, sessionSender)
	go func() {
		handler.Handle(c.conn)

	}()
	log.WithFields(
		log.Fields{
			"userName": c.userName,
		}).Infoln("game:客户端连接游戏服务器成功")

	return
}

//对话开启
func (c *Client) sessionOpen(s session.Session) error {
	//设置玩家
	//TODO 设置session配置
	ps := clientsession.NewPlayerSession(s, nil)
	nctx := clientsession.WithSession(s.Context(), ps)
	s.SetContext(nctx)
	c.auth(ps)
	return nil
}

func (c *Client) auth(ps clientsession.Session) {
	login.AuthGameServer(ps, c.options.GameServer.ServerId, c.token)
}

//对话关闭
func (c *Client) sessionClose(s session.Session) error {

	ps := clientsession.SessionInContext(s.Context())
	if ps == nil {
		panic("SessionClose: never reach here")
	}
	c.conn = nil
	c.Close()
	//TODO 移除别的东西
	return nil
}

func (c *Client) sessionReceive(s session.Session, msg []byte) error {

	// handleSessionRecvStats(s, msg)
	log.WithFields(
		log.Fields{
			"sessionId": s.Id(),
		}).Debug("对话处理器,接收消息")
	//解析
	m, err := clientcodec.Decode(msg)
	if err != nil {
		// log.WithFields(
		// 	log.Fields{
		// 		"sessionId": s.Id(),
		// 		"err":       err.Error(),
		// 	}).Warn("client:处理消息错误")
		return nil
	}
	//处理消息
	err = processor.GetDispatch().Handle(s, m)
	if err != nil {
		switch err.(type) {
		case dispatch.HandleError:
			// log.WithFields(
			// 	log.Fields{
			// 		"sessionId": s.Id(),
			// 		"err":       err.Error(),
			// 	}).Warn("client:处理消息错误")
			return nil
		}
		return err
	}
	return nil
}

func (c *Client) sessionSend(s session.Session, msg []byte) error {
	return nil
}

type contextKey string

var (
	clientContextKey contextKey = contextKey("fgame.client")
)

func ClientInContext(ctx context.Context) *Client {
	val := ctx.Value(clientContextKey)
	if val == nil {
		return nil
	}
	c, ok := val.(*Client)
	if !ok {
		return nil
	}
	return c
}

func WithClient(parent context.Context, c *Client) context.Context {
	ctx := context.WithValue(parent, clientContextKey, c)
	return ctx
}
