package player

import (
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

type ActivityObjInfoRefreshHandler interface {
	RefreshInfo(obj *PlayerOpenActivityObject) (err error)
}

type ActivityObjInfoRefreshHandlerFunc func(obj *PlayerOpenActivityObject) (err error)

func (f ActivityObjInfoRefreshHandlerFunc) RefreshInfo(obj *PlayerOpenActivityObject) (err error) {
	return f(obj)
}

var (
	infoRefreshMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]ActivityObjInfoRefreshHandler)
)

func RegisterInfoRefreshHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, h ActivityObjInfoRefreshHandler) {
	subHandlerMap, ok := infoRefreshMap[typ]
	if !ok {
		subHandlerMap = make(map[welfaretypes.OpenActivitySubType]ActivityObjInfoRefreshHandler)
		infoRefreshMap[typ] = subHandlerMap
	}
	_, ok = subHandlerMap[subType]
	if ok {
		panic(fmt.Errorf("welfare:repeat register activityData refresh handler;  type:%d,subType:%d", typ, subType))
	}

	subHandlerMap[subType] = h
}

func GetInfoRefreshHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) ActivityObjInfoRefreshHandler {
	subHandlerMap, ok := infoRefreshMap[typ]
	if !ok {
		return nil
	}
	h, ok := subHandlerMap[subType]
	if !ok {
		return nil
	}

	return h
}
