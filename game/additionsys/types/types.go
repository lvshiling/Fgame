package types

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	itemtypes "fgame/fgame/game/item/types"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

type AdditionSysType int32

const (
	//坐骑装备
	AdditionSysTypeMountEquip AdditionSysType = 1 + iota
	//战翼符石
	AdditionSysTypeWingStone AdditionSysType = 2
	//暗器机关
	AdditionSysTypeAnqiJiguan AdditionSysType = 3
	//法宝装备
	AdditionSysTypeFaBao AdditionSysType = 4
	//仙体装备
	AdditionSysTypeXianTi AdditionSysType = 5
	//领域装备
	AdditionSysTypeLingYu AdditionSysType = 6
	//身法装备
	AdditionSysTypeShenFa AdditionSysType = 7
	//噬魂幡装备
	AdditionSysTypeShiHunFan AdditionSysType = 8
	//天魔体装备
	AdditionSysTypeTianMoTi AdditionSysType = 9
	//灵童兵魂系统
	AdditionSysTypeLingTongWeapon AdditionSysType = 101
	//灵童坐骑系统
	AdditionSysTypeLingTongMount AdditionSysType = 102
	//灵童战翼系统
	AdditionSysTypeLingTongWing AdditionSysType = 103
	//灵童身法系统
	AdditionSysTypeLingTongShenFa AdditionSysType = 104
	//灵童领域系统
	AdditionSysTypeLingTongLingYu AdditionSysType = 105
	//灵童法宝系统
	AdditionSysTypeLingTongFaBao AdditionSysType = 106
	//灵童仙体系统
	AdditionSysTypeLingTongXianTi AdditionSysType = 107
	//圣痕-青龙
	AdditionSysTypeShengHenQingLong AdditionSysType = 201
	//圣痕-白虎
	AdditionSysTypeShengHenBaiHu AdditionSysType = 202
	//圣痕-朱雀
	AdditionSysTypeShengHenZhuQue AdditionSysType = 203
	//圣痕-玄武
	AdditionSysTypeShengHenXuanWu AdditionSysType = 204
	//灵童1装备
	AdditionSysTypeLingTongOneEquip AdditionSysType = 301
	//灵童2装备
	AdditionSysTypeLingTongTwoEquip AdditionSysType = 302
	//灵童3装备
	AdditionSysTypeLingTongThreeEquip AdditionSysType = 303
	//灵童4装备
	AdditionSysTypeLingTongFourEquip AdditionSysType = 304
	//灵童5装备
	AdditionSysTypeLingTongFiveEquip AdditionSysType = 305
	//灵童6装备
	AdditionSysTypeLingTongSixEquip AdditionSysType = 306
	//灵童7装备
	AdditionSysTypeLingTongSevenEquip AdditionSysType = 307
	//灵童8装备
	AdditionSysTypeLingTongEightEquip AdditionSysType = 308
	//灵童9装备
	AdditionSysTypeLingTongNineEquip AdditionSysType = 309
	//灵童10装备
	AdditionSysTypeLingTongTenEquip AdditionSysType = 310
)

const (
	MinType = AdditionSysTypeMountEquip
	MaxType = AdditionSysTypeLingTongTenEquip
)

var (
	additionSysTypeMap = map[AdditionSysType]string{
		AdditionSysTypeMountEquip:         "坐骑装备",
		AdditionSysTypeWingStone:          "战翼符石",
		AdditionSysTypeAnqiJiguan:         "暗器机关",
		AdditionSysTypeFaBao:              "法宝装备",
		AdditionSysTypeXianTi:             "仙体装备",
		AdditionSysTypeLingYu:             "领域装备",
		AdditionSysTypeShenFa:             "身法装备",
		AdditionSysTypeShiHunFan:          "噬魂幡装备",
		AdditionSysTypeTianMoTi:           "天魔体装备",
		AdditionSysTypeLingTongWeapon:     "灵童兵魂系统",
		AdditionSysTypeLingTongMount:      "灵童坐骑系统",
		AdditionSysTypeLingTongWing:       "灵童战翼系统",
		AdditionSysTypeLingTongShenFa:     "灵童身法系统",
		AdditionSysTypeLingTongLingYu:     "灵童领域系统",
		AdditionSysTypeLingTongFaBao:      "灵童法宝系统",
		AdditionSysTypeLingTongXianTi:     "灵童仙体系统",
		AdditionSysTypeLingTongOneEquip:   "灵童1装备系统",
		AdditionSysTypeLingTongTwoEquip:   "灵童2装备系统",
		AdditionSysTypeLingTongThreeEquip: "灵童3装备系统",
		AdditionSysTypeLingTongFourEquip:  "灵童4装备系统",
		AdditionSysTypeLingTongFiveEquip:  "灵童5装备系统",
		AdditionSysTypeLingTongSixEquip:   "灵童6装备系统",
		AdditionSysTypeLingTongSevenEquip: "灵童7装备系统",
		AdditionSysTypeLingTongEightEquip: "灵童8装备系统",
		AdditionSysTypeLingTongNineEquip:  "灵童9装备系统",
		AdditionSysTypeLingTongTenEquip:   "灵童10装备系统",
		AdditionSysTypeShengHenQingLong:   "圣痕青龙系统",
		AdditionSysTypeShengHenBaiHu:      "圣痕白虎系统",
		AdditionSysTypeShengHenZhuQue:     "圣痕朱雀系统",
		AdditionSysTypeShengHenXuanWu:     "圣痕玄武系统",
	}
)

func (spt AdditionSysType) Valid() bool {
	switch spt {
	case AdditionSysTypeMountEquip,
		AdditionSysTypeWingStone,
		AdditionSysTypeAnqiJiguan,
		AdditionSysTypeFaBao,
		AdditionSysTypeXianTi,
		AdditionSysTypeLingYu,
		AdditionSysTypeShenFa,
		AdditionSysTypeShiHunFan,
		AdditionSysTypeTianMoTi,
		AdditionSysTypeLingTongWeapon,
		AdditionSysTypeLingTongMount,
		AdditionSysTypeLingTongWing,
		AdditionSysTypeLingTongShenFa,
		AdditionSysTypeLingTongLingYu,
		AdditionSysTypeLingTongFaBao,
		AdditionSysTypeLingTongXianTi,
		AdditionSysTypeLingTongOneEquip,
		AdditionSysTypeLingTongTwoEquip,
		AdditionSysTypeLingTongThreeEquip,
		AdditionSysTypeLingTongFourEquip,
		AdditionSysTypeLingTongFiveEquip,
		AdditionSysTypeLingTongSixEquip,
		AdditionSysTypeLingTongSevenEquip,
		AdditionSysTypeLingTongEightEquip,
		AdditionSysTypeLingTongNineEquip,
		AdditionSysTypeLingTongTenEquip,
		AdditionSysTypeShengHenQingLong,
		AdditionSysTypeShengHenBaiHu,
		AdditionSysTypeShengHenZhuQue,
		AdditionSysTypeShengHenXuanWu:
		return true
	}
	return false
}

func (spt AdditionSysType) String() string {
	return additionSysTypeMap[spt]
}

// 转换属性作用器类型
var propertyEffectTypeMap = map[AdditionSysType]playerpropertytypes.PropertyEffectorType{
	AdditionSysTypeMountEquip:         playerpropertytypes.PlayerPropertyEffectorTypeMount,
	AdditionSysTypeWingStone:          playerpropertytypes.PlayerPropertyEffectorTypeWing,
	AdditionSysTypeAnqiJiguan:         playerpropertytypes.PlayerPropertyEffectorTypeAnqi,
	AdditionSysTypeFaBao:              playerpropertytypes.PlayerPropertyEffectorTypeFaBao,
	AdditionSysTypeXianTi:             playerpropertytypes.PlayerPropertyEffectorTypeXianTi,
	AdditionSysTypeLingYu:             playerpropertytypes.PlayerPropertyEffectorTypeLingyu,
	AdditionSysTypeShenFa:             playerpropertytypes.PlayerPropertyEffectorTypeShenfa,
	AdditionSysTypeShiHunFan:          playerpropertytypes.PlayerPropertyEffectorTypeShiHunFan,
	AdditionSysTypeTianMoTi:           playerpropertytypes.PlayerPropertyEffectorTypeTianMoTi,
	AdditionSysTypeLingTongWeapon:     playerpropertytypes.PlayerPropertyEffectorTypeLingTongWeapon,
	AdditionSysTypeLingTongMount:      playerpropertytypes.PlayerPropertyEffectorTypeLingTongMount,
	AdditionSysTypeLingTongWing:       playerpropertytypes.PlayerPropertyEffectorTypeLingTongWing,
	AdditionSysTypeLingTongShenFa:     playerpropertytypes.PlayerPropertyEffectorTypeLingTongShenFa,
	AdditionSysTypeLingTongLingYu:     playerpropertytypes.PlayerPropertyEffectorTypeLingTongLingYu,
	AdditionSysTypeLingTongFaBao:      playerpropertytypes.PlayerPropertyEffectorTypeLingTongFaBao,
	AdditionSysTypeLingTongXianTi:     playerpropertytypes.PlayerPropertyEffectorTypeLingTongXianTi,
	AdditionSysTypeLingTongOneEquip:   playerpropertytypes.PlayerPropertyEffectorTypeLingTong,
	AdditionSysTypeLingTongTwoEquip:   playerpropertytypes.PlayerPropertyEffectorTypeLingTong,
	AdditionSysTypeLingTongThreeEquip: playerpropertytypes.PlayerPropertyEffectorTypeLingTong,
	AdditionSysTypeLingTongFourEquip:  playerpropertytypes.PlayerPropertyEffectorTypeLingTong,
	AdditionSysTypeLingTongFiveEquip:  playerpropertytypes.PlayerPropertyEffectorTypeLingTong,
	AdditionSysTypeLingTongSixEquip:   playerpropertytypes.PlayerPropertyEffectorTypeLingTong,
	AdditionSysTypeLingTongSevenEquip: playerpropertytypes.PlayerPropertyEffectorTypeLingTong,
	AdditionSysTypeLingTongEightEquip: playerpropertytypes.PlayerPropertyEffectorTypeLingTong,
	AdditionSysTypeLingTongNineEquip:  playerpropertytypes.PlayerPropertyEffectorTypeLingTong,
	AdditionSysTypeLingTongTenEquip:   playerpropertytypes.PlayerPropertyEffectorTypeLingTong,
	AdditionSysTypeShengHenQingLong:   playerpropertytypes.PlayerPropertyEffectorTypeShengHen,
	AdditionSysTypeShengHenBaiHu:      playerpropertytypes.PlayerPropertyEffectorTypeShengHen,
	AdditionSysTypeShengHenZhuQue:     playerpropertytypes.PlayerPropertyEffectorTypeShengHen,
	AdditionSysTypeShengHenXuanWu:     playerpropertytypes.PlayerPropertyEffectorTypeShengHen,
}

func (t AdditionSysType) ConvertToPropertyEffectType() (playerpropertytypes.PropertyEffectorType, bool) {
	typ, ok := propertyEffectTypeMap[t]
	return typ, ok
}

// 转换觉醒功能开启类型
var awakeOpenTypeMap = map[AdditionSysType]funcopentypes.FuncOpenType{
	AdditionSysTypeMountEquip:     funcopentypes.FuncOpenTypeJueXingMount,
	AdditionSysTypeWingStone:      funcopentypes.FuncOpenTypeJueXingWing,
	AdditionSysTypeAnqiJiguan:     funcopentypes.FuncOpenTypeJueXingAnQi,
	AdditionSysTypeFaBao:          funcopentypes.FuncOpenTypeJueXingFaBao,
	AdditionSysTypeXianTi:         funcopentypes.FuncOpenTypeJueXingXianTi,
	AdditionSysTypeLingYu:         funcopentypes.FuncOpenTypeJueXingLingYu,
	AdditionSysTypeShenFa:         funcopentypes.FuncOpenTypeJueXingShenFa,
	AdditionSysTypeShiHunFan:      funcopentypes.FuncOpenTypeJueXingShiHunFan,
	AdditionSysTypeTianMoTi:       funcopentypes.FuncOpenTypeJueXingTianMoTi,
	AdditionSysTypeLingTongWeapon: funcopentypes.FuncOpenTypeJueXingLingTongWeapon,
	AdditionSysTypeLingTongMount:  funcopentypes.FuncOpenTypeJueXingLingTongMount,
	AdditionSysTypeLingTongWing:   funcopentypes.FuncOpenTypeJueXingLingTongWing,
	AdditionSysTypeLingTongShenFa: funcopentypes.FuncOpenTypeJueXingLingTongShenFa,
	AdditionSysTypeLingTongLingYu: funcopentypes.FuncOpenTypeJueXingLingTongLingYu,
	AdditionSysTypeLingTongFaBao:  funcopentypes.FuncOpenTypeJueXingLingTongFaBao,
	AdditionSysTypeLingTongXianTi: funcopentypes.FuncOpenTypeJueXingLingTongXianTi,
}

func (t AdditionSysType) ConvertToAwakeFuncOpenType() (funcopentypes.FuncOpenType, bool) {
	typ, ok := awakeOpenTypeMap[t]
	return typ, ok
}

// 转换化灵功能开启类型
var huaLingOpenTypeMap = map[AdditionSysType]funcopentypes.FuncOpenType{
	AdditionSysTypeMountEquip:     funcopentypes.FuncOpenTypeHuaLingMount,
	AdditionSysTypeWingStone:      funcopentypes.FuncOpenTypeHuaLingWing,
	AdditionSysTypeAnqiJiguan:     funcopentypes.FuncOpenTypeHuaLingAnQi,
	AdditionSysTypeFaBao:          funcopentypes.FuncOpenTypeHuaLingFaBao,
	AdditionSysTypeXianTi:         funcopentypes.FuncOpenTypeHuaLingXianTi,
	AdditionSysTypeLingYu:         funcopentypes.FuncOpenTypeHuaLingLingYu,
	AdditionSysTypeShenFa:         funcopentypes.FuncOpenTypeHuaLingShenFa,
	AdditionSysTypeShiHunFan:      funcopentypes.FuncOpenTypeHuaLingShiHunFan,
	AdditionSysTypeTianMoTi:       funcopentypes.FuncOpenTypeHuaLingTianMoTi,
	AdditionSysTypeLingTongWeapon: funcopentypes.FuncOpenTypeHuaLingLingBing,
	AdditionSysTypeLingTongMount:  funcopentypes.FuncOpenTypeHuaLingLingQi,
	AdditionSysTypeLingTongWing:   funcopentypes.FuncOpenTypeHuaLingLingWing,
	AdditionSysTypeLingTongShenFa: funcopentypes.FuncOpenTypeHuaLingLingShen,
	AdditionSysTypeLingTongLingYu: funcopentypes.FuncOpenTypeHuaLingLingArea,
	AdditionSysTypeLingTongFaBao:  funcopentypes.FuncOpenTypeHuaLingLingBao,
	AdditionSysTypeLingTongXianTi: funcopentypes.FuncOpenTypeHuaLingLingTi,
}

func (t AdditionSysType) ConvertToHuaLingFuncOpenType() (funcopentypes.FuncOpenType, bool) {
	typ, ok := huaLingOpenTypeMap[t]
	return typ, ok
}

// 转换神铸功能开启类型
var shenZhuOpenTypeMap = map[AdditionSysType]funcopentypes.FuncOpenType{
	AdditionSysTypeMountEquip:         funcopentypes.FuncOpenTypeShenZhuMount,
	AdditionSysTypeWingStone:          funcopentypes.FuncOpenTypeShenZhuWing,
	AdditionSysTypeAnqiJiguan:         funcopentypes.FuncOpenTypeShenZhuAnQi,
	AdditionSysTypeFaBao:              funcopentypes.FuncOpenTypeShenZhuFaBao,
	AdditionSysTypeXianTi:             funcopentypes.FuncOpenTypeShenZhuXianTi,
	AdditionSysTypeLingYu:             funcopentypes.FuncOpenTypeShenZhuLingYu,
	AdditionSysTypeShenFa:             funcopentypes.FuncOpenTypeShenZhuShenFa,
	AdditionSysTypeShiHunFan:          funcopentypes.FuncOpenTypeShenZhuShiHunFan,
	AdditionSysTypeTianMoTi:           funcopentypes.FuncOpenTypeShenZhuTianMoTi,
	AdditionSysTypeLingTongWeapon:     funcopentypes.FuncOpenTypeShenZhuLingTongWeapon,
	AdditionSysTypeLingTongMount:      funcopentypes.FuncOpenTypeShenZhuLingTongMount,
	AdditionSysTypeLingTongWing:       funcopentypes.FuncOpenTypeShenZhuLingTongWing,
	AdditionSysTypeLingTongShenFa:     funcopentypes.FuncOpenTypeShenZhuLingTongShenFa,
	AdditionSysTypeLingTongLingYu:     funcopentypes.FuncOpenTypeShenZhuLingTongLingYu,
	AdditionSysTypeLingTongFaBao:      funcopentypes.FuncOpenTypeShenZhuLingTongFaBao,
	AdditionSysTypeLingTongXianTi:     funcopentypes.FuncOpenTypeShenZhuLingTongXianTi,
	AdditionSysTypeLingTongOneEquip:   funcopentypes.FuncOpenTypeShenZhuLingTongEquip,
	AdditionSysTypeLingTongTwoEquip:   funcopentypes.FuncOpenTypeShenZhuLingTongEquip,
	AdditionSysTypeLingTongThreeEquip: funcopentypes.FuncOpenTypeShenZhuLingTongEquip,
	AdditionSysTypeLingTongFourEquip:  funcopentypes.FuncOpenTypeShenZhuLingTongEquip,
	AdditionSysTypeLingTongFiveEquip:  funcopentypes.FuncOpenTypeShenZhuLingTongEquip,
	AdditionSysTypeLingTongSixEquip:   funcopentypes.FuncOpenTypeShenZhuLingTongEquip,
	AdditionSysTypeLingTongSevenEquip: funcopentypes.FuncOpenTypeShenZhuLingTongEquip,
	AdditionSysTypeLingTongEightEquip: funcopentypes.FuncOpenTypeShenZhuLingTongEquip,
	AdditionSysTypeLingTongNineEquip:  funcopentypes.FuncOpenTypeShenZhuLingTongEquip,
	AdditionSysTypeLingTongTenEquip:   funcopentypes.FuncOpenTypeShenZhuLingTongEquip,
}

func (t AdditionSysType) ConvertToShenZhuFuncOpenType() (funcopentypes.FuncOpenType, bool) {
	typ, ok := shenZhuOpenTypeMap[t]
	return typ, ok
}

// 转换通灵功能开启类型
var tongLingOpenTypeMap = map[AdditionSysType]funcopentypes.FuncOpenType{
	AdditionSysTypeMountEquip:     funcopentypes.FuncOpenTypeTongLingMount,
	AdditionSysTypeWingStone:      funcopentypes.FuncOpenTypeTongLingWing,
	AdditionSysTypeAnqiJiguan:     funcopentypes.FuncOpenTypeTongLingAnQi,
	AdditionSysTypeFaBao:          funcopentypes.FuncOpenTypeTongLingFaBao,
	AdditionSysTypeXianTi:         funcopentypes.FuncOpenTypeTongLingXianTi,
	AdditionSysTypeLingYu:         funcopentypes.FuncOpenTypeTongLingLingYu,
	AdditionSysTypeShenFa:         funcopentypes.FuncOpenTypeTongLingShenFa,
	AdditionSysTypeShiHunFan:      funcopentypes.FuncOpenTypeTongLingShiHunFan,
	AdditionSysTypeTianMoTi:       funcopentypes.FuncOpenTypeTongLingTianMoTi,
	AdditionSysTypeLingTongWeapon: funcopentypes.FuncOpenTypeTongLingLingTongWeapon,
	AdditionSysTypeLingTongMount:  funcopentypes.FuncOpenTypeTongLingLingTongMount,
	AdditionSysTypeLingTongWing:   funcopentypes.FuncOpenTypeTongLingLingTongWing,
	AdditionSysTypeLingTongShenFa: funcopentypes.FuncOpenTypeTongLingLingTongShenFa,
	AdditionSysTypeLingTongLingYu: funcopentypes.FuncOpenTypeTongLingLingTongLingYu,
	AdditionSysTypeLingTongFaBao:  funcopentypes.FuncOpenTypeTongLingLingTongFaBao,
	AdditionSysTypeLingTongXianTi: funcopentypes.FuncOpenTypeTongLingLingTongXianTi,
}

func (t AdditionSysType) ConvertToTongLingFuncOpenType() (funcopentypes.FuncOpenType, bool) {
	typ, ok := tongLingOpenTypeMap[t]
	return typ, ok
}

// 转换 功能开启类型
var openTypeMap = map[AdditionSysType]funcopentypes.FuncOpenType{
	AdditionSysTypeMountEquip:         funcopentypes.FuncOpenTypeMount,
	AdditionSysTypeWingStone:          funcopentypes.FuncOpenTypeWing,
	AdditionSysTypeAnqiJiguan:         funcopentypes.FuncOpenTypeAnQi,
	AdditionSysTypeFaBao:              funcopentypes.FuncOpenTypeFaBao,
	AdditionSysTypeXianTi:             funcopentypes.FuncOpenTypeXianTi,
	AdditionSysTypeLingYu:             funcopentypes.FuncOpenTypeLingYu,
	AdditionSysTypeShenFa:             funcopentypes.FuncOpenTypeShenfa,
	AdditionSysTypeShiHunFan:          funcopentypes.FuncOpenTypeShiHunFan,
	AdditionSysTypeTianMoTi:           funcopentypes.FuncOpenTypeTianMo,
	AdditionSysTypeLingTongWeapon:     funcopentypes.FuncOpenTypeLingTongWeapon,
	AdditionSysTypeLingTongMount:      funcopentypes.FuncOpenTypeLingTongMount,
	AdditionSysTypeLingTongWing:       funcopentypes.FuncOpenTypeLingTongWing,
	AdditionSysTypeLingTongShenFa:     funcopentypes.FuncOpenTypeLingTongShenFa,
	AdditionSysTypeLingTongLingYu:     funcopentypes.FuncOpenTypeLingTongLingYu,
	AdditionSysTypeLingTongFaBao:      funcopentypes.FuncOpenTypeLingTongFaBao,
	AdditionSysTypeLingTongXianTi:     funcopentypes.FuncOpenTypeLingTongXianTi,
	AdditionSysTypeShengHenQingLong:   funcopentypes.FuncOpenTypeShengHen,
	AdditionSysTypeShengHenBaiHu:      funcopentypes.FuncOpenTypeShengHen,
	AdditionSysTypeShengHenZhuQue:     funcopentypes.FuncOpenTypeShengHen,
	AdditionSysTypeShengHenXuanWu:     funcopentypes.FuncOpenTypeShengHen,
	AdditionSysTypeLingTongOneEquip:   funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongTwoEquip:   funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongThreeEquip: funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongFourEquip:  funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongFiveEquip:  funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongSixEquip:   funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongSevenEquip: funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongEightEquip: funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongNineEquip:  funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongTenEquip:   funcopentypes.FuncOpenTypeLingTongEquip,
}

func (t AdditionSysType) ConvertToFuncOpenType() (funcopentypes.FuncOpenType, bool) {
	typ, ok := openTypeMap[t]
	return typ, ok
}

// 转换 附加系统装备功能开启类型
var equipOpenTypeMap = map[AdditionSysType]funcopentypes.FuncOpenType{
	AdditionSysTypeMountEquip:         funcopentypes.FuncOpenTypeMountEquipment,
	AdditionSysTypeWingStone:          funcopentypes.FuncOpenTypeWingRune,
	AdditionSysTypeAnqiJiguan:         funcopentypes.FuncOpenTypeAnQiJiGuan,
	AdditionSysTypeFaBao:              funcopentypes.FuncOpenTypeFaBaoSuit,
	AdditionSysTypeXianTi:             funcopentypes.FuncOpenTypeXianTiLingYu,
	AdditionSysTypeLingYu:             funcopentypes.FuncOpenTypeLingYuEquip,
	AdditionSysTypeShenFa:             funcopentypes.FuncOpenTypeShenFaEquip,
	AdditionSysTypeShiHunFan:          funcopentypes.FuncOpenTypeShiHunFanEquip,
	AdditionSysTypeTianMoTi:           funcopentypes.FuncOpenTypeTianMoEquipment,
	AdditionSysTypeLingTongWeapon:     funcopentypes.FuncOpenTypeLingTongWeaponEquip,
	AdditionSysTypeLingTongMount:      funcopentypes.FuncOpenTypeLingTongMountEquip,
	AdditionSysTypeLingTongWing:       funcopentypes.FuncOpenTypeLingTongWingEquip,
	AdditionSysTypeLingTongShenFa:     funcopentypes.FuncOpenTypeLingTongShenFaEquip,
	AdditionSysTypeLingTongLingYu:     funcopentypes.FuncOpenTypeLingTongLingYuEquip,
	AdditionSysTypeLingTongFaBao:      funcopentypes.FuncOpenTypeLingTongFaBaoEquip,
	AdditionSysTypeLingTongXianTi:     funcopentypes.FuncOpenTypeLingTongXianTiEquip,
	AdditionSysTypeShengHenQingLong:   funcopentypes.FuncOpenTypeShengHen,
	AdditionSysTypeShengHenBaiHu:      funcopentypes.FuncOpenTypeShengHen,
	AdditionSysTypeShengHenZhuQue:     funcopentypes.FuncOpenTypeShengHen,
	AdditionSysTypeShengHenXuanWu:     funcopentypes.FuncOpenTypeShengHen,
	AdditionSysTypeLingTongOneEquip:   funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongTwoEquip:   funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongThreeEquip: funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongFourEquip:  funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongFiveEquip:  funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongSixEquip:   funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongSevenEquip: funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongEightEquip: funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongNineEquip:  funcopentypes.FuncOpenTypeLingTongEquip,
	AdditionSysTypeLingTongTenEquip:   funcopentypes.FuncOpenTypeLingTongEquip,
}

func (t AdditionSysType) ConvertToEquipFuncOpenType() (funcopentypes.FuncOpenType, bool) {
	typ, ok := equipOpenTypeMap[t]
	return typ, ok
}

// 转换 系统升级功能开启类型
var shengJiOpenTypeMap = map[AdditionSysType]funcopentypes.FuncOpenType{
	AdditionSysTypeMountEquip:       funcopentypes.FuncOpenTypeMountUpgrade,
	AdditionSysTypeWingStone:        funcopentypes.FuncOpenTypeWingShengJi,
	AdditionSysTypeAnqiJiguan:       funcopentypes.FuncOpenTypeAnQiShengJi,
	AdditionSysTypeFaBao:            funcopentypes.FuncOpenTypeFaBaoShengJi,
	AdditionSysTypeXianTi:           funcopentypes.FuncOpenTypeXianTiShengJi,
	AdditionSysTypeLingYu:           funcopentypes.FuncOpenTypeLingYuShengJi,
	AdditionSysTypeShenFa:           funcopentypes.FuncOpenTypeShenfaShengJi,
	AdditionSysTypeShiHunFan:        funcopentypes.FuncOpenTypeShiHunFanShengJi,
	AdditionSysTypeTianMoTi:         funcopentypes.FuncOpenTypeTianMoUplevel,
	AdditionSysTypeLingTongWeapon:   funcopentypes.FuncOpenTypeLingTongWeaponUpgrade,
	AdditionSysTypeLingTongMount:    funcopentypes.FuncOpenTypeLingTongMountUpgrade,
	AdditionSysTypeLingTongWing:     funcopentypes.FuncOpenTypeLingTongWingUpgrade,
	AdditionSysTypeLingTongShenFa:   funcopentypes.FuncOpenTypeLingTongShenFaUpgrade,
	AdditionSysTypeLingTongLingYu:   funcopentypes.FuncOpenTypeLingTongLingYuUpgrade,
	AdditionSysTypeLingTongFaBao:    funcopentypes.FuncOpenTypeLingTongFaBaoUpgrade,
	AdditionSysTypeLingTongXianTi:   funcopentypes.FuncOpenTypeLingTongXianTiUpgrade,
	AdditionSysTypeShengHenQingLong: funcopentypes.FuncOpenTypeShengHen,
	AdditionSysTypeShengHenBaiHu:    funcopentypes.FuncOpenTypeShengHen,
	AdditionSysTypeShengHenZhuQue:   funcopentypes.FuncOpenTypeShengHen,
	AdditionSysTypeShengHenXuanWu:   funcopentypes.FuncOpenTypeShengHen,
}

func (t AdditionSysType) ConvertToShengJiFuncOpenType() (funcopentypes.FuncOpenType, bool) {
	typ, ok := shengJiOpenTypeMap[t]
	return typ, ok
}

// 转换 物品类型
var itemTypeMap = map[AdditionSysType]itemtypes.ItemType{
	AdditionSysTypeMountEquip:         itemtypes.ItemTypeMountEquip,
	AdditionSysTypeWingStone:          itemtypes.ItemTypeWingStone,
	AdditionSysTypeAnqiJiguan:         itemtypes.ItemTypeAnqiJiguan,
	AdditionSysTypeFaBao:              itemtypes.ItemTypeFaBaoSuit,
	AdditionSysTypeXianTi:             itemtypes.ItemTypeXianTiLingYu,
	AdditionSysTypeLingYu:             itemtypes.ItemTypeLingyuEquip,
	AdditionSysTypeShenFa:             itemtypes.ItemTypeShenfaEquip,
	AdditionSysTypeShiHunFan:          itemtypes.ItemTypeShiHunFanEquip,
	AdditionSysTypeTianMoTi:           itemtypes.ItemTypeTianMoTiEquip,
	AdditionSysTypeLingTongWeapon:     itemtypes.ItemTypeLingTongWeaponEquip,
	AdditionSysTypeLingTongMount:      itemtypes.ItemTypeLingTongMountEquip,
	AdditionSysTypeLingTongWing:       itemtypes.ItemTypeLingTongWingEquip,
	AdditionSysTypeLingTongShenFa:     itemtypes.ItemTypeLingTongShenFaEquip,
	AdditionSysTypeLingTongLingYu:     itemtypes.ItemTypeLingTongLingYuEquip,
	AdditionSysTypeLingTongFaBao:      itemtypes.ItemTypeLingTongFaBaoEquip,
	AdditionSysTypeLingTongXianTi:     itemtypes.ItemTypeLingTongXianTiEquip,
	AdditionSysTypeShengHenQingLong:   itemtypes.ItemTypeShengHenEquipQingLong,
	AdditionSysTypeShengHenBaiHu:      itemtypes.ItemTypeShengHenEquipBaiHu,
	AdditionSysTypeShengHenZhuQue:     itemtypes.ItemTypeShengHenEquipZhuQue,
	AdditionSysTypeShengHenXuanWu:     itemtypes.ItemTypeShengHenEquipXuanWu,
	AdditionSysTypeLingTongOneEquip:   itemtypes.ItemTypeLingTongEquip,
	AdditionSysTypeLingTongTwoEquip:   itemtypes.ItemTypeLingTongEquip,
	AdditionSysTypeLingTongThreeEquip: itemtypes.ItemTypeLingTongEquip,
	AdditionSysTypeLingTongFourEquip:  itemtypes.ItemTypeLingTongEquip,
	AdditionSysTypeLingTongFiveEquip:  itemtypes.ItemTypeLingTongEquip,
	AdditionSysTypeLingTongSixEquip:   itemtypes.ItemTypeLingTongEquip,
	AdditionSysTypeLingTongSevenEquip: itemtypes.ItemTypeLingTongEquip,
	AdditionSysTypeLingTongEightEquip: itemtypes.ItemTypeLingTongEquip,
	AdditionSysTypeLingTongNineEquip:  itemtypes.ItemTypeLingTongEquip,
	AdditionSysTypeLingTongTenEquip:   itemtypes.ItemTypeLingTongEquip,
}

func (t AdditionSysType) ConvertToItemType() (itemtypes.ItemType, bool) {
	typ, ok := itemTypeMap[t]
	return typ, ok
}

// 转换 灵珠功能开启
var lingzhuOpenTypeMap = map[AdditionSysType]funcopentypes.FuncOpenType{
	AdditionSysTypeLingTongOneEquip:   funcopentypes.FuncOpenTypeLingTongEquipWuXingLingZhu,
	AdditionSysTypeLingTongTwoEquip:   funcopentypes.FuncOpenTypeLingTongEquipWuXingLingZhu,
	AdditionSysTypeLingTongThreeEquip: funcopentypes.FuncOpenTypeLingTongEquipWuXingLingZhu,
	AdditionSysTypeLingTongFourEquip:  funcopentypes.FuncOpenTypeLingTongEquipWuXingLingZhu,
	AdditionSysTypeLingTongFiveEquip:  funcopentypes.FuncOpenTypeLingTongEquipWuXingLingZhu,
	AdditionSysTypeLingTongSixEquip:   funcopentypes.FuncOpenTypeLingTongEquipWuXingLingZhu,
	AdditionSysTypeLingTongSevenEquip: funcopentypes.FuncOpenTypeLingTongEquipWuXingLingZhu,
	AdditionSysTypeLingTongEightEquip: funcopentypes.FuncOpenTypeLingTongEquipWuXingLingZhu,
	AdditionSysTypeLingTongNineEquip:  funcopentypes.FuncOpenTypeLingTongEquipWuXingLingZhu,
	AdditionSysTypeLingTongTenEquip:   funcopentypes.FuncOpenTypeLingTongEquipWuXingLingZhu,
}

func (t AdditionSysType) ConvertToLingZhuOpenType() (funcopentypes.FuncOpenType, bool) {
	typ, ok := lingzhuOpenTypeMap[t]
	return typ, ok
}

//转换为Template的AdditionSysType
var csAdditionSysTypeMap = map[AdditionSysType]AdditionSysType{
	AdditionSysTypeLingTongOneEquip:   AdditionSysTypeLingTongOneEquip,
	AdditionSysTypeLingTongTwoEquip:   AdditionSysTypeLingTongOneEquip,
	AdditionSysTypeLingTongThreeEquip: AdditionSysTypeLingTongOneEquip,
	AdditionSysTypeLingTongFourEquip:  AdditionSysTypeLingTongOneEquip,
	AdditionSysTypeLingTongFiveEquip:  AdditionSysTypeLingTongOneEquip,
	AdditionSysTypeLingTongSixEquip:   AdditionSysTypeLingTongOneEquip,
	AdditionSysTypeLingTongSevenEquip: AdditionSysTypeLingTongOneEquip,
	AdditionSysTypeLingTongEightEquip: AdditionSysTypeLingTongOneEquip,
	AdditionSysTypeLingTongNineEquip:  AdditionSysTypeLingTongOneEquip,
	AdditionSysTypeLingTongTenEquip:   AdditionSysTypeLingTongOneEquip,
}

func (t AdditionSysType) ConvertToTemplateAdditionSysType() AdditionSysType {
	typ, ok := csAdditionSysTypeMap[t]
	if !ok {
		typ = t
	}
	return typ
}

// 暂时用于灵珠系统(非通用)
const (
	MinAdditionSysTypeLingTongEquipType = AdditionSysTypeLingTongOneEquip
	MaxAdditionSysTypeLingTongEquipType = AdditionSysTypeLingTongTenEquip
)

// 灵童表id转换为AdditionSysType
var lingtongIdToAdditionSysTypeMap = map[int]AdditionSysType{
	1:  AdditionSysTypeLingTongOneEquip,
	2:  AdditionSysTypeLingTongTwoEquip,
	3:  AdditionSysTypeLingTongThreeEquip,
	4:  AdditionSysTypeLingTongFourEquip,
	5:  AdditionSysTypeLingTongFiveEquip,
	6:  AdditionSysTypeLingTongSixEquip,
	7:  AdditionSysTypeLingTongSevenEquip,
	8:  AdditionSysTypeLingTongEightEquip,
	9:  AdditionSysTypeLingTongNineEquip,
	10: AdditionSysTypeLingTongTenEquip,
}

var additionSysTypeToLingtongIdMap = map[AdditionSysType]int{
	AdditionSysTypeLingTongOneEquip:   1,
	AdditionSysTypeLingTongTwoEquip:   2,
	AdditionSysTypeLingTongThreeEquip: 3,
	AdditionSysTypeLingTongFourEquip:  4,
	AdditionSysTypeLingTongFiveEquip:  5,
	AdditionSysTypeLingTongSixEquip:   6,
	AdditionSysTypeLingTongSevenEquip: 7,
	AdditionSysTypeLingTongEightEquip: 8,
	AdditionSysTypeLingTongNineEquip:  9,
	AdditionSysTypeLingTongTenEquip:   10,
}

func ConvertLingTongIdToAdditionSysType(id int) (AdditionSysType, bool) {
	typ, ok := lingtongIdToAdditionSysTypeMap[id]
	return typ, ok
}

func (t AdditionSysType) ConvertAdditionSysTypeToLingTongId() (int, bool) {
	typ, ok := additionSysTypeToLingtongIdMap[t]
	return typ, ok
}
