package types

//点击大类
type ClickType int32

const (
	//装备类型
	ClickTypeEquip ClickType = 1 + iota
	//技能类型
	ClickTypeSkill
	//帝魂类型
	ClickTypeSoul
	//仙盟
	ClickTypeAlliance
	//升阶系统
	ClickTypeAdvanced
	//升级系统
	ClickTypeUpgradeSys
	//元神金装类型
	ClickTypeGoldEquipment
	//灵童
	ClickTypeLingTong
)

func (c ClickType) Valid() bool {
	switch c {
	case ClickTypeEquip,
		ClickTypeSkill,
		ClickTypeSoul,
		ClickTypeAlliance,
		ClickTypeAdvanced,
		ClickTypeUpgradeSys,
		ClickTypeGoldEquipment,
		ClickTypeLingTong:
		return true
	}
	return false
}

type ClickSubType interface {
	SubType() int32
	Valid() bool
}

type ClickSubTypeFactory interface {
	GetClickSubType(subType int32) ClickSubType
}

type ClickSubTypeFactoryFunc func(subType int32) ClickSubType

func (c ClickSubTypeFactoryFunc) GetClickSubType(subType int32) ClickSubType {
	return c(subType)
}

//点击装备子类型
type ClickSubTypeEquip int32

const (
	//装备强化 || 装备一键强化
	ClickSubTypeEquipStrength ClickSubTypeEquip = 1 + iota
	//装备升阶
	ClickSubTypeEquipUpgrade
	//装备升星
	ClickSubTypeEquipUpgradeStar
)

func (c ClickSubTypeEquip) SubType() int32 {
	return int32(c)
}

func (c ClickSubTypeEquip) Valid() bool {
	switch c {
	case ClickSubTypeEquipStrength,
		ClickSubTypeEquipUpgrade,
		ClickSubTypeEquipUpgradeStar:
		return true
	}
	return false
}

func GetClickSubTypeEquip(subType int32) ClickSubType {
	return ClickSubTypeEquip(subType)
}

//点击技能子类型
type ClickSubTypeSkill int32

const (
	//一键升级 || 升级职业技能
	ClickSubTypeSkillUpgrade ClickSubTypeSkill = 1 + iota
)

func (c ClickSubTypeSkill) SubType() int32 {
	return int32(c)
}

func (c ClickSubTypeSkill) Valid() bool {
	switch c {
	case ClickSubTypeSkillUpgrade:
		return true
	}
	return false
}

func GetClickSubTypeSkill(subType int32) ClickSubType {
	return ClickSubTypeSkill(subType)
}

//点击帝魂子类型
type ClickSubTypeSoul int32

const (
	//强化帝魂
	ClickSubTypeSoulStrengthen ClickSubTypeSoul = 1 + iota
	//魂技升级
	ClickSubTypeSoulUpgrade
)

func (c ClickSubTypeSoul) SubType() int32 {
	return int32(c)
}

func (c ClickSubTypeSoul) Valid() bool {
	switch c {
	case ClickSubTypeSoulStrengthen,
		ClickSubTypeSoulUpgrade:
		return true
	}
	return false
}

func GetClickSubTypeSoul(subType int32) ClickSubType {
	return ClickSubTypeSoul(subType)
}

//点击仙盟子类型
type ClickSubTypeAlliance int32

const (
	//申请加入仙盟
	ClickSubTypeAllianceApplyJoin ClickSubTypeAlliance = 1 + iota
)

func (c ClickSubTypeAlliance) SubType() int32 {
	return int32(c)
}

func (c ClickSubTypeAlliance) Valid() bool {
	switch c {
	case ClickSubTypeAllianceApplyJoin:
		return true
	}
	return false
}

func GetClickSubTypeAlliance(subType int32) ClickSubType {
	return ClickSubTypeAlliance(subType)
}

//点击升阶子类型
type ClickSubTypeAdvanced int32

const (
	//坐骑升阶
	ClickSubTypeAdvancedMount ClickSubTypeAdvanced = iota
	//战翼升阶
	ClickSubTypeAdvancedWing
	//身法升阶
	ClickSubTypeAdvancedShenFa
	//法宝升阶
	ClickSubTypeAdvancedFaBao
	//仙体升阶
	ClickSubTypeAdvancedXianTi
	//暗器升阶
	ClickSubTypeAdvancedAnQi
)

func (c ClickSubTypeAdvanced) SubType() int32 {
	return int32(c)
}

func (c ClickSubTypeAdvanced) Valid() bool {
	switch c {
	case ClickSubTypeAdvancedMount,
		ClickSubTypeAdvancedWing,
		ClickSubTypeAdvancedShenFa,
		ClickSubTypeAdvancedFaBao,
		ClickSubTypeAdvancedXianTi,
		ClickSubTypeAdvancedAnQi:
		return true
	}
	return false
}

func GetClickSubTypeAdvanced(subType int32) ClickSubType {
	return ClickSubTypeAdvanced(subType)
}

//点击升级系统
type ClickSubTypeUpgradeSys int32

const (
	//坐骑系统
	ClickSubTypeUpgradeSysMount ClickSubTypeUpgradeSys = 1 + iota
	//战翼系统
	ClickSubTypeUpgradeSysWing
	//暗器系统
	ClickSubTypeUpgradeSysAnQi
)

func (c ClickSubTypeUpgradeSys) SubType() int32 {
	return int32(c)
}

func (c ClickSubTypeUpgradeSys) Valid() bool {
	switch c {
	case ClickSubTypeUpgradeSysMount,
		ClickSubTypeUpgradeSysWing,
		ClickSubTypeUpgradeSysAnQi:
		return true
	}
	return false
}

func GetClickSubTypeUpgradeSys(subType int32) ClickSubType {
	return ClickSubTypeUpgradeSys(subType)
}

//点击强化元神金装
type ClickSubTypeGoldEquipment int32

const (
	//武器
	ClickSubTypeGoldEquipmentWeapon ClickSubTypeGoldEquipment = iota
	//衣服
	ClickSubTypeGoldEquipmentClothes
	//头盔
	ClickSubTypeGoldEquipmentHelmet
	//战靴
	ClickSubTypeGoldEquipmentCaliga
	//腰带
	ClickSubTypeGoldEquipmentBelt
	//护手
	ClickSubTypeGoldEquipmentHand
	//玉坠
	ClickSubTypeGoldEquipmentJade
	//项链
	ClickSubTypeGoldEquipmentNecklace
)

func (c ClickSubTypeGoldEquipment) SubType() int32 {
	return int32(c)
}

func (c ClickSubTypeGoldEquipment) Valid() bool {
	switch c {
	case ClickSubTypeGoldEquipmentWeapon,
		ClickSubTypeGoldEquipmentClothes,
		ClickSubTypeGoldEquipmentHelmet,
		ClickSubTypeGoldEquipmentCaliga,
		ClickSubTypeGoldEquipmentBelt,
		ClickSubTypeGoldEquipmentHand,
		ClickSubTypeGoldEquipmentJade,
		ClickSubTypeGoldEquipmentNecklace:
		return true
	}
	return false
}

func GetClickSubTypeGoldEquipment(subType int32) ClickSubType {
	return ClickSubTypeGoldEquipment(subType)
}

//点击灵童
type ClickSubTypeLingTong int32

const (
	//灵童升级
	ClickSubTypeLingTongShengJi ClickSubTypeLingTong = iota
)

func (c ClickSubTypeLingTong) SubType() int32 {
	return int32(c)
}

func (c ClickSubTypeLingTong) Valid() bool {
	switch c {
	case ClickSubTypeLingTongShengJi:
		return true
	}
	return false
}

func GetClickSubTypeLingTong(subType int32) ClickSubType {
	return ClickSubTypeLingTong(subType)
}

var (
	clickSubTypeFactoryMap = make(map[ClickType]ClickSubTypeFactory)
)

func GetClickSubType(typ ClickType, subType int32) ClickSubType {
	factory, ok := clickSubTypeFactoryMap[typ]
	if !ok {
		return nil
	}
	return factory.GetClickSubType(subType)
}

func init() {
	clickSubTypeFactoryMap[ClickTypeEquip] = ClickSubTypeFactoryFunc(GetClickSubTypeEquip)
	clickSubTypeFactoryMap[ClickTypeSkill] = ClickSubTypeFactoryFunc(GetClickSubTypeSkill)

	clickSubTypeFactoryMap[ClickTypeSoul] = ClickSubTypeFactoryFunc(GetClickSubTypeSoul)
	clickSubTypeFactoryMap[ClickTypeAlliance] = ClickSubTypeFactoryFunc(GetClickSubTypeAlliance)

	clickSubTypeFactoryMap[ClickTypeAdvanced] = ClickSubTypeFactoryFunc(GetClickSubTypeAdvanced)
	clickSubTypeFactoryMap[ClickTypeUpgradeSys] = ClickSubTypeFactoryFunc(GetClickSubTypeUpgradeSys)

	clickSubTypeFactoryMap[ClickTypeGoldEquipment] = ClickSubTypeFactoryFunc(GetClickSubTypeGoldEquipment)
	clickSubTypeFactoryMap[ClickTypeLingTong] = ClickSubTypeFactoryFunc(GetClickSubTypeLingTong)
}
