package websocket

import (
	"context"
	"sync"

	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

type WebsocketSession struct {
	id            string
	conn          *websocket.Conn
	inputHandler  session.Handler
	outputHandler session.Handler
	ctx           context.Context
	m             sync.Mutex
	closed        bool
}

func (ws *WebsocketSession) Id() string {
	return ws.id
}

func (ws *WebsocketSession) Ip() string {
	return ws.conn.Request().RemoteAddr
}

func (ws *WebsocketSession) Receive(msg []byte) error {
	if ws.inputHandler == nil {
		log.WithFields(log.Fields{
			"id": ws.id,
		}).Warn("session,no input handler")
		return nil
	}

	log.WithFields(log.Fields{
		"id":  ws.id,
		"msg": msg,
	}).Debug("session,msg receive")
	return ws.inputHandler.Handle(ws, msg)
}

func (ws *WebsocketSession) Send(msg []byte) error {
	if ws.outputHandler != nil {
		err := ws.outputHandler.Handle(ws, msg)
		if err != nil {
			return err
		}
	}
	log.WithFields(log.Fields{
		"id":  ws.id,
		"msg": msg,
	}).Debug("session,msg send")
	// ws.conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	return websocket.Message.Send(ws.conn, msg)
}

func (ws *WebsocketSession) Context() context.Context {
	ws.m.Lock()
	defer ws.m.Unlock()
	if ws.ctx == nil {
		ws.ctx = context.Background()
	}
	return ws.ctx
}

func (ws *WebsocketSession) SetContext(ctx context.Context) {
	if ctx == nil {
		panic("nil context")
	}
	ws.m.Lock()
	defer ws.m.Unlock()
	ws.ctx = ctx
}

func (ws *WebsocketSession) Close() error {
	ws.m.Lock()
	defer ws.m.Unlock()
	if ws.closed {
		return nil
	}
	ws.closed = true
	ws.conn.Close()
	return nil
}

func NewWebsocketSession(conn *websocket.Conn, ih session.Handler, oh session.Handler) (ws *WebsocketSession, err error) {
	if conn == nil {
		panic("conn should not be nil")
	}
	ws = &WebsocketSession{}

	uid := uuid.NewV4()

	ws.id = uid.String()
	ws.conn = conn
	ws.inputHandler = ih
	ws.outputHandler = oh
	return ws, nil
}
