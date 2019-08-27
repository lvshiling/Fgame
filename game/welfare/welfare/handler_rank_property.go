package welfare

import (
	rankentity "fgame/fgame/game/rank/entity"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

type RankPropertyDataHandler interface {
	GetRankList(groupId int32, page int32) ([]*rankentity.PlayerPropertyData, int64)
}

type RankPropertyDataHandlerFunc func(groupId int32, page int32) ([]*rankentity.PlayerPropertyData, int64)

func (h RankPropertyDataHandlerFunc) GetRankList(groupId int32, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	return h(groupId, page)
}

var (
	rankPropertyHandlerMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]RankPropertyDataHandler)
)

// 注册
func RegisterRankPropertyDataHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, h RankPropertyDataHandler) {
	subHandlerMap, ok := rankPropertyHandlerMap[typ]
	if !ok {
		subHandlerMap = make(map[welfaretypes.OpenActivitySubType]RankPropertyDataHandler)
		rankPropertyHandlerMap[typ] = subHandlerMap
	}
	_, ok = subHandlerMap[subType]
	if ok {
		panic(fmt.Errorf("welfare:repeat register rankData handler; type:%d,subType:%d", typ, subType))
	}

	subHandlerMap[subType] = h
}

// 获取
func GetRankPropertyDataHandler(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) RankPropertyDataHandler {
	subHandlerMap, ok := rankPropertyHandlerMap[typ]
	if !ok {
		return nil
	}
	h, ok := subHandlerMap[subType]
	if !ok {
		return nil
	}

	return h
}
