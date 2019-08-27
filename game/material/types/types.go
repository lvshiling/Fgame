package types

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
)

type MaterialType int32

const (
	//坐骑副本
	MaterialTypeMount MaterialType = iota + 1
	//战翼副本
	MaterialTypeWing
	//身法副本
	MaterialTypeShenfa
	//法宝副本
	MaterialTypeFabao
	//仙体副本
	MaterialTypeXianti
	//6天魔副本
	MaterialTypeTianMo
	//灵童副本
	MaterialTypeLingTong
	//灵兵副本
	MaterialTypeLingBing
	//灵域副本
	MaterialTypeLingYu
)

var materialTypeMap = map[MaterialType]string{
	MaterialTypeMount:    "坐骑副本",
	MaterialTypeWing:     "战翼副本",
	MaterialTypeXianti:   "仙体副本",
	MaterialTypeFabao:    "法宝副本",
	MaterialTypeShenfa:   "身法副本",
	MaterialTypeTianMo:   "天魔副本",
	MaterialTypeLingTong: "灵童副本",
	MaterialTypeLingBing: "灵兵副本",
	MaterialTypeLingYu:   "灵域副本",
}

func (t MaterialType) Valid() bool {
	switch t {
	case MaterialTypeMount,
		MaterialTypeWing,
		MaterialTypeXianti,
		MaterialTypeFabao,
		MaterialTypeShenfa,
		MaterialTypeTianMo,
		MaterialTypeLingTong,
		MaterialTypeLingBing,
		MaterialTypeLingYu:
		return true
	default:
		return false
	}
}

func (t MaterialType) String() string {
	return materialTypeMap[t]
}

func (t MaterialType) GetFuncOpenType() funcopentypes.FuncOpenType {
	return funcopenMap[t]
}

func GetMaterialMap() map[MaterialType]string {
	return materialTypeMap
}

const (
	MinMaterialType = MaterialTypeMount
	MaxMaterialType = MaterialTypeLingYu
)

var (
	funcopenMap = map[MaterialType]funcopentypes.FuncOpenType{
		MaterialTypeMount:    funcopentypes.FuncOpenTypeCaiLiaoMount,
		MaterialTypeWing:     funcopentypes.FuncOpenTypeCaiLiaoWing,
		MaterialTypeXianti:   funcopentypes.FuncOpenTypeCaiLiaoXianTi,
		MaterialTypeFabao:    funcopentypes.FuncOpenTypeCaiLiaoFaBao,
		MaterialTypeShenfa:   funcopentypes.FuncOpenTypeCaiLiaoShenFa,
		MaterialTypeTianMo:   funcopentypes.FuncOpenTypeCaiLiaoTianMo,
		MaterialTypeLingTong: funcopentypes.FuncOpenTypeCaiLiaoLingTong,
		MaterialTypeLingBing: funcopentypes.FuncOpenTypeCaiLiaoLingBing,
		MaterialTypeLingYu:   funcopentypes.FuncOpenTypeCaiLiaoLingYu,
	}
)
