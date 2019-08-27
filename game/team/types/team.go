package types

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
)

type TeamPurposeType int32

const (
	//通常组队
	TeamPurposeTypeNormal TeamPurposeType = iota
	//组队副本.银两
	TeamPurposeTypeFuBenSilver
	//组队副本.星尘
	TeamPurposeTypeFuBenXingChen
	//组队副本.血魔
	TeamPurposeTypeFuBenXueMo
	//组队副本.转生装备
	TeamPurposeTypeFuBenZhuangShengEquip
	//组队副本.兵魂
	TeamPurposeTypeFuBenWeapon
	//组队副本.强化
	TeamPurposeTypeFuBenStrength
	//组队副本.灵童
	TeamPurposeTypeFuBenLingTong
	//组队副本.升星
	TeamPurposeTypeFuBenUpstar
	//3v3
	TeamPurposeTypeArena
)

func (t TeamPurposeType) Vaild() bool {
	switch t {
	case TeamPurposeTypeFuBenSilver,
		TeamPurposeTypeFuBenXingChen,
		TeamPurposeTypeFuBenXueMo,
		TeamPurposeTypeFuBenZhuangShengEquip,
		TeamPurposeTypeFuBenWeapon,
		TeamPurposeTypeFuBenStrength,
		TeamPurposeTypeFuBenLingTong,
		TeamPurposeTypeFuBenUpstar,
		TeamPurposeTypeArena:
		return true
	}
	return false
}

func (t TeamPurposeType) GetFuncOpenType() funcopentypes.FuncOpenType {
	return teamPurposeMap[t]
}

var teamPurposeMap = map[TeamPurposeType]funcopentypes.FuncOpenType{
	TeamPurposeTypeFuBenSilver:           funcopentypes.FuncOpenTypeTeamCopySilver,
	TeamPurposeTypeFuBenXingChen:         funcopentypes.FuncOpenTypeTeamCopyXingChen,
	TeamPurposeTypeFuBenXueMo:            funcopentypes.FuncOpenTypeTeamCopyXueMo,
	TeamPurposeTypeFuBenZhuangShengEquip: funcopentypes.FuncOpenTypeTeamCopyZhuangShengEquip,
	TeamPurposeTypeFuBenWeapon:           funcopentypes.FuncOpenTypeTeamCopyWeapon,
	TeamPurposeTypeFuBenStrength:         funcopentypes.FuncOpenTypeTeamCopyStrength,
	TeamPurposeTypeFuBenLingTong:         funcopentypes.FuncOpenTypeTeamCopyLingTong,
	TeamPurposeTypeFuBenUpstar:           funcopentypes.FuncOpenTypeTeamCopyUpstar,
	TeamPurposeTypeArena:                 funcopentypes.FuncOpenTypeArena,
}
