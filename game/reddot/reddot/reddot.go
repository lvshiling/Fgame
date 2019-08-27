package reddot

import (
	"fgame/fgame/game/player"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

func handleDefault(pl player.Player) (bool, bool) {
	return false, false
}

type Handler interface {
	Handle(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) bool
}

type HandlerFunc func(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) bool

func (hf HandlerFunc) Handle(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) bool {
	return hf(pl, obj)
}

func Register(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, h Handler) {
	subTypeMap, exist := redDotHandlerMap[typ]
	if !exist {
		subTypeMap = make(map[welfaretypes.OpenActivitySubType]Handler)
		redDotHandlerMap[typ] = subTypeMap
	}
	_, ok := subTypeMap[subType]
	if ok {
		panic(fmt.Sprintf("repeat register redDotSubType %d", subType))
	}
	subTypeMap[subType] = h
}

func Handle(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isReddot, isChanged bool) {
	if obj == nil {
		return
	}

	typ := obj.GetActivityType()
	subType := obj.GetActivitySubType()
	groupId := obj.GetGroupId()

	if !welfarelogic.IsOnActivityTime(groupId) {
		return false, false
	}

	subTypeMap, exist := redDotHandlerMap[typ]
	if !exist {
		return handleDefault(pl)
	}
	h, exist := subTypeMap[subType]
	if !exist {
		return handleDefault(pl)
	}

	isReddot = h.Handle(pl, obj)
	if obj.GetIsReddot() != isReddot {
		isChanged = true
		obj.SetIsReddot(isReddot)
	}

	return
}

var (
	redDotHandlerMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]Handler)
)
