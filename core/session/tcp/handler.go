package tcp

import (
	"io"
	"runtime/debug"

	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
)

//tcp处理器
type TCPHandler struct {
	openHandler   session.SessionHandler
	closeHandler  session.SessionHandler
	inputHandler  session.Handler
	outputHandler session.Handler
}

func (wh *TCPHandler) Handle(conn *TCPConnection) {
	//最后关闭
	defer conn.Close()

	//创建对话
	ts := NewTCPSession(conn, wh.inputHandler, wh.outputHandler)

	if wh.openHandler != nil {
		err := wh.openHandler.Handle(ts)
		if err != nil {
			log.WithFields(
				log.Fields{
					"id":    ts.Id(),
					"ip":    conn.Connection().RemoteAddr().String(),
					"error": err,
				}).Error("TCPHandler:打开错误")
			return
		}
	}

	defer func() {
		//消息处理异常
		if rerr := recover(); rerr != nil {
			debug.PrintStack()
			log.WithFields(
				log.Fields{
					"id":    ts.Id(),
					"ip":    conn.Connection().RemoteAddr().String(),
					"error": rerr,
					"stack": string(debug.Stack()),
				}).Error("TCPHandler:异常")
		}
		//可能关闭处理错误
		defer func() {
			if rerr := recover(); rerr != nil {
				debug.PrintStack()
				log.WithFields(
					log.Fields{
						"id":    ts.Id(),
						"ip":    conn.Connection().RemoteAddr().String(),
						"error": rerr,
						"stack": string(debug.Stack()),
					}).Error("TCPHandler:关闭异常")
			}
		}()
		//对话关闭处理
		if wh.closeHandler != nil {
			err := wh.closeHandler.Handle(ts)
			if err != nil {
				log.WithFields(
					log.Fields{
						"id":    ts.Id(),
						"ip":    conn.Connection().RemoteAddr().String(),
						"error": err,
					}).Error("TCPHandler:关闭错误")
				return
			}
		}
		log.WithFields(
			log.Fields{
				"id": ts.Id(),
				"ip": conn.Connection().RemoteAddr().String(),
			}).Info("TCPHandler:关闭")
	}()
	//对话开启成功
	log.WithFields(
		log.Fields{
			"id": ts.Id(),
			"ip": conn.Connection().RemoteAddr().String(),
		}).Info("TCPHandler:打开")

	var content []byte

	for {

		err := Receive(conn, &content)
		if err != nil {
			if err != io.EOF {
				log.WithFields(
					log.Fields{
						"id":    ts.Id(),
						"ip":    conn.Connection().RemoteAddr().String(),
						"error": err,
					}).Error("TCPHandler:接收,错误")
			}
			break
		}

		//消息处理异常
		if err = ts.Receive(content); err != nil {
			log.WithFields(
				log.Fields{
					"id":    ts.Id(),
					"ip":    conn.Connection().RemoteAddr().String(),
					"error": err.Error(),
				}).Error("TCPHandler:消息处理,错误")
			break
		}
	}
}

//websocket 处理器
func NewTCPHandler(openHandler session.SessionHandler, closeHandler session.SessionHandler, inputHandler session.Handler, outputHandler session.Handler) *TCPHandler {
	wh := &TCPHandler{}
	wh.openHandler = openHandler
	wh.closeHandler = closeHandler
	wh.inputHandler = inputHandler
	wh.outputHandler = outputHandler
	return wh
}
