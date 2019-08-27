package welfare

import (
	rankentity "fgame/fgame/game/rank/entity"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

type RankSystemDataHandler interface {
	GetRankList(groupId int32, page int32) ([]*rankentity.PlayerOrderData, int64)
}

type RankSystemDataHandlerFunc func(groupId int32, page int32) ([]*rankentity.PlayerOrderData, int64)

func (h RankSystemDataHandlerFunc) GetRankList(groupId int32, page int32) ([]*rankentity.PlayerOrderData, int64) {
	return h(groupId, page)
}

var (
	rankSystemHandlerMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]RankSystemDataHandler)
)

// 注册
func RegisterRankSystemDataHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, h RankSystemDataHandler) {
	subHandlerMap, ok := rankSystemHandlerMap[typ]
	if !ok {
		subHandlerMap = make(map[welfaretypes.OpenActivitySubType]RankSystemDataHandler)
		rankSystemHandlerMap[typ] = subHandlerMap
	}
	_, ok = subHandlerMap[subType]
	if ok {
		panic(fmt.Errorf("welfare:repeat register rankData handler; type:%d,subType:%d", typ, subType))
	}

	subHandlerMap[subType] = h
}

// 获取
func GetRankSystemDataHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) RankSystemDataHandler {
	subHandlerMap, ok := rankSystemHandlerMap[typ]
	if !ok {
		return nil
	}
	h, ok := subHandlerMap[subType]
	if !ok {
		return nil
	}

	return h
}
