package dispatch

import (
	"errors"
	"fgame/fgame/common/codec"
	"fgame/fgame/core/session"
	"fmt"
)

type Handler interface {
	Handle(session session.Session, msg interface{}) error
}

type HandlerFunc func(session session.Session, msg interface{}) error

func (hf HandlerFunc) Handle(session session.Session, msg interface{}) error {
	return hf(session, msg)
}

type Dispatch interface {
	Handler
	Register(msgType codec.MessageType, h Handler)
}

//TODO 使用sync.Map 以防不正常使用注册机制
type dispatch struct {
	handlerMap map[codec.MessageType]Handler
}

func (d *dispatch) Register(msgType codec.MessageType, h Handler) {
	_, exist := d.handlerMap[msgType]
	if exist {
		panic(fmt.Sprintf("repeat register msgType %d", msgType))
	}
	d.handlerMap[msgType] = h
}

func NewDispatch() Dispatch {
	d := &dispatch{}
	d.handlerMap = make(map[codec.MessageType]Handler)
	return d
}

var (
	ErrorMsgCanNotHandle = errors.New("msg can not handle")
)

type HandleError interface {
	error
	MessageType() codec.MessageType
}

type handleError struct {
	messageType codec.MessageType
}

func (e *handleError) Error() string {
	return fmt.Sprintf("msg [%d] can not handle", int32(e.messageType))
}

func (e *handleError) MessageType() codec.MessageType {
	return e.messageType
}

func newHandleError(msgType codec.MessageType) HandleError {
	e := &handleError{
		messageType: msgType,
	}
	return e
}

func (d *dispatch) Handle(session session.Session, msg interface{}) (err error) {
	m, ok := msg.(*codec.Message)
	if !ok {
		err = ErrorMsgCanNotHandle
		return
	}
	h, exist := d.handlerMap[m.MessageType]
	if !exist {
		err = newHandleError(m.MessageType)
		return
	}
	return h.Handle(session, m.Body)
}

var (
	d Dispatch
)

func init() {
	td := &dispatch{}
	td.handlerMap = make(map[codec.MessageType]Handler)
	d = td
}

func Register(msgType codec.MessageType, h Handler) {
	d.Register(msgType, h)
}

func GetDispatch() Dispatch {
	return d
}
