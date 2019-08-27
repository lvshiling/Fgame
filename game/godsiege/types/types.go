package types

import (
	activitytypes "fgame/fgame/game/activity/types"
)

type GodSiegeBornType int32

const (
	//boss
	GodSiegeBornTypeBoss GodSiegeBornType = iota
	//玩家
	GodSiegeBornTypePlayer
)

func (t GodSiegeBornType) Valid() bool {
	switch t {
	case GodSiegeBornTypeBoss,
		GodSiegeBornTypePlayer:
		return true
	}
	return false
}

type GodSiegeType int32

const (
	//无类型
	GodSiegeTypeNo GodSiegeType = iota
	//麒麟来袭
	GodSiegeTypeQiLin
	//火凤来袭
	GodSiegeTypeHuoFeng
	//毒龙来袭
	GodSiegeTypeDuLong
	//金银密窟
	GodSiegeTypeDenseWat
	//麒麟来袭(本服)
	GodSiegeTypeLocalQiLin
)

func (t GodSiegeType) Valid() bool {
	switch t {
	case GodSiegeTypeQiLin,
		GodSiegeTypeHuoFeng,
		GodSiegeTypeDuLong,
		GodSiegeTypeDenseWat,
		GodSiegeTypeLocalQiLin:
		return true
	}
	return false
}

// 四神类型To活动类型
var godSiegeTypeMap = map[GodSiegeType]activitytypes.ActivityType{
	GodSiegeTypeQiLin:      activitytypes.ActivityTypeGodSiegeQiLin,
	GodSiegeTypeHuoFeng:    activitytypes.ActivityTypeGodSiegeHuoFeng,
	GodSiegeTypeDuLong:     activitytypes.ActivityTypeGodSiegeDuLong,
	GodSiegeTypeDenseWat:   activitytypes.ActivityTypeDenseWat,
	GodSiegeTypeLocalQiLin: activitytypes.ActivityTypeLocalGodSiegeQiLin,
}

func (t GodSiegeType) GetActivityType() (activitytypes.ActivityType, bool) {
	acType, isExist := godSiegeTypeMap[t]
	return acType, isExist
}

// 活动类型To四神类型
var activityTypeToGodSiegeTypeMap = map[activitytypes.ActivityType]GodSiegeType{
	activitytypes.ActivityTypeGodSiegeQiLin:      GodSiegeTypeQiLin,
	activitytypes.ActivityTypeGodSiegeHuoFeng:    GodSiegeTypeHuoFeng,
	activitytypes.ActivityTypeGodSiegeDuLong:     GodSiegeTypeDuLong,
	activitytypes.ActivityTypeDenseWat:           GodSiegeTypeDenseWat,
	activitytypes.ActivityTypeLocalGodSiegeQiLin: GodSiegeTypeLocalQiLin,
}

func GetGodSiegeType(t activitytypes.ActivityType) (GodSiegeType, bool) {
	siegeType, isExist := activityTypeToGodSiegeTypeMap[t]
	return siegeType, isExist
}

type GodSiegePosType int32

const (
	//出生点1
	GodSiegePosTypeOne GodSiegePosType = 1 + iota
	//出生点2
	GodSiegePosTypeTwo
	//出生点3
	GodSiegePosTypeThree
	//出生点4
	GodSiegePosTypeFour
)

const (
	GodSiegePosTypeMin = GodSiegePosTypeOne
	GodSiegePosTypeMax = GodSiegePosTypeOne
)

func (t GodSiegePosType) Valid() bool {
	switch t {
	case GodSiegePosTypeOne,
		GodSiegePosTypeTwo,
		GodSiegePosTypeThree,
		GodSiegePosTypeFour:
		return true
	}
	return false
}

type GodSiegeBossStatusType int32

const (
	//待刷新
	GodSiegeBossStatusTypeInit GodSiegeBossStatusType = iota
	//已刷新
	GodSiegeBossStatusTypeLive
	//已死亡
	GodSiegeBossStatusTypeDead
)

//四神服务类型
type GodSiegeServerType int32

const (
	GodSiegeServerTypeLocal GodSiegeServerType = iota //本服
	GodSiegeServerTypeCross                           //跨服
)

var (
	godSiegeServerMap = map[GodSiegeServerType][]GodSiegeType{
		GodSiegeServerTypeCross: {
			GodSiegeTypeQiLin,
			GodSiegeTypeHuoFeng,
			GodSiegeTypeDuLong,
			GodSiegeTypeDenseWat,
		},
		GodSiegeServerTypeLocal: {
			GodSiegeTypeLocalQiLin,
		},
	}
)

// 四神类型列表
func (t GodSiegeServerType) GetGodSiegeTypeList() []GodSiegeType {
	return godSiegeServerMap[t]
}
