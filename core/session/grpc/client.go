package grpc

import (
	"context"
	"io"
	"runtime/debug"
	"sync"

	"fgame/fgame/core/session"
	grpcpb "fgame/fgame/core/session/grpc/pb"

	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
)

type GrpcClientSession struct {
	id            string
	conn          grpcpb.Connection_ConnectClient
	inputHandler  session.Handler
	outputHandler session.Handler
	ctx           context.Context
	rwm           sync.RWMutex
	closed        bool
}

func (gs *GrpcClientSession) Id() string {
	return gs.id
}

//TODO 获取ip
func (gs *GrpcClientSession) Ip() string {
	return ""
}

func (gs *GrpcClientSession) Receive(msg []byte) error {
	if gs.inputHandler == nil {
		log.WithFields(log.Fields{
			"id": gs.id,
		}).Warn("session,no input handler")
		return nil
	}

	log.WithFields(log.Fields{
		"id":  gs.id,
		"msg": msg,
	}).Debug("session,msg receive")
	return gs.inputHandler.Handle(gs, msg)
}

func (gs *GrpcClientSession) Send(msg []byte) error {
	if gs.outputHandler != nil {
		err := gs.outputHandler.Handle(gs, msg)
		if err != nil {
			return err
		}
	}
	log.WithFields(log.Fields{
		"id":  gs.id,
		"msg": msg,
	}).Debug("session,msg send")
	// ws.conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	m := &grpcpb.Message{}
	m.Body = msg
	return gs.conn.Send(m)
}

func (ws *GrpcClientSession) Context() context.Context {
	ws.rwm.RLock()
	defer ws.rwm.RUnlock()
	if ws.ctx == nil {
		ws.ctx = context.Background()
	}
	return ws.ctx
}

func (ws *GrpcClientSession) SetContext(ctx context.Context) {
	if ctx == nil {
		panic("nil context")
	}
	ws.rwm.Lock()
	defer ws.rwm.Unlock()
	ws.ctx = ctx
}

func (ws *GrpcClientSession) Close() error {
	ws.rwm.Lock()
	defer ws.rwm.Unlock()
	if ws.closed {
		return nil
	}
	ws.closed = true
	ws.conn.CloseSend()
	return nil
}

func NewGrpcClientSession(conn grpcpb.Connection_ConnectClient, ih session.Handler, oh session.Handler) (gs *GrpcClientSession, err error) {
	if conn == nil {
		panic("conn should not be nil")
	}
	gs = &GrpcClientSession{}

	uid := uuid.NewV4()

	gs.id = uid.String()
	gs.conn = conn
	gs.ctx = conn.Context()
	gs.inputHandler = ih
	gs.outputHandler = oh
	return gs, nil
}

//grpc处理器
type GrpcClientHandler struct {
	openHandler   session.SessionHandler
	closeHandler  session.SessionHandler
	inputHandler  session.Handler
	outputHandler session.Handler
}

func (wh *GrpcClientHandler) Handle(conn grpcpb.Connection_ConnectClient) {

	//创建对话
	ts, err := NewGrpcClientSession(conn, wh.inputHandler, wh.outputHandler)
	if err != nil {
		log.Error("create session  failed")
		return
	}
	if wh.openHandler != nil {
		err := wh.openHandler.Handle(ts)
		if err != nil {
			log.WithFields(
				log.Fields{
					"id": ts.Id(),
					// "ip": conn.RemoteAddr().String(),
				}).Error("session open failed")
			return
		}
	}

	defer func() {
		//消息处理异常
		if rerr := recover(); rerr != nil {
			debug.PrintStack()
			log.WithFields(
				log.Fields{
					"sessionId": ts.Id(),
					// "ip":        conn.RemoteAddr().String(),
					"error": rerr,
					"stack": string(debug.Stack()),
				}).Error("session 处理异常")
		}
		//对话关闭处理
		if wh.closeHandler != nil {
			err := wh.closeHandler.Handle(ts)
			if err != nil {
				log.WithFields(
					log.Fields{
						"id": ts.Id(),
						// "ip":    conn.RemoteAddr().String(),
						"error": err,
					}).Error("session  close failed")
				return
			}
		}
		log.WithFields(
			log.Fields{
				"id": ts.Id(),
				// "ip": conn.RemoteAddr().String(),
			}).Info("session close")
	}()
	//对话开启成功
	log.WithFields(
		log.Fields{
			"id": ts.Id(),
		}).Info("session open")

	var content []byte

	for {
		//TODO 使用内存池取代每次分配
		err := clientReceive(conn, &content)
		if err != nil {
			if err != io.EOF {
				log.WithFields(
					log.Fields{
						"id": ts.Id(),
						// "ip":    conn.RemoteAddr().String(),
						"error": err,
					}).Error("grpc receive error")
			}
			break
		}

		//消息处理异常
		if err = ts.Receive(content); err != nil {
			log.WithFields(
				log.Fields{
					"id": ts.Id(),
					// "ip":    conn.RemoteAddr().String(),
					"error": err.Error(),
				}).Error("session handle error")
			break
		}
	}
}

//websocket 处理器
func NewGrpcClientHandler(openHandler session.SessionHandler, closeHandler session.SessionHandler, inputHandler session.Handler, outputHandler session.Handler) *GrpcClientHandler {
	gh := &GrpcClientHandler{}
	gh.openHandler = openHandler
	gh.closeHandler = closeHandler
	gh.inputHandler = inputHandler
	gh.outputHandler = outputHandler
	return gh
}
