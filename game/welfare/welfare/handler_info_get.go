package welfare

import (
	"fgame/fgame/game/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

type InfoGetHandler interface {
	GetInfo(pl player.Player, groupId int32) error
}

type InfoGetHandlerFunc func(pl player.Player, groupId int32) error

func (h InfoGetHandlerFunc) GetInfo(pl player.Player, groupId int32) error {
	return h(pl, groupId)
}

var (
	infoGetHandlerMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]InfoGetHandler)
)

// 注册
func RegisterInfoGetHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, h InfoGetHandler) {
	subHandlerMap, ok := infoGetHandlerMap[typ]
	if !ok {
		subHandlerMap = make(map[welfaretypes.OpenActivitySubType]InfoGetHandler)
		infoGetHandlerMap[typ] = subHandlerMap
	}
	_, ok = subHandlerMap[subType]
	if ok {
		panic(fmt.Errorf("welfare:repeat register InfoGet handler;  type:%d,subType:%d", typ, subType))
	}

	subHandlerMap[subType] = h
}

// 获取
func GetInfoGetHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) InfoGetHandler {
	subHandlerMap, ok := infoGetHandlerMap[typ]
	if !ok {
		return nil
	}
	h, ok := subHandlerMap[subType]
	if !ok {
		return nil
	}

	return h
}
