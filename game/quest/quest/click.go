package quest

import (
	clicktypes "fgame/fgame/game/click/types"
	"fgame/fgame/game/player"
	"fmt"
)

type Handler interface {
	Handle(pl player.Player, clickSubType clicktypes.ClickSubType) error
}

type HandlerFunc func(pl player.Player, clickSubType clicktypes.ClickSubType) error

func (hf HandlerFunc) Handle(pl player.Player, clickSubType clicktypes.ClickSubType) error {
	return hf(pl, clickSubType)
}

func RegisterClick(clickType clicktypes.ClickType, h Handler) {
	_, exist := handlerMap[clickType]
	if exist {
		panic(fmt.Sprintf("repeat register clickType %d", clickType))
	}
	handlerMap[clickType] = h
}

func ClickHandle(pl player.Player, clickType clicktypes.ClickType, clickSubType clicktypes.ClickSubType) (err error) {
	h, exist := handlerMap[clickType]
	if !exist {
		err = fmt.Errorf("clickType [%d] can not handle", int32(clickType))
		return
	}
	return h.Handle(pl, clickSubType)
}

var (
	handlerMap = make(map[clicktypes.ClickType]Handler)
)
