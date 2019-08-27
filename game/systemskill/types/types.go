package types

import (
	additionsystypes "fgame/fgame/game/additionsys/types"
	skilltypes "fgame/fgame/game/skill/types"
)

type SystemSkillType int32

const (
	//坐骑技能
	SystemSkillTypeMount SystemSkillType = 1 + iota
	//战翼技能
	SystemSkillTypeWing SystemSkillType = 2
	//暗器技能
	SystemSkillTypeAnQi SystemSkillType = 3
	//法宝技能
	SystemSkillTypeFaBao SystemSkillType = 4
	//仙体技能
	SystemSkillTypeXianTi SystemSkillType = 5
	//领域技能
	SystemSkillTypeLingYu SystemSkillType = 6
	//身法技能
	SystemSkillTypeShenFa SystemSkillType = 7
	//噬魂幡技能
	SystemSkillTypeShiHunFan SystemSkillType = 8
	//天魔体技能
	SystemSkillTypeTianMo SystemSkillType = 9
	//灵兵
	SystemSkillTypeLingTongWeapon SystemSkillType = 101
	//灵骑
	SystemSkillTypeLingTongMount SystemSkillType = 102
	//灵翼
	SystemSkillTypeLingTongWing SystemSkillType = 103
	//灵身
	SystemSkillTypeLingTongShenFa SystemSkillType = 104
	//灵域
	SystemSkillTypeLingTongLingYu SystemSkillType = 105
	//灵宝
	SystemSkillTypeLingTongFaBao SystemSkillType = 106
	//灵体
	SystemSkillTypeLingTongXianTi SystemSkillType = 107
	//圣痕技能-青龙
	SystemSkillTypeShengHenQingLong SystemSkillType = 201
	//圣痕技能-白虎
	SystemSkillTypeShengHenBaiHu SystemSkillType = 202
	//圣痕技能-朱雀
	SystemSkillTypeShengHenZhuQue SystemSkillType = 203
	//圣痕技能-玄武
	SystemSkillTypeShengHenXuanWu SystemSkillType = 204
)

func (s SystemSkillType) Valid() bool {
	switch s {
	case SystemSkillTypeMount,
		SystemSkillTypeWing,
		SystemSkillTypeAnQi,
		SystemSkillTypeFaBao,
		SystemSkillTypeXianTi,
		SystemSkillTypeLingYu,
		SystemSkillTypeShenFa,
		SystemSkillTypeShiHunFan,
		SystemSkillTypeTianMo,
		SystemSkillTypeLingTongWeapon,
		SystemSkillTypeLingTongMount,
		SystemSkillTypeLingTongWing,
		SystemSkillTypeLingTongShenFa,
		SystemSkillTypeLingTongLingYu,
		SystemSkillTypeLingTongFaBao,
		SystemSkillTypeLingTongXianTi,
		SystemSkillTypeShengHenQingLong,
		SystemSkillTypeShengHenBaiHu,
		SystemSkillTypeShengHenZhuQue,
		SystemSkillTypeShengHenXuanWu:
		return true
	default:
		return false
	}
}

func (s SystemSkillType) String() string {
	return systemSkillTypeStringMap[s]
}

func (s SystemSkillType) GetSkillFirstType() skilltypes.SkillFirstType {
	return systemSkillTypeMap[s]
}

func (s SystemSkillType) ConvertToAdditionSysType() (additionsystypes.AdditionSysType, bool) {
	typ, ok := additionsystypesMap[s]
	return typ, ok
}

var systemSkillTypeMap = map[SystemSkillType]skilltypes.SkillFirstType{
	SystemSkillTypeMount:            skilltypes.SkillFirstTypeMountEquip,
	SystemSkillTypeWing:             skilltypes.SkillFirstTypeWingStone,
	SystemSkillTypeAnQi:             skilltypes.SkillFirstTypeAnQiJiGuan,
	SystemSkillTypeFaBao:            skilltypes.SkillFirstTypeFaBao,
	SystemSkillTypeXianTi:           skilltypes.SkillFirstTypeXianTi,
	SystemSkillTypeLingYu:           skilltypes.SkillFirstTypeLingYuSystemSkill,
	SystemSkillTypeShenFa:           skilltypes.SkillFirstTypeShenFaSystemSkill,
	SystemSkillTypeShiHunFan:        skilltypes.SkillFirstTypeShiHunFanSystemSkill,
	SystemSkillTypeTianMo:           skilltypes.SkillFirstTypeTianMoSystemSkill,
	SystemSkillTypeShengHenQingLong: skilltypes.SkillFirstTypeShengHenQingLong,
	SystemSkillTypeShengHenBaiHu:    skilltypes.SkillFirstTypeShengHenBaiHu,
	SystemSkillTypeShengHenZhuQue:   skilltypes.SkillFirstTypeShengHenZhuQue,
	SystemSkillTypeShengHenXuanWu:   skilltypes.SkillFirstTypeShengHenXuanWu,
	SystemSkillTypeLingTongWeapon:   skilltypes.SkillFirstTypeLingTongWeaponSystemSkill,
	SystemSkillTypeLingTongMount:    skilltypes.SkillFirstTypeLingTongMountSystemSkill,
	SystemSkillTypeLingTongWing:     skilltypes.SkillFirstTypeLingTongWingSystemSkill,
	SystemSkillTypeLingTongShenFa:   skilltypes.SkillFirstTypeLingTongShenFaSystemSkill,
	SystemSkillTypeLingTongLingYu:   skilltypes.SkillFirstTypeLingTongLingYuSystemSkill,
	SystemSkillTypeLingTongFaBao:    skilltypes.SkillFirstTypeLingTongFaBaoSystemSkill,
	SystemSkillTypeLingTongXianTi:   skilltypes.SkillFirstTypeLingTongXianTiSystemSkill,
}

var (
	systemSkillTypeStringMap = map[SystemSkillType]string{
		SystemSkillTypeMount:            "坐骑技能",
		SystemSkillTypeWing:             "战翼技能",
		SystemSkillTypeAnQi:             "暗器技能",
		SystemSkillTypeFaBao:            "法宝技能",
		SystemSkillTypeXianTi:           "仙体技能",
		SystemSkillTypeLingYu:           "领域技能",
		SystemSkillTypeShenFa:           "身法技能",
		SystemSkillTypeShiHunFan:        "噬魂幡技能",
		SystemSkillTypeTianMo:           "天魔技能",
		SystemSkillTypeLingTongWeapon:   "灵兵技能",
		SystemSkillTypeLingTongMount:    "灵骑技能",
		SystemSkillTypeLingTongWing:     "灵翼技能",
		SystemSkillTypeLingTongShenFa:   "灵身技能",
		SystemSkillTypeLingTongLingYu:   "灵域技能",
		SystemSkillTypeLingTongFaBao:    "灵宝技能",
		SystemSkillTypeLingTongXianTi:   "灵体技能",
		SystemSkillTypeShengHenQingLong: "圣痕技能-青龙",
		SystemSkillTypeShengHenBaiHu:    "圣痕技能-白虎",
		SystemSkillTypeShengHenZhuQue:   "圣痕技能-朱雀",
		SystemSkillTypeShengHenXuanWu:   "圣痕技能-玄武",
	}
)

var additionsystypesMap = map[SystemSkillType]additionsystypes.AdditionSysType{
	SystemSkillTypeShengHenQingLong: additionsystypes.AdditionSysTypeShengHenQingLong,
	SystemSkillTypeShengHenBaiHu:    additionsystypes.AdditionSysTypeShengHenBaiHu,
	SystemSkillTypeShengHenZhuQue:   additionsystypes.AdditionSysTypeShengHenZhuQue,
	SystemSkillTypeShengHenXuanWu:   additionsystypes.AdditionSysTypeShengHenXuanWu,
}

type SystemSkillSubType int32

const (
	SystemSkillSubTypeOne SystemSkillSubType = 1 + iota
	SystemSkillSubTypeTwo
	SystemSkillSubTypeThree
	SystemSkillSubTypeFour
	SystemSkillSubTypeFive
	SystemSkillSubTypeSix
)

func (s SystemSkillSubType) Valid() bool {
	switch s {
	case SystemSkillSubTypeOne,
		SystemSkillSubTypeTwo,
		SystemSkillSubTypeThree,
		SystemSkillSubTypeFour,
		SystemSkillSubTypeFive,
		SystemSkillSubTypeSix:
		return true
	default:
		return false
	}
}
