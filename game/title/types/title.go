package types

import (
	ranktypes "fgame/fgame/game/rank/types"
)

type TitleType int32

const (
	//排行榜称号
	TitleTypeRank TitleType = 1 + iota
	//活动称号 (必现限时)
	TitleTypeActivity
	//情侣称号
	TitleTypeCouples
	// 普通称号
	TitleTypeNormal
	// 自定义称号
	TitleTypeDefine
	// NPC称号
	TitleTypeNPC
	//采集物称号
	TitleTypeColThings
	//大皇帝
	TitleTypeKing
	//神魔战场称号
	TitleTypeShenMo
	//特殊称号
	TitleTypeSpecial
)

func (tt TitleType) Valid() bool {
	switch tt {
	case TitleTypeRank,
		TitleTypeActivity,
		TitleTypeCouples,
		TitleTypeNormal,
		TitleTypeDefine,
		TitleTypeNPC,
		TitleTypeColThings,
		TitleTypeKing,
		TitleTypeShenMo,
		TitleTypeSpecial:
		break
	default:
		return false
	}
	return true
}

type TitleSubType interface {
	SubType() int32
	Valid() bool
}

type TitleSubTypeFactory interface {
	CreateTitleSubType(subType int32) TitleSubType
}

type TitleSubTypeFactoryFunc func(subType int32) TitleSubType

func (tstff TitleSubTypeFactoryFunc) CreateTitleSubType(subType int32) TitleSubType {
	return tstff(subType)
}

//无子类型
type TitleCommonSubType int32

const (
	//默认子类型
	TitleCommonSubTypeDefault TitleCommonSubType = iota
)

func (tcst TitleCommonSubType) SubType() int32 {
	return int32(tcst)
}

func (tcst TitleCommonSubType) Valid() bool {
	switch tcst {
	case
		TitleCommonSubTypeDefault:
		return true
	}
	return false
}
func CreateTitleCommonSubType(subType int32) TitleSubType {
	return TitleCommonSubType(subType)
}

//排行榜子类型
type TitleRankSubType int32

const (
	//战力排行榜第一称号
	TitleRankSubTypeForce TitleRankSubType = 1 + iota
	//兵魂排行榜第一称号
	TitleRankSubTypeWeapon
	//天劫塔排行榜第一称号
	//TitleRankSubTypeRealm
)

func (tstr TitleRankSubType) SubType() int32 {
	return int32(tstr)
}

func (tstr TitleRankSubType) Valid() bool {
	switch tstr {
	case TitleRankSubTypeForce,
		TitleRankSubTypeWeapon:
		//TitleRankSubTypeRealm:
		return true
	}
	return false
}

func CreateTitleRankSubType(subType int32) TitleSubType {
	return TitleRankSubType(subType)
}

var titleRankSubTypeMap = map[TitleRankSubType]string{
	TitleRankSubTypeForce:  "战力排行榜第一称号",
	TitleRankSubTypeWeapon: "兵魂排行榜第一称号",
	//TitleRankSubTypeRealm:  "天劫塔排行榜第一称号",
}

func (tstr TitleRankSubType) String() string {
	return titleRankSubTypeMap[tstr]
}

var titleRankSubMap = map[TitleRankSubType]ranktypes.RankType{
	TitleRankSubTypeForce:  ranktypes.RankTypeForce,
	TitleRankSubTypeWeapon: ranktypes.RankTypeWeapon,
}

func GetTitleRankSubMap() map[TitleRankSubType]ranktypes.RankType {
	return titleRankSubMap
}

func (tstr TitleRankSubType) TitleRankSubTypeToRankType() ranktypes.RankType {
	return titleRankSubMap[tstr]
}

var rankTitleSubMap = map[ranktypes.RankType]TitleRankSubType{
	ranktypes.RankTypeForce:  TitleRankSubTypeForce,
	ranktypes.RankTypeWeapon: TitleRankSubTypeWeapon,
}

func RankTypeToTitleRankSubType(rankType ranktypes.RankType) TitleRankSubType {
	return rankTitleSubMap[rankType]
}

var (
	titleSubTypeFactoryMap = make(map[TitleType]TitleSubTypeFactory)
)

func CreateTitleSubType(typ TitleType, subType int32) TitleSubType {
	factory, ok := titleSubTypeFactoryMap[typ]
	if !ok {
		return nil
	}
	return factory.CreateTitleSubType(subType)
}

func init() {
	titleSubTypeFactoryMap[TitleTypeRank] = TitleSubTypeFactoryFunc(CreateTitleRankSubType)
	titleSubTypeFactoryMap[TitleTypeActivity] = TitleSubTypeFactoryFunc(CreateTitleCommonSubType)
	titleSubTypeFactoryMap[TitleTypeCouples] = TitleSubTypeFactoryFunc(CreateTitleCommonSubType)
	titleSubTypeFactoryMap[TitleTypeNormal] = TitleSubTypeFactoryFunc(CreateTitleCommonSubType)
	titleSubTypeFactoryMap[TitleTypeDefine] = TitleSubTypeFactoryFunc(CreateTitleCommonSubType)
	titleSubTypeFactoryMap[TitleTypeNPC] = TitleSubTypeFactoryFunc(CreateTitleCommonSubType)
	titleSubTypeFactoryMap[TitleTypeColThings] = TitleSubTypeFactoryFunc(CreateTitleCommonSubType)
	titleSubTypeFactoryMap[TitleTypeKing] = TitleSubTypeFactoryFunc(CreateTitleCommonSubType)
	titleSubTypeFactoryMap[TitleTypeShenMo] = TitleSubTypeFactoryFunc(CreateTitleCommonSubType)
	titleSubTypeFactoryMap[TitleTypeSpecial] = TitleSubTypeFactoryFunc(CreateTitleCommonSubType)

}
