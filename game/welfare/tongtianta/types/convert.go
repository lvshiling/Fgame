package types

import (
	playerpropertytypes "fgame/fgame/game/property/player/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//属性到通天塔
var (
	convertToWelfareTongTianTaSubTypeMap = map[playerpropertytypes.PropertyEffectorType]welfaretypes.TongTianTaSubType{
		playerpropertytypes.PlayerPropertyEffectorTypeLingTong:        welfaretypes.TongTianTaSubTypeLingTong,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongFashion: welfaretypes.TongTianTaSubTypeLingTong,
		playerpropertytypes.PlayerPropertyEffectorTypeMingGe:          welfaretypes.TongTianTaSubTypeMingGe,
		playerpropertytypes.PlayerPropertyEffectorTypeTuLongEquip:     welfaretypes.TongTianTaSubTypeTuLong,
		playerpropertytypes.PlayerPropertyEffectorTypeShengHen:        welfaretypes.TongTianTaSubTypeShengHen,
		playerpropertytypes.PlayerPropertyEffectorTypeZhenFa:          welfaretypes.TongTianTaSubTypeZhenFa,
		playerpropertytypes.PlayerPropertyEffectorTypeBaby:            welfaretypes.TongTianTaSubTypeBaby,
		playerpropertytypes.PlayerPropertyEffectorTypeDianXing:        welfaretypes.TongTianTaSubTypeDianXing,
	}
)

func PlayerPropertyEffectTypeToTongTianTaSubType(effectType playerpropertytypes.PropertyEffectorType) (subType welfaretypes.TongTianTaSubType, isExist bool) {
	subType, isExist = convertToWelfareTongTianTaSubTypeMap[effectType]
	return
}

//通天塔到属性
var (
	convertToPropertyEffectTypeMap = map[welfaretypes.OpenActivitySubType][]playerpropertytypes.PropertyEffectorType{
		welfaretypes.TongTianTaSubTypeLingTong: []playerpropertytypes.PropertyEffectorType{playerpropertytypes.PlayerPropertyEffectorTypeLingTong, playerpropertytypes.PlayerPropertyEffectorTypeLingTongFashion},
		welfaretypes.TongTianTaSubTypeMingGe:   []playerpropertytypes.PropertyEffectorType{playerpropertytypes.PlayerPropertyEffectorTypeMingGe},
		welfaretypes.TongTianTaSubTypeTuLong:   []playerpropertytypes.PropertyEffectorType{playerpropertytypes.PlayerPropertyEffectorTypeTuLongEquip},
		welfaretypes.TongTianTaSubTypeShengHen: []playerpropertytypes.PropertyEffectorType{playerpropertytypes.PlayerPropertyEffectorTypeShengHen},
		welfaretypes.TongTianTaSubTypeZhenFa:   []playerpropertytypes.PropertyEffectorType{playerpropertytypes.PlayerPropertyEffectorTypeZhenFa},
		welfaretypes.TongTianTaSubTypeBaby:     []playerpropertytypes.PropertyEffectorType{playerpropertytypes.PlayerPropertyEffectorTypeBaby},
		welfaretypes.TongTianTaSubTypeDianXing: []playerpropertytypes.PropertyEffectorType{playerpropertytypes.PlayerPropertyEffectorTypeDianXing},
	}
)

func TongTianTaSubTypeTOPlayerPropertyEffectType(subType welfaretypes.OpenActivitySubType) (effectType []playerpropertytypes.PropertyEffectorType, isExist bool) {
	effectType, isExist = convertToPropertyEffectTypeMap[subType]
	return
}
