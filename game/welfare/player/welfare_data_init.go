package player

import (
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

type ActivityObjInfoInitHandler interface {
	InitInfo(obj *PlayerOpenActivityObject)
}

type ActivityObjInfoInitFunc func(obj *PlayerOpenActivityObject)

func (f ActivityObjInfoInitFunc) InitInfo(obj *PlayerOpenActivityObject) {
	f(obj)
}

var (
	infoInitMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]ActivityObjInfoInitHandler)
)

func RegisterInfoInitHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, h ActivityObjInfoInitHandler) {
	subHandlerMap, ok := infoInitMap[typ]
	if !ok {
		subHandlerMap = make(map[welfaretypes.OpenActivitySubType]ActivityObjInfoInitHandler)
		infoInitMap[typ] = subHandlerMap
	}
	_, ok = subHandlerMap[subType]
	if ok {
		panic(fmt.Errorf("welfare:repeat register activityData init handler; type:%d,subType:%d", typ, subType))
	}

	subHandlerMap[subType] = h
}

func GetInfoInitHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) ActivityObjInfoInitHandler {
	subHandlerMap, ok := infoInitMap[typ]
	if !ok {
		return nil
	}
	h, ok := subHandlerMap[subType]
	if !ok {
		return nil
	}

	return h
}
