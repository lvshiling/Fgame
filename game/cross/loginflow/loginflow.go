package loginflow

import (
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	"fmt"
)

type Handler interface {
	Handle(pl player.Player, crossType crosstypes.CrossType, args ...string) error
}

type HandlerFunc func(pl player.Player, crossType crosstypes.CrossType, args ...string) error

func (hf HandlerFunc) Handle(pl player.Player, crossType crosstypes.CrossType, args ...string) error {
	return hf(pl, crossType, args...)
}

func RegisterCrossLoginFlow(crossType crosstypes.CrossType, h Handler) {
	_, exist := handlerMap[crossType]
	if exist {
		panic(fmt.Sprintf("repeat register crossType %d", crossType))
	}
	handlerMap[crossType] = h
}

func CrossLoginFlowHandle(pl player.Player, crossType crosstypes.CrossType, args ...string) (err error) {
	h, exist := handlerMap[crossType]
	if !exist {
		err = fmt.Errorf("crossType [%d] can not handle", int32(crossType))
		return
	}
	return h.Handle(pl, crossType, args...)
}

var (
	handlerMap = make(map[crosstypes.CrossType]Handler)
)
