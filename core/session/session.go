package session

import "context"

//对话
type Session interface {
	//id uuid生成
	Id() string
	//ip
	Ip() string
	//接收消息 处理
	Receive(msg []byte) error
	//发送消息
	Send(msg []byte) error
	//关闭
	Close() error

	//对话上下文
	Context() context.Context
	//设置上下文
	SetContext(ctx context.Context)
}

//对话消息处理
type Handler interface {
	Handle(s Session, msg []byte) error
}

type HandlerFunc func(s Session, msg []byte) error

func (hf HandlerFunc) Handle(s Session, msg []byte) error {
	return hf(s, msg)
}

//对话消息处理中间件
type MiddlewareHandlerFunc func(next Handler) Handler

func WrapMiddlewareHandler(h Handler) MiddlewareHandlerFunc {
	return func(next Handler) Handler {
		return HandlerFunc(func(s Session, msg []byte) error {
			if err := h.Handle(s, msg); err != nil {
				return err
			}
			return next.Handle(s, msg)
		})
	}
}

func WrapHandler(h Handler, middlewares ...MiddlewareHandlerFunc) Handler {
	return HandlerFunc(func(s Session, msg []byte) error {
		th := h
		for i := len(middlewares) - 1; i >= 0; i-- {
			th = middlewares[i](th)
		}
		return th.Handle(s, msg)
	})
}

//对话处理
type SessionHandler interface {
	Handle(s Session) error
}

type SessionHandlerFunc func(s Session) error

//对话中间件处理
type MiddlewareSessionHandlerFunc func(next SessionHandler) SessionHandler

func WrapMiddlewareSessionHandler(sh SessionHandler) MiddlewareSessionHandlerFunc {
	return func(next SessionHandler) SessionHandler {
		return SessionHandlerFunc(func(s Session) error {
			if err := sh.Handle(s); err != nil {
				return err
			}
			return next.Handle(s)
		})
	}
}

func WrapSessionHandler(h SessionHandler, middlewares ...MiddlewareSessionHandlerFunc) SessionHandler {
	return SessionHandlerFunc(func(s Session) error {
		th := h
		for i := len(middlewares) - 1; i >= 0; i-- {
			th = middlewares[i](th)
		}
		return th.Handle(s)
	})
}

func (shf SessionHandlerFunc) Handle(s Session) error {
	return shf(s)
}
