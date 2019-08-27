package types

import (
	propertytypes "fgame/fgame/game/property/types"
)

type MingGeType int32

const (
	//普通命格
	MingGeTypeNormal MingGeType = iota
	//超级命格
	MingGeTypeSuper
)

func (t MingGeType) Valid() bool {
	switch t {
	case MingGeTypeNormal,
		MingGeTypeSuper:
		return true
	}
	return false
}

type MingGeSubType interface {
	SubType() int32
	Valid() bool
}

type MingGeSubTypeFactory interface {
	CreateMingGeSubType(subType int32) MingGeSubType
}

type MingGeSubTypeFactoryFunc func(subType int32) MingGeSubType

func (t MingGeSubTypeFactoryFunc) CreateMingGeSubType(subType int32) MingGeSubType {
	return t(subType)
}

// 命格子类型
type MingGeAllSubType int32

const (
	//贪狼
	MingGeAllSubTypeTanLang MingGeAllSubType = 1 + iota
	//破军
	MingGeAllSubTypePoJun
	//七杀
	MingGeAllSubTypeQiSha
	//紫薇
	MingGeAllSubTypeZiWei
)

func (t MingGeAllSubType) SubType() int32 {
	return int32(t)
}

func (t MingGeAllSubType) Valid() bool {
	switch t {
	case MingGeAllSubTypeTanLang,
		MingGeAllSubTypePoJun,
		MingGeAllSubTypeQiSha,
		MingGeAllSubTypeZiWei:
		return true
	}
	return false
}

const (
	MingGeAllSubTypeMin = MingGeAllSubTypeTanLang
	MingGeAllSubTypeMax = MingGeAllSubTypeZiWei
)

func CreateMingGeAllSubType(subType int32) MingGeSubType {
	return MingGeAllSubType(subType)
}

var (
	mingGeSubTypeFactoryMap = make(map[MingGeType]MingGeSubTypeFactory)
)

func CreateMingGeSubType(typ MingGeType, subType int32) MingGeSubType {
	factory, ok := mingGeSubTypeFactoryMap[typ]
	if !ok {
		panic("mingge:CreateMingGeSubType 应该是ok的")
	}
	return factory.CreateMingGeSubType(subType)
}

func init() {
	mingGeSubTypeFactoryMap[MingGeTypeNormal] = MingGeSubTypeFactoryFunc(CreateMingGeAllSubType)
	mingGeSubTypeFactoryMap[MingGeTypeSuper] = MingGeSubTypeFactoryFunc(CreateMingGeAllSubType)
}

//命宫类型
type MingGongType int32

const (
	//命宫
	MingGongTypeGong MingGongType = iota
	//财帛宫
	MingGongTypeCaiJin
	//兄弟宫
	MingGongTypeXiongDi
	//田宅宫
	MingGongTypeTianZhai
	//子女宫
	MingGongTypeZiNv
	//奴仆宫
	MingGongTypeNuPu
	//夫妻宫
	MingGongTypeFuQi
	//疾恶宫
	MingGongTypeJiE
	//迁移宫
	MingGongTypeQianYi
	//官禄宫
	MingGongTypeGuanLu
	//福德宫
	MingGongTypeFuDe
	//相貌宫
	MingGongTypeXiangMao
)

func (t MingGongType) Valid() bool {
	switch t {
	case MingGongTypeGong,
		MingGongTypeCaiJin,
		MingGongTypeXiongDi,
		MingGongTypeTianZhai,
		MingGongTypeZiNv,
		MingGongTypeNuPu,
		MingGongTypeFuQi,
		MingGongTypeJiE,
		MingGongTypeQianYi,
		MingGongTypeGuanLu,
		MingGongTypeFuDe,
		MingGongTypeXiangMao:
		return true
	}
	return false
}

const (
	MingGongTypeGongMin = MingGongTypeGong
	MingGongTypeGongMax = MingGongTypeXiangMao
)

type MingGongSubType interface {
	SubType() int32
	Valid() bool
}

type MingGongSubTypeFactory interface {
	CreateMingGongSubType(subType int32) MingGongSubType
}

type MingGongSubTypeFactoryFunc func(subType int32) MingGongSubType

func (t MingGongSubTypeFactoryFunc) CreateMingGongSubType(subType int32) MingGongSubType {
	return t(subType)
}

// 命宫子类型
type MingGongAllSubType int32

const (
	//部位1
	MingGongAllSubTypeOne MingGongAllSubType = 1 + iota
	//部位2
	MingGongAllSubTypeTwo
	//部位3
	MingGongAllSubTypeThree
	//部位4
	MingGongAllSubTypeFours
	//部位5
	MingGongAllSubTypeFive
	//部位6
	MingGongAllSubTypeSix
)

const (
	MingGongAllSubTypeMin = MingGongAllSubTypeOne
	MingGongAllSubTypeMax = MingGongAllSubTypeSix
)

func (t MingGongAllSubType) SubType() int32 {
	return int32(t)
}

func (t MingGongAllSubType) Valid() bool {
	switch t {
	case MingGongAllSubTypeOne,
		MingGongAllSubTypeTwo,
		MingGongAllSubTypeThree,
		MingGongAllSubTypeFours,
		MingGongAllSubTypeFive,
		MingGongAllSubTypeSix:
		return true
	}
	return false
}

func CreateMingGongAllSubType(subType int32) MingGongSubType {
	return MingGongAllSubType(subType)
}

var (
	mingGongSubTypeFactoryMap = make(map[MingGongType]MingGongSubTypeFactory)
)

func CreateMingGongSubType(typ MingGongType, subType int32) MingGongSubType {
	factory, ok := mingGongSubTypeFactoryMap[typ]
	if !ok {
		panic("mingge:CreateMingGongSubType 应该是ok的")
	}
	return factory.CreateMingGongSubType(subType)
}

func init() {
	mingGongSubTypeFactoryMap[MingGongTypeGong] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
	mingGongSubTypeFactoryMap[MingGongTypeCaiJin] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
	mingGongSubTypeFactoryMap[MingGongTypeXiongDi] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
	mingGongSubTypeFactoryMap[MingGongTypeTianZhai] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
	mingGongSubTypeFactoryMap[MingGongTypeZiNv] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
	mingGongSubTypeFactoryMap[MingGongTypeNuPu] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
	mingGongSubTypeFactoryMap[MingGongTypeFuQi] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
	mingGongSubTypeFactoryMap[MingGongTypeJiE] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
	mingGongSubTypeFactoryMap[MingGongTypeQianYi] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
	mingGongSubTypeFactoryMap[MingGongTypeGuanLu] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
	mingGongSubTypeFactoryMap[MingGongTypeFuDe] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
	mingGongSubTypeFactoryMap[MingGongTypeXiangMao] = MingGongSubTypeFactoryFunc(CreateMingGongAllSubType)
}

type MingGePropertyType int32

const (
	//生命
	MingGePropertyTypeLife MingGePropertyType = 1 + iota
	//攻击
	MingGePropertyTypeAttack
	//防御
	MingGePropertyTypeDefend
	//暴击
	MingGePropertyTypeCrit
	//免暴
	MingGePropertyTypeTough
	//破格
	MingGePropertyTypeAbnormality
	//格挡
	MingGePropertyTypeBlock
)

func (t MingGePropertyType) Valid() bool {
	switch t {
	case MingGePropertyTypeLife,
		MingGePropertyTypeAttack,
		MingGePropertyTypeDefend,
		MingGePropertyTypeCrit,
		MingGePropertyTypeTough,
		MingGePropertyTypeAbnormality,
		MingGePropertyTypeBlock:
		return true
	}
	return false
}

func (t MingGePropertyType) GetPropertyType() propertytypes.BattlePropertyType {
	return mingGePropertyTypeMap[t]
}

func (t MingGePropertyType) Mask() int32 {
	val := uint(t) - 1
	return 1 << val
}

var mingGePropertyTypeMap = map[MingGePropertyType]propertytypes.BattlePropertyType{
	MingGePropertyTypeLife:        propertytypes.BattlePropertyTypeMaxHP,
	MingGePropertyTypeAttack:      propertytypes.BattlePropertyTypeAttack,
	MingGePropertyTypeDefend:      propertytypes.BattlePropertyTypeDefend,
	MingGePropertyTypeCrit:        propertytypes.BattlePropertyTypeCrit,
	MingGePropertyTypeTough:       propertytypes.BattlePropertyTypeTough,
	MingGePropertyTypeAbnormality: propertytypes.BattlePropertyTypeAbnormality,
	MingGePropertyTypeBlock:       propertytypes.BattlePropertyTypeBlock,
}

//命理槽位
type MingLiSlotType int32

const (
	MingLiSlotTypeOne MingLiSlotType = 1 + iota
	MingLiSlotTypeTwo
	MingLiSlotTypeThree
)

func (t MingLiSlotType) Vaild() bool {
	switch t {
	case MingLiSlotTypeOne,
		MingLiSlotTypeTwo,
		MingLiSlotTypeThree:
		return true
	}
	return false
}

const (
	MingLiSlotTypeMin = MingLiSlotTypeOne
	MingLiSlotTypeMax = MingLiSlotTypeThree
)

//命格镶嵌槽位
type MingGeSlotType int32

const (
	MingGeSlotTypeOne MingGeSlotType = 1 + iota
	MingGeSlotTypeTwo
	MingGeSlotTypeThree
	MingGeSlotTypeFour
)

func (t MingGeSlotType) Vaild() bool {
	switch t {
	case MingGeSlotTypeOne,
		MingGeSlotTypeTwo,
		MingGeSlotTypeThree,
		MingGeSlotTypeFour:
		return true
	}
	return false
}
