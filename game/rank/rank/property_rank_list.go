package rank

import (
	rankentity "fgame/fgame/game/rank/entity"
	rankobj "fgame/fgame/game/rank/obj"
	ranktypes "fgame/fgame/game/rank/types"
)

func init() {
	RegisterPropertyRankListHandler(ranktypes.RankTypeLevel, PropertyRankListHandlerFunc(GetLevelListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeFeiSheng, PropertyRankListHandlerFunc(GetFeiShengListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeLingTongLevel, PropertyRankListHandlerFunc(GetLingTongLevelListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeCount, PropertyRankListHandlerFunc(GetCountListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeMarryDevelop, PropertyRankListHandlerFunc(GetMarryDevelopListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeCharm, PropertyRankListHandlerFunc(GetCharmListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeCost, PropertyRankListHandlerFunc(GetCostListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeCharge, PropertyRankListHandlerFunc(GetChargeListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeLingTongForce, PropertyRankListHandlerFunc(GetLingTongForceListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeGoldEquipForce, PropertyRankListHandlerFunc(GetGoldEquipForceListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeDianXingForce, PropertyRankListHandlerFunc(GetDianXingForceListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeShenQiForce, PropertyRankListHandlerFunc(GetShenQiForceListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeMingGeForce, PropertyRankListHandlerFunc(GetMingGeForceListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeShengHenForce, PropertyRankListHandlerFunc(GetShengHenForceListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeZhenFaForce, PropertyRankListHandlerFunc(GetZhenFaForceListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeTuLongEquipForce, PropertyRankListHandlerFunc(GetTuLongEquipForceListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeBabyForce, PropertyRankListHandlerFunc(GetBabyForceListByPage))
	RegisterPropertyRankListHandler(ranktypes.RankTypeZhuanSheng, PropertyRankListHandlerFunc(GetZhuanShengListByPage))
}

// 等级排行榜
func GetLevelListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.LevelRank)
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

// 飞升排行榜
func GetFeiShengListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.FeiShengRank)
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

// 灵童等级排行榜
func GetLingTongLevelListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.LingTongLevelRank)
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

// 次数排行榜
func GetCountListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.CountRank)
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

// 表白等级排行榜
func GetMarryDevelopListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.MarryDevelopRank)
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

// 魅力排行榜
func GetCharmListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.CharmRank)
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

// 消费排行榜
func GetCostListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.CostRank)
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

// 充值排行榜
func GetChargeListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.ChargeRank)
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

// 灵童战斗力排行榜
func GetLingTongForceListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.LingTongForceRank)
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

// 元神金装战力排行榜
func GetGoldEquipForceListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.GoldEquipForceRank)
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

// 点星战力排行榜
func GetDianXingForceListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.DianXingForceRank)
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

// 神器战力排行榜
func GetShenQiForceListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.ShenQiForceRank)
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

// 命格战力排行榜
func GetMingGeForceListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.MingGeForceRank)
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

// 圣痕战力排行榜
func GetShengHenForceListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.ShengHenForceRank)
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

// 阵法战力排行榜
func GetZhenFaForceListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.ZhenFaForceRank)
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

// 屠龙装战力排行榜
func GetTuLongEquipForceListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.TuLongEquipForceRank)
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

// 宝宝战力排行榜
func GetBabyForceListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.BabyForceRank)
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

// 转生排行榜
func GetZhuanShengListByPage(rankData rankobj.RankTypeData, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	pageIndex := page * ranktypes.PageLimit
	rankObj, ok := rankData.(*rankobj.ZhuanShengRank)
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
