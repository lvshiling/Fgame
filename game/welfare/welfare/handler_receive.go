package welfare

import (
	"fgame/fgame/game/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

type ReceiveHandler interface {
	ReceiveRew(pl player.Player, rewId int32) error
}

type ReceiveHandlerFunc func(pl player.Player, rewId int32) error

func (h ReceiveHandlerFunc) ReceiveRew(pl player.Player, rewId int32) error {
	return h(pl, rewId)
}

var (
	receiveHandlerMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]ReceiveHandler)
)

// 注册
func RegisterReceiveHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, h ReceiveHandler) {
	subHandlerMap, ok := receiveHandlerMap[typ]
	if !ok {
		subHandlerMap = make(map[welfaretypes.OpenActivitySubType]ReceiveHandler)
		receiveHandlerMap[typ] = subHandlerMap
	}
	_, ok = subHandlerMap[subType]
	if ok {
		panic(fmt.Errorf("welfare:repeat register InfoGet handler;  type:%d,subType:%d", typ, subType))
	}

	subHandlerMap[subType] = h
}

// 获取
func GetReceiveHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) ReceiveHandler {
	subHandlerMap, ok := receiveHandlerMap[typ]
	if !ok {
		return nil
	}
	h, ok := subHandlerMap[subType]
	if !ok {
		return nil
	}

	return h
}
