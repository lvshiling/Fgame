package rank

import (
	rankentity "fgame/fgame/game/rank/entity"
	rankobj "fgame/fgame/game/rank/obj"
	ranktypes "fgame/fgame/game/rank/types"
)

func init() {
	RegisterOrderRankListHandler(ranktypes.RankTypeMount, OrderRankListHandlerFunc(GetMountListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeWing, OrderRankListHandlerFunc(GetWingListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeBodyShield, OrderRankListHandlerFunc(GetBodyShieldListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeLingYu, OrderRankListHandlerFunc(GetLingYuListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeFeather, OrderRankListHandlerFunc(GetFeatherListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeShield, OrderRankListHandlerFunc(GetShieldListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeAnQi, OrderRankListHandlerFunc(GetAnQiListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeFaBao, OrderRankListHandlerFunc(GetFaBaoListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeXianTi, OrderRankListHandlerFunc(GetXianTiListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeShiHunFan, OrderRankListHandlerFunc(GetShiHunListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeTianMoTi, OrderRankListHandlerFunc(GetTianMoTiListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeShenFa, OrderRankListHandlerFunc(GetShenFaListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeLingBao, OrderRankListHandlerFunc(GetLingTongDevListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeLingBing, OrderRankListHandlerFunc(GetLingTongDevListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeLingQi, OrderRankListHandlerFunc(GetLingTongDevListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeLingShen, OrderRankListHandlerFunc(GetLingTongDevListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeLingTi, OrderRankListHandlerFunc(GetLingTongDevListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeLingTongYu, OrderRankListHandlerFunc(GetLingTongDevListByPage))
	RegisterOrderRankListHandler(ranktypes.RankTypeLingYi, OrderRankListHandlerFunc(GetLingTongDevListByPage))

}

// 坐骑排行榜
func GetMountListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.MountRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 战翼排行榜
func GetWingListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.WingRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 护体盾排行榜
func GetBodyShieldListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.BodyShieldRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 领域排行榜
func GetLingYuListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.LingYuRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 护体仙羽排行榜
func GetFeatherListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.FeatherRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 神盾尖刺排行榜
func GetShieldListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.ShieldRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 暗器排行榜
func GetAnQiListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.AnQiRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 法宝排行榜
func GetFaBaoListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.FaBaoRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 仙体排行榜
func GetXianTiListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.XianTiRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 噬魂幡排行榜
func GetShiHunListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.ShiHunFanRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 天魔体排行榜
func GetTianMoTiListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.TianMoTiRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 身法排行榜
func GetShenFaListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.ShenFaRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

// 灵童养成排行榜
func GetLingTongDevListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerOrderData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.LingTongDevRank)
	if !ok {
		return nil, 0
	}

	rankList, rankTime := rankObj.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}
