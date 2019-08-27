package types

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	ranktypes "fgame/fgame/game/rank/types"
)

//运营活动排行榜类型To系统排行榜类型
var (
	rankTypeMap = map[OpenActivityRankSubType]ranktypes.RankType{
		OpenActivityRankSubTypeCharge:           ranktypes.RankTypeCharge,
		OpenActivityRankSubTypeBodyshield:       ranktypes.RankTypeBodyShield,
		OpenActivityRankSubTypeCost:             ranktypes.RankTypeCost,
		OpenActivityRankSubTypeMount:            ranktypes.RankTypeMount,
		OpenActivityRankSubTypeWing:             ranktypes.RankTypeWing,
		OpenActivityRankSubTypeLingyu:           ranktypes.RankTypeLingYu,
		OpenActivityRankSubTypeShenfa:           ranktypes.RankTypeShenFa,
		OpenActivityRankSubTypeFeather:          ranktypes.RankTypeFeather,
		OpenActivityRankSubTypeShield:           ranktypes.RankTypeShield,
		OpenActivityRankSubTypeCharm:            ranktypes.RankTypeCharm,
		OpenActivityRankSubTypeAnqi:             ranktypes.RankTypeAnQi,
		OpenActivityRankSubTypeNumber:           ranktypes.RankTypeCount,
		OpenActivityRankSubTypeFaBao:            ranktypes.RankTypeFaBao,
		OpenActivityRankSubTypeXianTi:           ranktypes.RankTypeXianTi,
		OpenActivityRankSubTypeLevel:            ranktypes.RankTypeLevel,
		OpenActivityRankSubTypeShiHunFan:        ranktypes.RankTypeShiHunFan,
		OpenActivityRankSubTypeTianMoTi:         ranktypes.RankTypeTianMoTi,
		OpenActivityRankSubTypeLingBing:         ranktypes.RankTypeLingBing,
		OpenActivityRankSubTypeLingQi:           ranktypes.RankTypeLingQi,
		OpenActivityRankSubTypeLingYi:           ranktypes.RankTypeLingYi,
		OpenActivityRankSubTypeLingBao:          ranktypes.RankTypeLingBao,
		OpenActivityRankSubTypeLingTi:           ranktypes.RankTypeLingTi,
		OpenActivityRankSubTypeLingYu:           ranktypes.RankTypeLingTongYu,
		OpenActivityRankSubTypeLingShen:         ranktypes.RankTypeLingShen,
		OpenActivityRankSubTypeMarryDevelop:     ranktypes.RankTypeMarryDevelop,
		OpenActivityRankSubTypeNumberDay:        ranktypes.RankTypeCount,
		OpenActivityRankSubTypeLingTongForce:    ranktypes.RankTypeLingTongForce,
		OpenActivityRankSubTypeGoldEquipForce:   ranktypes.RankTypeGoldEquipForce,
		OpenActivityRankSubTypeDianXingForce:    ranktypes.RankTypeDianXingForce,
		OpenActivityRankSubTypeShenQiForce:      ranktypes.RankTypeShenQiForce,
		OpenActivityRankSubTypeMingGeForce:      ranktypes.RankTypeMingGeForce,
		OpenActivityRankSubTypeShengHenForce:    ranktypes.RankTypeShengHenForce,
		OpenActivityRankSubTypeZhenFaForce:      ranktypes.RankTypeZhenFaForce,
		OpenActivityRankSubTypeTuLongEquipForce: ranktypes.RankTypeTuLongEquipForce,
		OpenActivityRankSubTypeBabyForce:        ranktypes.RankTypeBabyForce,
		OpenActivityRankSubTypeZhuanSheng:       ranktypes.RankTypeZhuanSheng,
	}
)

func (t OpenActivityRankSubType) RankType() ranktypes.RankType {
	return rankTypeMap[t]
}

// 运营活动系统类型To灵童系统类型
var (
	convertTolingTongTypeMap = map[AdvancedType]lingtongdevtypes.LingTongDevSysType{
		AdvancedTypeLingBao:  lingtongdevtypes.LingTongDevSysTypeLingBao,
		AdvancedTypeLingBing: lingtongdevtypes.LingTongDevSysTypeLingBing,
		AdvancedTypeLingQi:   lingtongdevtypes.LingTongDevSysTypeLingQi,
		AdvancedTypeLingShen: lingtongdevtypes.LingTongDevSysTypeLingShen,
		AdvancedTypeLingTi:   lingtongdevtypes.LingTongDevSysTypeLingTi,
		AdvancedTypeLingYi:   lingtongdevtypes.LingTongDevSysTypeLingYi,
		AdvancedTypeLingYu:   lingtongdevtypes.LingTongDevSysTypeLingYu,
	}
)

func AdvancedTypeToLingTongDevType(adType AdvancedType) (lingTongDevType lingtongdevtypes.LingTongDevSysType, isExist bool) {
	lingTongDevType, isExist = convertTolingTongTypeMap[adType]
	return
}

//灵童系统类型To运营活动系统类型
var (
	convertToWelfareAdvancedTypeMap = map[lingtongdevtypes.LingTongDevSysType]AdvancedType{
		lingtongdevtypes.LingTongDevSysTypeLingBao:  AdvancedTypeLingBao,
		lingtongdevtypes.LingTongDevSysTypeLingBing: AdvancedTypeLingBing,
		lingtongdevtypes.LingTongDevSysTypeLingQi:   AdvancedTypeLingQi,
		lingtongdevtypes.LingTongDevSysTypeLingShen: AdvancedTypeLingShen,
		lingtongdevtypes.LingTongDevSysTypeLingTi:   AdvancedTypeLingTi,
		lingtongdevtypes.LingTongDevSysTypeLingYi:   AdvancedTypeLingYi,
		lingtongdevtypes.LingTongDevSysTypeLingYu:   AdvancedTypeLingYu,
	}
)

func LingTongDevTypeToAdvancedType(lingTongDevType lingtongdevtypes.LingTongDevSysType) (adType AdvancedType, isExist bool) {
	adType, isExist = convertToWelfareAdvancedTypeMap[lingTongDevType]
	return
}

//功能开启类型To进阶类型
var (
	convertToAdvancedTypeMap = map[funcopentypes.FuncOpenType]AdvancedType{
		funcopentypes.FuncOpenTypeMount:          AdvancedTypeMount,
		funcopentypes.FuncOpenTypeWing:           AdvancedTypeWing,
		funcopentypes.FuncOpenTypeAnQi:           AdvancedTypeAnqi,
		funcopentypes.FuncOpenTypeBodyShield:     AdvancedTypeBodyshield,
		funcopentypes.FuncOpenTypeLingYu:         AdvancedTypeLingyu,
		funcopentypes.FuncOpenTypeShenfa:         AdvancedTypeShenfa,
		funcopentypes.FuncOpenTypeShield:         AdvancedTypeShield,
		funcopentypes.FuncOpenTypeFaBao:          AdvancedTypeFaBao,
		funcopentypes.FuncOpenTypeXianTi:         AdvancedTypeXianTi,
		funcopentypes.FuncOpenTypeFeather:        AdvancedTypeFeather,
		funcopentypes.FuncOpenTypeShiHunFan:      AdvancedTypeShiHunFan,
		funcopentypes.FuncOpenTypeTianMo:         AdvancedTypeTianMoTi,
		funcopentypes.FuncOpenTypeLingTongWeapon: AdvancedTypeLingBing,
		funcopentypes.FuncOpenTypeLingTongMount:  AdvancedTypeLingQi,
		funcopentypes.FuncOpenTypeLingTongWing:   AdvancedTypeLingYi,
		funcopentypes.FuncOpenTypeLingTongFaBao:  AdvancedTypeLingBao,
		funcopentypes.FuncOpenTypeLingTongXianTi: AdvancedTypeLingTi,
		funcopentypes.FuncOpenTypeLingTongLingYu: AdvancedTypeLingYu,
		funcopentypes.FuncOpenTypeLingTongShenFa: AdvancedTypeLingShen,
	}
)

func FuncTypeToAdvancedType(funcType funcopentypes.FuncOpenType) (adType AdvancedType, isExist bool) {
	adType, isExist = convertToAdvancedTypeMap[funcType]
	return
}
