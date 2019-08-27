package types

//强化类型
type EquipmentStrengthenType int32

const (
	//强化
	EquipmentStrengthenTypeUpgrade EquipmentStrengthenType = iota + 1
	//升星
	EquipmentStrengthenTypeStar
)

func (est EquipmentStrengthenType) Valid() bool {
	switch est {
	case EquipmentStrengthenTypeStar,
		EquipmentStrengthenTypeUpgrade:
		return true
	}
	return false
}

type EquipmentStrengthenResultType int32

const (
	//强化成功
	EquipmentStrengthenResultTypeSuccess EquipmentStrengthenResultType = iota
	//强化失败
	EquipmentStrengthenResultTypeFailed
	//强化回退
	EquipmentStrengthenResultTypeBack
)
