package grpc_pbutil

import (
	rankentity "fgame/fgame/game/rank/entity"
	rankobj "fgame/fgame/game/rank/obj"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

func BuildForceInfoList(rankMap *rankobj.RankMap) (rankInfo *rankpb.AreaRankForce) {
	rankTypeData := rankMap.GetRankTypeData(ranktypes.RankTypeForce)
	rankData, ok := rankTypeData.(*rankobj.ForceRank)
	if !ok {
		return nil
	}
	rankList, rankTime := rankData.GetListAndTime()
	rankInfo = &rankpb.AreaRankForce{}
	rankInfo.RankTime = rankTime
	for _, rankObj := range rankList {
		rankInfo.ForceList = append(rankInfo.ForceList, buildForce(rankObj))
	}
	return rankInfo
}

func buildForce(data *rankentity.PlayerForceData) *rankpb.RankForceInfo {
	rankInfo := &rankpb.RankForceInfo{}
	rankInfo.ServerId = data.ServerId
	rankInfo.PlayerId = data.PlayerId
	rankInfo.PlayerName = data.PlayerName
	rankInfo.GangName = data.GangName
	rankInfo.Power = data.Force
	rankInfo.Role = int32(data.Role)
	rankInfo.Sex = int32(data.Sex)
	return rankInfo
}

func BuildGangInfoList(rankMap *rankobj.RankMap) (rankInfo *rankpb.AreaRankGang) {
	rankTypeData := rankMap.GetRankTypeData(ranktypes.RankTypeGang)
	rankData, ok := rankTypeData.(*rankobj.GangRank)
	if !ok {
		return nil
	}
	rankList, rankTime := rankData.GetListAndTime()
	rankInfo = &rankpb.AreaRankGang{}
	rankInfo.RankTime = rankTime
	for _, rankObj := range rankList {
		rankInfo.GangList = append(rankInfo.GangList, buildGang(rankObj))
	}
	return rankInfo
}

func buildGang(data *rankentity.PlayerGangData) *rankpb.RankGangInfo {
	rankInfo := &rankpb.RankGangInfo{}
	rankInfo.ServerId = data.ServerId
	rankInfo.GangId = data.GangId
	rankInfo.GangName = data.GangName
	rankInfo.LeaderId = data.LeadId
	rankInfo.LeaderName = data.LeadName
	rankInfo.Power = data.Power
	rankInfo.Role = int32(data.Role)
	rankInfo.Sex = int32(data.Sex)
	return rankInfo
}

func BuildMountInfoList(rankMap *rankobj.RankMap) (rankInfo *rankpb.AreaRankMount) {
	rankTypeData := rankMap.GetRankTypeData(ranktypes.RankTypeMount)
	rankData, ok := rankTypeData.(*rankobj.MountRank)
	if !ok {
		return nil
	}
	rankList, rankTime := rankData.GetListAndTime()
	rankInfo = &rankpb.AreaRankMount{}
	rankInfo.RankTime = rankTime
	for _, rankObj := range rankList {
		rankInfo.MountList = append(rankInfo.MountList, buildOrder(rankObj))
	}
	return rankInfo
}

func buildOrder(data *rankentity.PlayerOrderData) *rankpb.RankOrderInfo {
	rankInfo := &rankpb.RankOrderInfo{}
	rankInfo.ServerId = data.ServerId
	rankInfo.PlayerId = data.PlayerId
	rankInfo.PlayerName = data.PlayerName
	rankInfo.Order = data.Order
	rankInfo.Power = data.Power
	return rankInfo
}

func BuildWingInfoList(rankMap *rankobj.RankMap) (rankInfo *rankpb.AreaRankWing) {
	rankTypeData := rankMap.GetRankTypeData(ranktypes.RankTypeWing)
	rankData, ok := rankTypeData.(*rankobj.WingRank)
	if !ok {
		return nil
	}
	rankList, rankTime := rankData.GetListAndTime()
	rankInfo = &rankpb.AreaRankWing{}
	rankInfo.RankTime = rankTime
	for _, rankObj := range rankList {
		rankInfo.WingList = append(rankInfo.WingList, buildOrder(rankObj))
	}
	return rankInfo
}

func BuildWeaponInfoList(rankMap *rankobj.RankMap) (rankInfo *rankpb.AreaRankWeapon) {
	rankTypeData := rankMap.GetRankTypeData(ranktypes.RankTypeWeapon)
	rankData, ok := rankTypeData.(*rankobj.WeaponRank)
	if !ok {
		return nil
	}
	rankList, rankTime := rankData.GetListAndTime()
	rankInfo = &rankpb.AreaRankWeapon{}
	rankInfo.RankTime = rankTime
	for _, rankObj := range rankList {
		rankInfo.WeaponList = append(rankInfo.WeaponList, buildWeapon(rankObj))
	}
	return rankInfo
}

func buildWeapon(data *rankentity.PlayerWeaponData) *rankpb.RankWeaponInfo {
	rankInfo := &rankpb.RankWeaponInfo{}
	rankInfo.ServerId = data.ServerId
	rankInfo.PlayerId = data.PlayerId
	rankInfo.PlayerName = data.PlayerName
	rankInfo.Star = data.Star
	rankInfo.Power = data.Power
	rankInfo.Sex = int32(data.Sex)
	rankInfo.Role = int32(data.Role)
	return rankInfo
}

func BuildBodyShieldInfoList(rankMap *rankobj.RankMap) (rankInfo *rankpb.AreaRankBodyShield) {
	rankTypeData := rankMap.GetRankTypeData(ranktypes.RankTypeBodyShield)
	rankData, ok := rankTypeData.(*rankobj.BodyShieldRank)
	if !ok {
		return nil
	}
	rankList, rankTime := rankData.GetListAndTime()
	rankInfo = &rankpb.AreaRankBodyShield{}
	rankInfo.RankTime = rankTime
	for _, rankObj := range rankList {
		rankInfo.BodyShieldList = append(rankInfo.BodyShieldList, buildOrder(rankObj))
	}
	return rankInfo
}

func BuildShenFaInfoList(rankMap *rankobj.RankMap) (rankInfo *rankpb.AreaRankShenFa) {
	rankTypeData := rankMap.GetRankTypeData(ranktypes.RankTypeShenFa)
	rankData, ok := rankTypeData.(*rankobj.ShenFaRank)
	if !ok {
		return nil
	}
	rankList, rankTime := rankData.GetListAndTime()
	rankInfo = &rankpb.AreaRankShenFa{}
	rankInfo.RankTime = rankTime
	for _, rankObj := range rankList {
		rankInfo.ShenFaList = append(rankInfo.ShenFaList, buildOrder(rankObj))
	}
	return rankInfo
}

func BuildLingYuInfoList(rankMap *rankobj.RankMap) (rankInfo *rankpb.AreaRankLingYu) {
	rankTypeData := rankMap.GetRankTypeData(ranktypes.RankTypeLingYu)
	rankData, ok := rankTypeData.(*rankobj.LingYuRank)
	if !ok {
		return nil
	}
	rankList, rankTime := rankData.GetListAndTime()
	rankInfo = &rankpb.AreaRankLingYu{}
	rankInfo.RankTime = rankTime
	for _, rankObj := range rankList {
		rankInfo.LingYuList = append(rankInfo.LingYuList, buildOrder(rankObj))
	}
	return rankInfo
}

func BuildFeatherInfoList(rankMap *rankobj.RankMap) (rankInfo *rankpb.AreaRankFeather) {
	rankTypeData := rankMap.GetRankTypeData(ranktypes.RankTypeFeather)
	rankData, ok := rankTypeData.(*rankobj.FeatherRank)
	if !ok {
		return nil
	}
	rankList, rankTime := rankData.GetListAndTime()
	rankInfo = &rankpb.AreaRankFeather{}
	rankInfo.RankTime = rankTime
	for _, rankObj := range rankList {
		rankInfo.FeatherList = append(rankInfo.FeatherList, buildOrder(rankObj))
	}
	return rankInfo
}

func BuildShieldInfoList(rankMap *rankobj.RankMap) (rankInfo *rankpb.AreaRankShield) {
	rankTypeData := rankMap.GetRankTypeData(ranktypes.RankTypeShield)
	rankData, ok := rankTypeData.(*rankobj.ShieldRank)
	if !ok {
		return nil
	}
	rankList, rankTime := rankData.GetListAndTime()
	rankInfo = &rankpb.AreaRankShield{}
	rankInfo.RankTime = rankTime
	for _, rankObj := range rankList {
		rankInfo.ShieldList = append(rankInfo.ShieldList, buildOrder(rankObj))
	}
	return rankInfo
}

func BuildAnQiInfoList(rankMap *rankobj.RankMap) (rankInfo *rankpb.AreaRankAnQi) {
	rankTypeData := rankMap.GetRankTypeData(ranktypes.RankTypeAnQi)
	rankData, ok := rankTypeData.(*rankobj.AnQiRank)
	if !ok {
		return nil
	}
	rankList, rankTime := rankData.GetListAndTime()
	rankInfo = &rankpb.AreaRankAnQi{}
	rankInfo.RankTime = rankTime
	for _, rankObj := range rankList {
		rankInfo.AnQiList = append(rankInfo.AnQiList, buildOrder(rankObj))
	}
	return rankInfo
}
