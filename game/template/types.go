package template

import propertytypes "fgame/fgame/game/property/types"

// 系统装备接口
type AdditionSystemEquip interface {
	TemplateId() int
	HasCondition() bool
	GetReturnItemMap() map[int32]int32
	GetNeedItemMap() map[int32]int32
	GetTaozhuangTemplate() *SystemTaozhuangTemplate
	GetBattleProperty() map[propertytypes.BattlePropertyType]int64
	GetNextItemTemplate() *ItemTemplate
	// field
	GetTushiExp() int32
	GetHp() int32
	GetAttack() int32
	GetDefence() int32
	GetSuccessRate() int32
}
