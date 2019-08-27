package grpc

import (
	"context"
	"sync"

	"fgame/fgame/core/session"
	grpcpb "fgame/fgame/core/session/grpc/pb"

	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/transport"
)

type GrpcSession struct {
	id            string
	conn          grpcpb.Connection_ConnectServer
	inputHandler  session.Handler
	outputHandler session.Handler
	ctx           context.Context
	rwm           sync.RWMutex
	closed        bool
}

func (gs *GrpcSession) Id() string {
	return gs.id
}

//TODO 获取ip
func (gs *GrpcSession) Ip() string {
	return ""
}

func (gs *GrpcSession) Receive(msg []byte) error {
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

func (gs *GrpcSession) Send(msg []byte) error {
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

func (ws *GrpcSession) Context() context.Context {
	ws.rwm.RLock()
	defer ws.rwm.RUnlock()
	if ws.ctx == nil {
		ws.ctx = context.Background()
	}
	return ws.ctx
}

func (ws *GrpcSession) SetContext(ctx context.Context) {
	if ctx == nil {
		panic("nil context")
	}
	ws.rwm.Lock()
	defer ws.rwm.Unlock()
	ws.ctx = ctx
}

func (ws *GrpcSession) Close() error {
	ws.rwm.Lock()
	defer ws.rwm.Unlock()
	if ws.closed {
		return nil
	}
	ws.closed = true
	//TODO 关闭流
	//hack
	stream := grpc.ServerTransportStreamFromContext(ws.conn.Context())
	transportStream := stream.(*transport.Stream)
	s, _ := status.FromError(nil)
	transportStream.ServerTransport().WriteStatus(transportStream, s)
	return nil
}

func NewGrpcSession(conn grpcpb.Connection_ConnectServer, ih session.Handler, oh session.Handler) (gs *GrpcSession, err error) {
	if conn == nil {
		panic("conn should not be nil")
	}
	gs = &GrpcSession{}

	uid := uuid.NewV4()

	gs.id = uid.String()
	gs.conn = conn
	gs.inputHandler = ih
	gs.outputHandler = oh
	return gs, nil
}
