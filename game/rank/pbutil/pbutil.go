package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	rankentity "fgame/fgame/game/rank/entity"
)

func BuildSCRankMyGet(isArea bool, typ int32, pos int32) *uipb.SCRankMyGet {
	rankMyGet := &uipb.SCRankMyGet{}
	rankMyGet.RankType = &typ
	rankMyGet.Pos = &pos
	rankMyGet.IsArea = &isArea
	return rankMyGet
}

func BuildSCRankForceGet(showServer bool, isArea bool, page int32, forceList []*rankentity.PlayerForceData, rankTime int64) *uipb.SCRankForceGet {
	rankForceGet := &uipb.SCRankForceGet{}
	rankForceGet.Page = &page
	showServer = showServer || isArea
	for _, forceObj := range forceList {
		rankForceGet.ForceList = append(rankForceGet.ForceList, buildForce(showServer, forceObj))
	}
	rankForceGet.RankTime = &rankTime
	rankForceGet.IsArea = &isArea
	return rankForceGet
}

func BuildSCRankMountGet(isArea bool, page int32, mountList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankMountGet {
	rankMountGet := &uipb.SCRankMountGet{}
	rankMountGet.Page = &page
	for _, mountObj := range mountList {
		rankMountGet.MountList = append(rankMountGet.MountList, buildOrder(isArea, mountObj))
	}
	rankMountGet.RankTime = &rankTime
	rankMountGet.IsArea = &isArea
	return rankMountGet
}

func BuildSCRankWingGet(isArea bool, page int32, wingList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankWingGet {
	rankWingGet := &uipb.SCRankWingGet{}
	rankWingGet.Page = &page
	for _, wingObj := range wingList {
		rankWingGet.WingList = append(rankWingGet.WingList, buildOrder(isArea, wingObj))
	}
	rankWingGet.RankTime = &rankTime
	rankWingGet.IsArea = &isArea
	return rankWingGet
}

func BuildSCRankBodyShieldGet(isArea bool, page int32, bodyShieldList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankBodyShieldGet {
	rankBodyShieldGet := &uipb.SCRankBodyShieldGet{}
	rankBodyShieldGet.Page = &page
	for _, bodyShieldObj := range bodyShieldList {
		rankBodyShieldGet.BodyShieldList = append(rankBodyShieldGet.BodyShieldList, buildOrder(isArea, bodyShieldObj))
	}
	rankBodyShieldGet.RankTime = &rankTime
	rankBodyShieldGet.IsArea = &isArea
	return rankBodyShieldGet
}

func BuildSCRankWeaponGet(isArea bool, page int32, weaponList []*rankentity.PlayerWeaponData, rankTime int64) *uipb.SCRankWeaponGet {
	rankWeaponGet := &uipb.SCRankWeaponGet{}
	rankWeaponGet.Page = &page
	for _, weaponObj := range weaponList {
		rankWeaponGet.WeaponList = append(rankWeaponGet.WeaponList, buildWeapon(isArea, weaponObj))
	}
	rankWeaponGet.RankTime = &rankTime
	rankWeaponGet.IsArea = &isArea
	return rankWeaponGet
}

func BuildSCRankGangGet(isArea bool, page int32, gangList []*rankentity.PlayerGangData, rankTime int64) *uipb.SCRankGangGet {
	rankGangGet := &uipb.SCRankGangGet{}
	rankGangGet.Page = &page
	for _, gangObj := range gangList {
		rankGangGet.GangList = append(rankGangGet.GangList, buildGang(isArea, gangObj))
	}
	rankGangGet.RankTime = &rankTime
	rankGangGet.IsArea = &isArea
	return rankGangGet
}

func BuildSCRankShenFaGet(isArea bool, page int32, shenFaList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankShenFaGet {
	rankShenFaGet := &uipb.SCRankShenFaGet{}
	rankShenFaGet.Page = &page
	for _, shenFaObj := range shenFaList {
		rankShenFaGet.ShenFaList = append(rankShenFaGet.ShenFaList, buildOrder(isArea, shenFaObj))
	}
	rankShenFaGet.RankTime = &rankTime
	rankShenFaGet.IsArea = &isArea
	return rankShenFaGet
}

func BuildSCRankLingYuGet(isArea bool, page int32, lingYuList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankLingYuGet {
	rankLingYuGet := &uipb.SCRankLingYuGet{}
	rankLingYuGet.Page = &page
	for _, lingYuObj := range lingYuList {
		rankLingYuGet.LingYuList = append(rankLingYuGet.LingYuList, buildOrder(isArea, lingYuObj))
	}
	rankLingYuGet.RankTime = &rankTime
	rankLingYuGet.IsArea = &isArea
	return rankLingYuGet
}

func BuildSCRankFeatherGet(isArea bool, page int32, featherList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankFeatherGet {
	rankFeatherGet := &uipb.SCRankFeatherGet{}
	rankFeatherGet.Page = &page
	for _, featherObj := range featherList {
		rankFeatherGet.FeahterList = append(rankFeatherGet.FeahterList, buildOrder(isArea, featherObj))
	}
	rankFeatherGet.RankTime = &rankTime
	rankFeatherGet.IsArea = &isArea
	return rankFeatherGet
}

func BuildSCRankShieldGet(isArea bool, page int32, shieldList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankShieldGet {
	rankShieldGet := &uipb.SCRankShieldGet{}
	rankShieldGet.Page = &page
	for _, shieldObj := range shieldList {
		rankShieldGet.ShieldList = append(rankShieldGet.ShieldList, buildOrder(isArea, shieldObj))
	}
	rankShieldGet.RankTime = &rankTime
	rankShieldGet.IsArea = &isArea
	return rankShieldGet
}

func BuildSCRankAnQiGet(isArea bool, page int32, anQiList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankAnQiGet {
	rankAnQiGet := &uipb.SCRankAnQiGet{}
	rankAnQiGet.Page = &page
	for _, anQiObj := range anQiList {
		rankAnQiGet.AnQiList = append(rankAnQiGet.AnQiList, buildOrder(isArea, anQiObj))
	}
	rankAnQiGet.RankTime = &rankTime
	rankAnQiGet.IsArea = &isArea
	return rankAnQiGet
}

func BuildSCRankFaBaoGet(isArea bool, page int32, faBaoList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankFaBaoGet {
	rankFaBaoGet := &uipb.SCRankFaBaoGet{}
	rankFaBaoGet.Page = &page
	for _, faBaoObj := range faBaoList {
		rankFaBaoGet.FaBaoList = append(rankFaBaoGet.FaBaoList, buildOrder(isArea, faBaoObj))
	}
	rankFaBaoGet.RankTime = &rankTime
	rankFaBaoGet.IsArea = &isArea
	return rankFaBaoGet
}

func BuildSCRankXianTiGet(isArea bool, page int32, xianTiList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankXianTiGet {
	rankXianTiGet := &uipb.SCRankXianTiGet{}
	rankXianTiGet.Page = &page
	for _, xianTiObj := range xianTiList {
		rankXianTiGet.XianTiList = append(rankXianTiGet.XianTiList, buildOrder(isArea, xianTiObj))
	}
	rankXianTiGet.RankTime = &rankTime
	rankXianTiGet.IsArea = &isArea
	return rankXianTiGet
}

func BuildSCRankShiHunFanGet(isArea bool, page int32, xianTiList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankShiHunFanGet {
	scMsg := &uipb.SCRankShiHunFanGet{}
	scMsg.Page = &page
	for _, xianTiObj := range xianTiList {
		scMsg.ShiHunFanList = append(scMsg.ShiHunFanList, buildOrder(isArea, xianTiObj))
	}
	scMsg.RankTime = &rankTime
	scMsg.IsArea = &isArea
	return scMsg
}

func BuildSCRankTianMoTiGet(isArea bool, page int32, xianTiList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankTianMoTiGet {
	scMsg := &uipb.SCRankTianMoTiGet{}
	scMsg.Page = &page
	for _, xianTiObj := range xianTiList {
		scMsg.TianMoTiList = append(scMsg.TianMoTiList, buildOrder(isArea, xianTiObj))
	}
	scMsg.RankTime = &rankTime
	scMsg.IsArea = &isArea
	return scMsg
}

func BuildSCRankLingTongDevGet(classType int32, isArea bool, page int32, lingTongDevList []*rankentity.PlayerOrderData, rankTime int64) *uipb.SCRankLingTongDevGet {
	rankLingTongDevGet := &uipb.SCRankLingTongDevGet{}
	rankLingTongDevGet.ClassType = &classType
	rankLingTongDevGet.Page = &page
	for _, lingTongDevObj := range lingTongDevList {
		rankLingTongDevGet.LingTongDevList = append(rankLingTongDevGet.LingTongDevList, buildOrder(isArea, lingTongDevObj))
	}
	rankLingTongDevGet.RankTime = &rankTime
	rankLingTongDevGet.IsArea = &isArea
	return rankLingTongDevGet
}

func BuildSCRankLingTongLevelGet(isArea bool, page int32, lingTongLevelList []*rankentity.PlayerPropertyData, rankTime int64) *uipb.SCRankLingTongLevelGet {
	rankLingTongLevelGet := &uipb.SCRankLingTongLevelGet{}
	rankLingTongLevelGet.Page = &page
	for _, lingTongLevelObj := range lingTongLevelList {
		rankLingTongLevelGet.LingtongLevelList = append(rankLingTongLevelGet.LingtongLevelList, buildProperty(isArea, lingTongLevelObj))
	}
	rankLingTongLevelGet.RankTime = &rankTime
	rankLingTongLevelGet.IsArea = &isArea
	return rankLingTongLevelGet
}

func BuildSCRankFeiShengGet(isArea bool, page int32, feiShengList []*rankentity.PlayerPropertyData, rankTime int64) *uipb.SCRankFeiShengGet {
	scMsg := &uipb.SCRankFeiShengGet{}
	scMsg.Page = &page
	for _, rankObj := range feiShengList {
		scMsg.FeiShengList = append(scMsg.FeiShengList, buildProperty(isArea, rankObj))
	}
	scMsg.RankTime = &rankTime
	scMsg.IsArea = &isArea
	return scMsg
}

func buildProperty(isArea bool, propotyData *rankentity.PlayerPropertyData) *uipb.RankProperty {
	rankProperty := &uipb.RankProperty{}
	id := propotyData.PlayerId
	name := propotyData.PlayerName
	num := propotyData.Num
	power := propotyData.Power
	if isArea {
		serverId := propotyData.ServerId
		rankProperty.ServerId = &serverId
	}
	rankProperty.PlayerId = &id
	rankProperty.PlayerName = &name
	rankProperty.Num = &num
	rankProperty.Power = &power

	return rankProperty
}

func buildForce(isArea bool, forceData *rankentity.PlayerForceData) *uipb.RankForce {
	rankForce := &uipb.RankForce{}
	id := forceData.PlayerId
	name := forceData.PlayerName
	force := forceData.Force
	role := forceData.Role
	sex := forceData.Sex
	gangName := forceData.GangName
	if isArea {
		serverId := forceData.ServerId
		rankForce.ServerId = &serverId
	}
	rankForce.PlayerId = &id
	rankForce.PlayerName = &name
	rankForce.GangName = &gangName
	rankForce.Power = &force
	rankForce.Role = &role
	rankForce.Sex = &sex
	return rankForce
}

func buildOrder(isArea bool, orderData *rankentity.PlayerOrderData) *uipb.RankOrder {
	rankOder := &uipb.RankOrder{}
	id := orderData.PlayerId
	name := orderData.PlayerName
	order := orderData.Order
	power := orderData.Power
	if isArea {
		serverId := orderData.ServerId
		rankOder.ServerId = &serverId
	}
	rankOder.PlayerId = &id
	rankOder.PlayerName = &name
	rankOder.Order = &order
	rankOder.Power = &power
	return rankOder
}

func buildWeapon(isArea bool, weaponData *rankentity.PlayerWeaponData) *uipb.RankWeapon {
	rankWeapon := &uipb.RankWeapon{}
	id := weaponData.PlayerId
	name := weaponData.PlayerName
	star := weaponData.Star
	wear := weaponData.WearId
	power := weaponData.Power
	role := weaponData.Role
	sex := weaponData.Sex

	if isArea {
		serverId := weaponData.ServerId
		rankWeapon.ServerId = &serverId
	}

	rankWeapon.PlayerId = &id
	rankWeapon.PlayerName = &name
	rankWeapon.Star = &star
	rankWeapon.WearId = &wear
	rankWeapon.Power = &power
	rankWeapon.Role = &role
	rankWeapon.Sex = &sex
	return rankWeapon
}

func buildGang(isArea bool, gangData *rankentity.PlayerGangData) *uipb.RankGang {
	rankGang := &uipb.RankGang{}
	name := gangData.GangName
	leadName := gangData.LeadName
	leadId := gangData.LeadId
	power := gangData.Power
	role := gangData.Role
	sex := gangData.Sex

	if isArea {
		serverId := gangData.ServerId
		rankGang.ServerId = &serverId
	}

	rankGang.GangName = &name
	rankGang.LeaderName = &leadName
	rankGang.LeaderId = &leadId
	rankGang.Power = &power
	rankGang.Role = &role
	rankGang.Sex = &sex
	return rankGang
}
