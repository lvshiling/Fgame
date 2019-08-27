package types

import (
	centertypes "fgame/fgame/center/types"
	activitytypes "fgame/fgame/game/activity/types"
)

type CrossType int32

const (
	CrossTypeNone CrossType = iota
	//3v3竞技场
	CrossTypeArena
	//世界boss
	CrossTypeWorldboss
	//屠龙
	CrossTypeTuLong
	//血矿
	CrossTypeXueKuang
	//无间炼狱--------5
	CrossTypeLianYu
	//神兽攻城--麒麟
	CrossTypeGodSiegeQiLin
	//神兽攻城--火凤
	CrossTypeGodSiegeHuoFeng
	//神兽攻城--毒龙
	CrossTypeGodSiegeDuLong
	//组队副本
	CrossTypeTeamCopy
	//金银密窟-----------10
	CrossTypeDenseWat
	//神魔战场
	CrossTypeShenMoWar
	//运营活动副本
	CrossTypeWelfare
	//跨服之战
	CrossTypeArenapvp
	//创世之战
	CrossTypeChuangShi
	//圣兽boss
	CrossTypeShengShou
	//珍惜boss
	CrossTypeZhenXi
)

func (t CrossType) Valid() bool {
	switch t {
	case CrossTypeArena,
		CrossTypeWorldboss,
		CrossTypeTuLong,
		CrossTypeXueKuang,
		CrossTypeLianYu,
		CrossTypeGodSiegeQiLin,
		CrossTypeGodSiegeHuoFeng,
		CrossTypeGodSiegeDuLong,
		CrossTypeTeamCopy,
		CrossTypeDenseWat,
		CrossTypeShenMoWar,
		CrossTypeWelfare,
		CrossTypeArenapvp,
		CrossTypeChuangShi,
		CrossTypeShengShou,
		CrossTypeZhenXi:
		return true
	default:
		return false
	}
}

var (
	serverTypeMap = map[CrossType]centertypes.GameServerType{
		CrossTypeNone:            centertypes.GameServerTypeSingle,
		CrossTypeArena:           centertypes.GameServerTypeRegion,
		CrossTypeWorldboss:       centertypes.GameServerTypeRegion,
		CrossTypeTuLong:          centertypes.GameServerTypeRegion,
		CrossTypeXueKuang:        centertypes.GameServerTypeRegion,
		CrossTypeLianYu:          centertypes.GameServerTypeRegion,
		CrossTypeGodSiegeQiLin:   centertypes.GameServerTypeRegion,
		CrossTypeGodSiegeHuoFeng: centertypes.GameServerTypeRegion,
		CrossTypeGodSiegeDuLong:  centertypes.GameServerTypeRegion,
		CrossTypeTeamCopy:        centertypes.GameServerTypeGroup,
		CrossTypeDenseWat:        centertypes.GameServerTypeRegion,
		CrossTypeShenMoWar:       centertypes.GameServerTypePlatform,
		CrossTypeWelfare:         centertypes.GameServerTypePlatform,
		CrossTypeArenapvp:        centertypes.GameServerTypeAll,
		CrossTypeChuangShi:       centertypes.GameServerTypeAll,
		CrossTypeShengShou:       centertypes.GameServerTypePlatform,
		CrossTypeZhenXi:          centertypes.GameServerTypeRegion,
	}
)

func (t CrossType) GetServerType() centertypes.GameServerType {
	return serverTypeMap[t]
}

var (
	crossTypeMap = map[CrossType]string{
		CrossTypeNone: "没有",
		//3v3竞技场
		CrossTypeArena: "3v3竞技场",
		//世界boss
		CrossTypeWorldboss: "世界boss",
		//屠龙
		CrossTypeTuLong: "仙盟屠龙",
		//血矿
		CrossTypeXueKuang: "血矿",
		//无间炼狱
		CrossTypeLianYu: "无间炼狱",
		//神兽攻城-麒麟
		CrossTypeGodSiegeQiLin: "麒麟来袭",
		//神兽攻城-火凤
		CrossTypeGodSiegeHuoFeng: "火凤来袭",
		//神兽攻城-毒龙
		CrossTypeGodSiegeDuLong: "毒龙来袭",
		//组队副本
		CrossTypeTeamCopy: "组队副本",
		//金银密窟
		CrossTypeDenseWat: "金银密窟",
		//神魔战场
		CrossTypeShenMoWar: "神魔战场",
		//奇遇岛
		CrossTypeWelfare: "运营活动副本",
		//跨服之战
		CrossTypeArenapvp: "跨服之战",
		//创世之战
		CrossTypeChuangShi: "创世之战",
		CrossTypeShengShou: "圣兽秘境",
		CrossTypeZhenXi:    "珍稀boss",
	}
)

func (t CrossType) String() string {
	return crossTypeMap[t]
}

func (t CrossType) CrossTypeToActivityType() (activitytypes.ActivityType, bool) {
	acType, ok := crossSceneActivityMap[t]
	return acType, ok
}

var (
	crossSceneActivityMap = map[CrossType]activitytypes.ActivityType{
		CrossTypeArena:           activitytypes.ActivityTypeArena,
		CrossTypeGodSiegeQiLin:   activitytypes.ActivityTypeGodSiegeQiLin,
		CrossTypeGodSiegeHuoFeng: activitytypes.ActivityTypeGodSiegeHuoFeng,
		CrossTypeGodSiegeDuLong:  activitytypes.ActivityTypeGodSiegeDuLong,
		CrossTypeTuLong:          activitytypes.ActivityTypeCoressTuLong,
		CrossTypeLianYu:          activitytypes.ActivityTypeLianYu,
		CrossTypeDenseWat:        activitytypes.ActivityTypeDenseWat,
		CrossTypeShenMoWar:       activitytypes.ActivityTypeShenMoWar,
		CrossTypeArenapvp:        activitytypes.ActivityTypeArenapvp,
	}
)

type CrossBehaviorType int32

const (
	CrossBehaviorTypeNormal CrossBehaviorType = iota
	CrossBehaviorTypeTrack
)
