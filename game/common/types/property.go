package types

import (
	propertytypes "fgame/fgame/game/property/types"
)

type DanBeiPropertyType int32

const (
	//生命
	DanBeiPropertyTypeHp DanBeiPropertyType = 1 + iota
	//攻击
	DanBeiPropertyTypeAttack
	//防御
	DanBeiPropertyTypeDefend
	//暴击
	DanBeiPropertyTypeCrit
	//免暴
	DanBeiPropertyTypeTough
	//破格
	DanBeiPropertyTypeAbnormality
	//格挡
	DanBeiPropertyTypeBlock
)

func (t DanBeiPropertyType) Valid() bool {
	switch t {
	case DanBeiPropertyTypeHp,
		DanBeiPropertyTypeAttack,
		DanBeiPropertyTypeDefend,
		DanBeiPropertyTypeCrit,
		DanBeiPropertyTypeTough,
		DanBeiPropertyTypeAbnormality,
		DanBeiPropertyTypeBlock:
		return true
	}
	return false
}

func (t DanBeiPropertyType) GetPropertyType() propertytypes.BattlePropertyType {
	return danbeiPropertyTypeMap[t]
}

func (t DanBeiPropertyType) Mask() int32 {
	val := uint(t) - 1
	return 1 << val
}

var danbeiPropertyTypeMap = map[DanBeiPropertyType]propertytypes.BattlePropertyType{
	DanBeiPropertyTypeHp:          propertytypes.BattlePropertyTypeMaxHP,
	DanBeiPropertyTypeAttack:      propertytypes.BattlePropertyTypeAttack,
	DanBeiPropertyTypeDefend:      propertytypes.BattlePropertyTypeDefend,
	DanBeiPropertyTypeCrit:        propertytypes.BattlePropertyTypeCrit,
	DanBeiPropertyTypeTough:       propertytypes.BattlePropertyTypeTough,
	DanBeiPropertyTypeAbnormality: propertytypes.BattlePropertyTypeAbnormality,
	DanBeiPropertyTypeBlock:       propertytypes.BattlePropertyTypeBlock,
}
