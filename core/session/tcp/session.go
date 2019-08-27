package tcp

import (
	"context"
	"sync"

	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
)

type TCPSession struct {
	id string
	// conn          *websocket.Conn
	inputHandler  session.Handler
	outputHandler session.Handler
	ctx           context.Context
	rwm           sync.RWMutex
	closed        bool
	conn          *TCPConnection
}

func (s *TCPSession) Id() string {
	return s.id
}

func (s *TCPSession) Ip() string {
	return s.conn.Connection().RemoteAddr().String()
}

func (s *TCPSession) Receive(msg []byte) error {
	if s.inputHandler == nil {
		log.WithFields(log.Fields{
			"id": s.id,
		}).Warn("TCPSession: no input handler")
		return nil
	}

	log.WithFields(log.Fields{
		"id":  s.id,
		"msg": msg,
	}).Debug("TCPSession: msg receive")
	return s.inputHandler.Handle(s, msg)
}

func (s *TCPSession) Send(msg []byte) error {
	if s.outputHandler != nil {
		err := s.outputHandler.Handle(s, msg)
		if err != nil {
			return err
		}
	}
	log.WithFields(log.Fields{
		"id":  s.id,
		"msg": msg,
	}).Debug("TCPSession: msg send")

	err := Send(s.conn, msg)
	if err != nil {
		return err
	}
	return nil
}

func (s *TCPSession) Context() context.Context {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if s.ctx == nil {
		s.ctx = context.Background()
	}
	return s.ctx
}

func (s *TCPSession) SetContext(ctx context.Context) {
	if ctx == nil {
		panic("nil context")
	}
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.ctx = ctx
}

func (s *TCPSession) Close() error {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if s.closed {
		return nil
	}
	s.closed = true
	s.conn.Close()
	return nil
}

func NewTCPSession(conn *TCPConnection, ih session.Handler, oh session.Handler) (s *TCPSession) {
	if conn == nil {
		panic("conn should not be nil")
	}
	s = &TCPSession{}

	uid := uuid.NewV4()

	s.id = uid.String()
	s.conn = conn
	s.inputHandler = ih
	s.outputHandler = oh
	return s
}
