package grpc

import (
	"io"
	"runtime/debug"

	"fgame/fgame/core/session"
	grpcpb "fgame/fgame/core/session/grpc/pb"

	log "github.com/Sirupsen/logrus"
)

//grpc处理器
type GrpcHandler struct {
	openHandler   session.SessionHandler
	closeHandler  session.SessionHandler
	inputHandler  session.Handler
	outputHandler session.Handler
}

func (wh *GrpcHandler) Handle(conn grpcpb.Connection_ConnectServer) {
	//创建对话
	ts, err := NewGrpcSession(conn, wh.inputHandler, wh.outputHandler)
	if err != nil {
		log.Error("session:create grpc session failed")
		return
	}
	if wh.openHandler != nil {
		err := wh.openHandler.Handle(ts)
		if err != nil {
			log.WithFields(
				log.Fields{
					"id": ts.Id(),
					// "ip": conn.RemoteAddr().String(),
				}).Error("session: grpc session open failed")
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
				}).Error("session: grpc session 处理异常")
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
					}).Error("session:grpc session  close failed")
				return
			}
		}
		log.WithFields(
			log.Fields{
				"id": ts.Id(),
				// "ip": conn.RemoteAddr().String(),
			}).Info("session: grpc session close")
	}()
	//对话开启成功
	log.WithFields(
		log.Fields{
			"id": ts.Id(),
		}).Info("session:grpc session open")

	var content []byte

	for {

		//TODO 使用内存池取代每次分配
		err := serverReceive(conn, &content)
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
func NewGrpcHandler(openHandler session.SessionHandler, closeHandler session.SessionHandler, inputHandler session.Handler, outputHandler session.Handler) *GrpcHandler {
	gh := &GrpcHandler{}
	gh.openHandler = openHandler
	gh.closeHandler = closeHandler
	gh.inputHandler = inputHandler
	gh.outputHandler = outputHandler
	return gh
}
