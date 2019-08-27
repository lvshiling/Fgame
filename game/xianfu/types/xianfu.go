package types

import (
	constanttypes "fgame/fgame/game/constant/types"
)

type XianfuType int32

const (
	//银两副本
	XianfuTypeSilver XianfuType = iota
	//经验副本
	XianfuTypeExp
	//材料副本
	XianfuTypeItem
	//元宝副本
	XianfuTypeGold
	//双修副本(major模块处理；前端显示，不要删)
	XianfuTypeCouples
)

var xianfuTypeMap = map[XianfuType]string{
	XianfuTypeSilver:  "银两",
	XianfuTypeExp:     "经验",
	XianfuTypeItem:    "材料",
	XianfuTypeGold:    "元宝",
	XianfuTypeCouples: "双修",
}

func (t XianfuType) Valid() bool {
	switch t {
	case XianfuTypeSilver,
		XianfuTypeExp,
		XianfuTypeItem,
		XianfuTypeGold,
		XianfuTypeCouples:
		return true
	default:
		return false
	}
}

const (
	MinXianfuType = XianfuTypeSilver
	MaxXianfuType = XianfuTypeCouples
)

func (t XianfuType) String() string {
	return xianfuTypeMap[t]
}

func (t XianfuType) GetChallengeNumConstantType() constanttypes.ConstantType {
	switch t {
	case XianfuTypeSilver:
		return constanttypes.ConstantTypeXianfuSilverChallengeNum
	case XianfuTypeExp:
		return constanttypes.ConstantTypeXianfuExpChallengeNum
	}
	return 0
}

func (t XianfuType) GetFinishAllConstantType() constanttypes.ConstantType {
	switch t {
	case XianfuTypeSilver:
		return constanttypes.ConstantTypeXianfuSilverFinishCostGold
	case XianfuTypeExp:
		return constanttypes.ConstantTypeXianfuExpFinishCostGold
	}
	return 0
}

func (t XianfuType) GetSaoDangNeedLevelConstantType() constanttypes.ConstantType {
	switch t {
	case XianfuTypeSilver:
		return constanttypes.ConstantTypeXianFuSilverSaoDangNeedLevel
	case XianfuTypeExp:
		return constanttypes.ConstantTypeXianFuExpSaoDangNeedLevel
	}
	return 0
}

type XianfuState int32

const (
	XianfuStateWaitedToUpgrade XianfuState = iota //待升级
	XianfuStateUpgrading                          //升级中
)
