package types

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerlingtongtypes "fgame/fgame/game/lingtong/player/types"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	viptypes "fgame/fgame/game/vip/types"
)

type LingTongDevSysType int32

const (
	LingTongDevSysTypeDefault LingTongDevSysType = iota
	//灵兵
	LingTongDevSysTypeLingBing
	//灵骑
	LingTongDevSysTypeLingQi
	//灵翼
	LingTongDevSysTypeLingYi
	//灵身
	LingTongDevSysTypeLingShen
	//灵域
	LingTongDevSysTypeLingYu
	//灵宝
	LingTongDevSysTypeLingBao
	//灵体
	LingTongDevSysTypeLingTi
)

func (t LingTongDevSysType) Vaild() bool {
	switch t {
	case LingTongDevSysTypeLingBao,
		LingTongDevSysTypeLingBing,
		LingTongDevSysTypeLingQi,
		LingTongDevSysTypeLingShen,
		LingTongDevSysTypeLingTi,
		LingTongDevSysTypeLingYi,
		LingTongDevSysTypeLingYu:
		return true
	}
	return false
}

func (t LingTongDevSysType) GetVipType() viptypes.CostLevelRuleType {
	return lingTongDevSysTypeVipMap[t]
}

func (t LingTongDevSysType) String() string {
	return lingTongDevSysTypeStringMap[t]
}

var lingTongDevSysTypeStringMap = map[LingTongDevSysType]string{
	LingTongDevSysTypeLingBao:  "灵宝",
	LingTongDevSysTypeLingBing: "灵兵",
	LingTongDevSysTypeLingQi:   "灵骑",
	LingTongDevSysTypeLingShen: "灵身",
	LingTongDevSysTypeLingTi:   "灵体",
	LingTongDevSysTypeLingYi:   "灵翼",
	LingTongDevSysTypeLingYu:   "灵域",
}

var lingTongDevSysTypeVipMap = map[LingTongDevSysType]viptypes.CostLevelRuleType{
	LingTongDevSysTypeLingBao:  viptypes.CostLevelRuleTypeLingTongFaBao,
	LingTongDevSysTypeLingBing: viptypes.CostLevelRuleTypeLingTongWeapon,
	LingTongDevSysTypeLingQi:   viptypes.CostLevelRuleTypeLingTongMount,
	LingTongDevSysTypeLingShen: viptypes.CostLevelRuleTypeLingTongShenFa,
	LingTongDevSysTypeLingTi:   viptypes.CostLevelRuleTypeLingTongXianTi,
	LingTongDevSysTypeLingYi:   viptypes.CostLevelRuleTypeLingTongWing,
	LingTongDevSysTypeLingYu:   viptypes.CostLevelRuleTypeLingTongLingYu,
}

var lingTongDevSysTypePropertyEffectMap = map[LingTongDevSysType]playerpropertytypes.PropertyEffectorType{
	LingTongDevSysTypeLingBao:  playerpropertytypes.PlayerPropertyEffectorTypeLingTongFaBao,
	LingTongDevSysTypeLingBing: playerpropertytypes.PlayerPropertyEffectorTypeLingTongWeapon,
	LingTongDevSysTypeLingQi:   playerpropertytypes.PlayerPropertyEffectorTypeLingTongMount,
	LingTongDevSysTypeLingShen: playerpropertytypes.PlayerPropertyEffectorTypeLingTongShenFa,
	LingTongDevSysTypeLingTi:   playerpropertytypes.PlayerPropertyEffectorTypeLingTongXianTi,
	LingTongDevSysTypeLingYi:   playerpropertytypes.PlayerPropertyEffectorTypeLingTongWing,
	LingTongDevSysTypeLingYu:   playerpropertytypes.PlayerPropertyEffectorTypeLingTongLingYu,
}

func (t LingTongDevSysType) GetPlayerPropertyEffectType() playerpropertytypes.PropertyEffectorType {
	return lingTongDevSysTypePropertyEffectMap[t]
}

var lingTongDevSysTypeLingTongPropertyEffectMap = map[LingTongDevSysType]playerlingtongtypes.PropertyEffectorType{
	LingTongDevSysTypeLingBao:  playerlingtongtypes.LingTongPropertyEffectorTypeFaBao,
	LingTongDevSysTypeLingBing: playerlingtongtypes.LingTongPropertyEffectorTypeWeapon,
	LingTongDevSysTypeLingQi:   playerlingtongtypes.LingTongPropertyEffectorTypeMount,
	LingTongDevSysTypeLingShen: playerlingtongtypes.LingTongPropertyEffectorTypeShenFa,
	LingTongDevSysTypeLingTi:   playerlingtongtypes.LingTongPropertyEffectorTypeXianTi,
	LingTongDevSysTypeLingYi:   playerlingtongtypes.LingTongPropertyEffectorTypeWing,
	LingTongDevSysTypeLingYu:   playerlingtongtypes.LingTongPropertyEffectorTypeLingYu,
}

func (t LingTongDevSysType) GetLingTongPropertyEffectType() playerlingtongtypes.PropertyEffectorType {
	return lingTongDevSysTypeLingTongPropertyEffectMap[t]
}

var lingTongDevSysTypeFuncOpenMap = map[LingTongDevSysType]funcopentypes.FuncOpenType{
	LingTongDevSysTypeLingBao:  funcopentypes.FuncOpenTypeLingTongFaBao,
	LingTongDevSysTypeLingBing: funcopentypes.FuncOpenTypeLingTongWeapon,
	LingTongDevSysTypeLingQi:   funcopentypes.FuncOpenTypeLingTongMount,
	LingTongDevSysTypeLingShen: funcopentypes.FuncOpenTypeLingTongShenFa,
	LingTongDevSysTypeLingTi:   funcopentypes.FuncOpenTypeLingTongXianTi,
	LingTongDevSysTypeLingYi:   funcopentypes.FuncOpenTypeLingTongWing,
	LingTongDevSysTypeLingYu:   funcopentypes.FuncOpenTypeLingTongLingYu,
}

func (t LingTongDevSysType) GetFuncOpenType() funcopentypes.FuncOpenType {
	return lingTongDevSysTypeFuncOpenMap[t]
}

var lingTongDevSysTypeAdvanceFuncOpenMap = map[LingTongDevSysType]funcopentypes.FuncOpenType{
	LingTongDevSysTypeLingBao:  funcopentypes.FuncOpenTypeLingTongFaBaoAdvanced,
	LingTongDevSysTypeLingBing: funcopentypes.FuncOpenTypeLingTongWeaponAdvanced,
	LingTongDevSysTypeLingQi:   funcopentypes.FuncOpenTypeLingTongMountAdvanced,
	LingTongDevSysTypeLingShen: funcopentypes.FuncOpenTypeLingTongShenFaAdvanced,
	LingTongDevSysTypeLingTi:   funcopentypes.FuncOpenTypeLingTongXianTiAdvanced,
	LingTongDevSysTypeLingYi:   funcopentypes.FuncOpenTypeLingTongWingAdvanced,
	LingTongDevSysTypeLingYu:   funcopentypes.FuncOpenTypeLingTongLingYuAdvanced,
}

func (t LingTongDevSysType) GetAdvanceFuncOpenType() funcopentypes.FuncOpenType {
	return lingTongDevSysTypeAdvanceFuncOpenMap[t]
}

const (
	LingTongDevSysTypeMin = LingTongDevSysTypeLingBing
	LingTongDevSysTypeMax = LingTongDevSysTypeLingTi
)

type LingTongDevType int32

const (
	//初始化
	LingTongDevTypeInit LingTongDevType = iota
	//进阶灵童
	LingTongDevTypeAdvanced
	//灵童皮肤
	LingTongDevTypeSkin
	//时效性
	LingTongDevTypeEffective
)

func (t LingTongDevType) Valid() bool {
	switch t {
	case LingTongDevTypeInit,
		LingTongDevTypeAdvanced,
		LingTongDevTypeSkin,
		LingTongDevTypeEffective:
		return true
	}
	return false
}

//幻化条件
type LingTongDevUCondType int32

const (
	//没有限制
	LingTongDevUCondTypeZ LingTongDevUCondType = iota
	//灵童阶别
	LingTongDevUCondTypeX
	//食用幻化丹数量
	LingTongDevUCondTypeU
	//消耗物品
	LingTongDevUCondTypeI
)

func (t LingTongDevUCondType) Valid() bool {
	switch t {
	case LingTongDevUCondTypeZ,
		LingTongDevUCondTypeX,
		LingTongDevUCondTypeU,
		LingTongDevUCondTypeI:
		return true
	}
	return false
}
