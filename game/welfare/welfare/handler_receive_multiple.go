package welfare

import (
	"fgame/fgame/game/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

type ReceiveMultipleHandler interface {
	ReceiveRewMultiple(pl player.Player, rewId int32, receiveType welfaretypes.ReceiveType) error
}

type ReceiveMultipleHandlerFunc func(pl player.Player, rewId int32, receiveType welfaretypes.ReceiveType) error

func (h ReceiveMultipleHandlerFunc) ReceiveRewMultiple(pl player.Player, rewId int32, receiveType welfaretypes.ReceiveType) error {
	return h(pl, rewId, receiveType)
}

var (
	receiveMultipleHandlerMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]ReceiveMultipleHandler)
)

// 注册
func RegisterReceiveMultipleHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, h ReceiveMultipleHandler) {
	subHandlerMap, ok := receiveMultipleHandlerMap[typ]
	if !ok {
		subHandlerMap = make(map[welfaretypes.OpenActivitySubType]ReceiveMultipleHandler)
		receiveMultipleHandlerMap[typ] = subHandlerMap
	}
	_, ok = subHandlerMap[subType]
	if ok {
		panic(fmt.Errorf("welfare:repeat register InfoGet handler;  type:%d,subType:%d", typ, subType))
	}

	subHandlerMap[subType] = h
}

// 获取
func GetReceiveMultipleHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) ReceiveMultipleHandler {
	subHandlerMap, ok := receiveMultipleHandlerMap[typ]
	if !ok {
		return nil
	}
	h, ok := subHandlerMap[subType]
	if !ok {
		return nil
	}

	return h
}
