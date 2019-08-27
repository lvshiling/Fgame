package websocket

import (
	"io"
	"runtime"

	session "fgame/fgame/gm/gamegm/session"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/websocket"
)

//websocket处理器
type WebsocketHandler struct {
	openHandler   session.SessionHandler
	closeHandler  session.SessionHandler
	inputHandler  session.Handler
	outputHandler session.Handler
}

func (wh *WebsocketHandler) Handle(conn *websocket.Conn) {
	//创建对话
	ts := NewWebsocketSession(conn, wh.inputHandler, wh.outputHandler)

	if wh.openHandler != nil {
		err := wh.openHandler.Handle(ts)
		if err != nil {
			log.WithFields(
				log.Fields{
					"id": ts.Id(),
				}).Error("session open failed")
			//conn.Close()
			return
		}
	}

	defer func() {
		//消息处理异常
		if rerr := recover(); rerr != nil {
			stackBuffer := make([]byte, 10240)
			tempPos := runtime.Stack(stackBuffer, true)
			log.WithFields(
				log.Fields{
					"sessionId": ts.Id(),
					"error":     rerr,
					"stack":     string(stackBuffer[:tempPos]),
				}).Error("session 处理异常")
		}
		//对话关闭处理
		if wh.closeHandler != nil {
			err := wh.closeHandler.Handle(ts)
			if err != nil {
				log.WithFields(
					log.Fields{
						"id":    ts.Id(),
						"error": err,
					}).Error("session  close failed")
				return
			}
		}
		log.WithFields(
			log.Fields{
				"id": ts.Id(),
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
		err := websocket.Message.Receive(conn, &content)
		if err != nil {
			if err != io.EOF {
				log.WithFields(
					log.Fields{
						"id":    ts.Id(),
						"error": err,
					}).Error("websocket receive error")
			}
			break
		}

		//消息处理异常
		if err = ts.Receive(content); err != nil {
			log.WithFields(
				log.Fields{
					"id":    ts.Id(),
					"error": err.Error(),
				}).Error("session handle error")
			//	conn.Close()
			break
		}
	}
}

//websocket 处理器
func NewWebsocketHandler(openHandler session.SessionHandler, closeHandler session.SessionHandler, inputHandler session.Handler, outputHandler session.Handler) *WebsocketHandler {
	wh := &WebsocketHandler{}
	wh.openHandler = openHandler
	wh.closeHandler = closeHandler
	wh.inputHandler = inputHandler
	wh.outputHandler = outputHandler
	return wh
}
