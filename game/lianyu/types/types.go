package types

import (
	activitytypes "fgame/fgame/game/activity/types"
)

type LianYuBornType int32

const (
	//本服boss
	LianYuBornTypeBoss LianYuBornType = iota
	//本服玩家
	LianYuBornTypePlayer
	//跨服boss
	LianYuBornTypeCrossBoss
	//跨服玩家
	LianYuBornTypeCrossPlayer
)

func (t LianYuBornType) Valid() bool {
	switch t {
	case LianYuBornTypeBoss,
		LianYuBornTypePlayer,
		LianYuBornTypeCrossBoss,
		LianYuBornTypeCrossPlayer:
		return true
	}
	return false
}

var (
	acTypeToBornTypeMap = map[activitytypes.ActivityType]LianYuBornType{
		activitytypes.ActivityTypeLianYu:      LianYuBornTypeCrossBoss,
		activitytypes.ActivityTypeLocalLianYu: LianYuBornTypeBoss,
	}
)

func ActivityTypeToBornType(acType activitytypes.ActivityType) (LianYuBornType, bool) {
	bornType, ok := acTypeToBornTypeMap[acType]
	return bornType, ok
}

const (
	//杀气排名
	ShaQiRankSize = 10
)

type LianYuPosType int32

const (
	//出生点1
	LianYuPosTypeOne LianYuPosType = 1 + iota
	//出生点2
	LianYuPosTypeTwo
	//出生点3
	LianYuPosTypeThree
	//出生点4
	LianYuPosTypeFour
)

const (
	LianYuPosTypeMin = LianYuPosTypeOne
	LianYuPosTypeMax = LianYuPosTypeOne
)

func (t LianYuPosType) Valid() bool {
	switch t {
	case LianYuPosTypeOne,
		LianYuPosTypeTwo,
		LianYuPosTypeThree,
		LianYuPosTypeFour:
		return true
	}
	return false
}

type LianYuBossStatusType int32

const (
	//待刷新
	LianYuBossStatusTypeInit LianYuBossStatusType = iota
	//已刷新
	LianYuBossStatusTypeLive
	//已死亡
	LianYuBossStatusTypeDead
)
