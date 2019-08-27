package monitor

import (
	"fgame/fgame/gm/gamegm/basic/pb"
	"fgame/fgame/gm/gamegm/session"
	"fmt"
)

//消息处理器
type Dispatcher struct {
	handlerMap map[int32]MessageHandler
}

func (d *Dispatcher) Register(messageType int32, h MessageHandler) error {
	d.handlerMap[messageType] = h
	return nil
}

func (d *Dispatcher) Handle(s session.Session, msg *pb.Message) error {
	h, exist := d.handlerMap[msg.GetMessageType()]
	if !exist {
		return fmt.Errorf("no exist handler for message type %d", msg.GetMessageType())
	}

	return h.Handle(s, msg)
}

func NewDispatch() *Dispatcher {
	d := &Dispatcher{}
	d.handlerMap = make(map[int32]MessageHandler)
	return d
}
