package xiuxianbook

import (
	"fgame/fgame/game/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//TODO:jzy 优化
type CountLevelHandler interface {
	CountLevel(pl player.Player) (level int32, err error)
}

type CountLevelHandlerFunc func(pl player.Player) (level int32, err error)

func (h CountLevelHandlerFunc) CountLevel(pl player.Player) (level int32, err error) {
	return h(pl)
}

var (
	countLevelHandlerMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]CountLevelHandler)
)

// 注册
func RegisterCountLevelHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, h CountLevelHandler) {
	subHandlerMap, ok := countLevelHandlerMap[typ]
	if !ok {
		subHandlerMap = make(map[welfaretypes.OpenActivitySubType]CountLevelHandler)
		countLevelHandlerMap[typ] = subHandlerMap
	}
	_, ok = subHandlerMap[subType]
	if ok {
		panic(fmt.Errorf("welfare:repeat register XiuxianBook count level handler;  type:%d,subType:%d", typ, subType))
	}

	subHandlerMap[subType] = h
}

// 获取
func GetCountLevelHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) CountLevelHandler {
	subHandlerMap, ok := countLevelHandlerMap[typ]
	if !ok {
		return nil
	}
	h, ok := subHandlerMap[subType]
	if !ok {
		return nil
	}

	return h
}
