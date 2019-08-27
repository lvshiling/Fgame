package monitor

import (
	"fgame/fgame/gm/gamegm/basic/pb"
	"fgame/fgame/gm/gamegm/session"
)

type Message struct {
	s   session.Session
	msg *pb.Message
}

func (m *Message) Session() session.Session {
	return m.s
}

func (m *Message) Msg() *pb.Message {
	return m.msg
}

func NewMessage(s session.Session, msg *pb.Message) *Message {
	m := &Message{
		s:   s,
		msg: msg,
	}
	return m
}

type MessageHandler interface {
	Handle(s session.Session, msg *pb.Message) error
}

type MessageHandlerFunc func(s session.Session, msg *pb.Message) error

func (mhf MessageHandlerFunc) Handle(s session.Session, msg *pb.Message) error {
	return mhf(s, msg)
}

//消息中间件
type MessageHandlerMiddleware interface {
	Handle(s session.Session, msg *pb.Message, next MessageHandler)
}

type MessageHandlerMiddlewareFunc func(s session.Session, msg *pb.Message, next MessageHandler) error

func (mhmf MessageHandlerMiddlewareFunc) Handle(s session.Session, msg *pb.Message, next MessageHandler) error {
	return mhmf(s, msg, next)
}

func WrapMessageHandlers(handlers ...MessageHandler) MessageHandler {
	return MessageHandlerFunc(
		func(s session.Session, msg *pb.Message) error {

			for _, handler := range handlers {
				err := handler.Handle(s, msg)
				if err != nil {
					return err
				}
			}
			return nil
		})
}
