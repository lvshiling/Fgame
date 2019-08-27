package relate_handler

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/game/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

type RelateHandler interface {
	AttendDrewRelate(pl player.Player, relateGroupId int32, msg *uipb.SCOpenActivityChargeDrewAttend)
}

type RelateHandlerFunc func(pl player.Player, relateGroupId int32, msg *uipb.SCOpenActivityChargeDrewAttend)

func (f RelateHandlerFunc) AttendDrewRelate(pl player.Player, relateGroupId int32, msg *uipb.SCOpenActivityChargeDrewAttend) {
	f(pl, relateGroupId, msg)
}

var (
	realteHandlerMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]RelateHandler)
)

func RegistRelateHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, h RelateHandler) {
	subHandlerMap, ok := realteHandlerMap[typ]
	if !ok {
		subHandlerMap = make(map[welfaretypes.OpenActivitySubType]RelateHandler)
		realteHandlerMap[typ] = subHandlerMap
	}
	_, ok = subHandlerMap[subType]
	if ok {
		panic(fmt.Errorf("welfare:repeat register rankData handler; type:%d,subType:%d", typ, subType))
	}

	subHandlerMap[subType] = h
}

func GetRelateHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) RelateHandler {
	subHandlerMap, ok := realteHandlerMap[typ]
	if !ok {
		return nil
	}
	h, ok := subHandlerMap[subType]
	if !ok {
		return nil
	}

	return h
}
