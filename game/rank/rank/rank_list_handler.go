package rank

import (
	rankentity "fgame/fgame/game/rank/entity"
	rankobj "fgame/fgame/game/rank/obj"
	ranktypes "fgame/fgame/game/rank/types"
	"fmt"
)

// PropertyData
type PropertyRankListHandler interface {
	GetPropertyListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64)
}

type PropertyRankListHandlerFunc func(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64)

func (f PropertyRankListHandlerFunc) GetPropertyListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	return f(rankData, page)
}

var (
	propertyListMap = make(map[ranktypes.RankType]PropertyRankListHandler)
)

func RegisterPropertyRankListHandler(rankType ranktypes.RankType, h PropertyRankListHandler) {

	_, ok := propertyListMap[rankType]
	if ok {
		panic(fmt.Errorf("排行榜列表重复注册，类型：%s", rankType.String()))
	}

	propertyListMap[rankType] = h
}

func GetPropertyRankListHandler(rankType ranktypes.RankType) PropertyRankListHandler {
	h, ok := propertyListMap[rankType]
	if !ok {
		return nil
	}

	return h
}

// OrderData
type OrderRankListHandler interface {
	GetOrderListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64)
}

type OrderRankListHandlerFunc func(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64)

func (f OrderRankListHandlerFunc) GetOrderListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	return f(rankData, page)
}

var (
	orderListMap = make(map[ranktypes.RankType]OrderRankListHandler)
)

func RegisterOrderRankListHandler(rankType ranktypes.RankType, h OrderRankListHandler) {

	_, ok := orderListMap[rankType]
	if ok {
		panic(fmt.Errorf("排行榜列表重复注册，类型：%s", rankType.String()))
	}

	orderListMap[rankType] = h
}

func GetOrderRankListHandler(rankType ranktypes.RankType) OrderRankListHandler {
	h, ok := orderListMap[rankType]
	if !ok {
		return nil
	}

	return h
}
