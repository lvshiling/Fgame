package player

import (
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	lingtongplayerproperty "fgame/fgame/game/lingtong/player/property"
	propertytypes "fgame/fgame/game/property/types"
)

//是否属性变化过
func (m *PlayerLingTongDataManager) IsPropertyChanged() bool {
	return m.battlePropertyGroup.IsChanged()
}

//重置改变标记位
func (m *PlayerLingTongDataManager) resetChanged() {
	m.battlePropertyGroup.ResetChanged()
}

func (m *PlayerLingTongDataManager) GetChangedBattlePropertiesAndReset() (battleChanged map[int32]int64) {
	battleChanged = m.battlePropertyGroup.GetChangedTypes()
	m.battlePropertyGroup.ResetChanged()
	return
}

func (m *PlayerLingTongDataManager) GetAllSystemBattleProperties() map[int32]int64 {
	properties := make(map[int32]int64)
	for typ := propertytypes.MinBattlePropertyType; typ <= propertytypes.MaxBattlePropertyType; typ++ {
		val := m.battlePropertyGroup.Get(typ)
		properties[int32(typ)] = val
	}
	return properties
}

//获取战斗属性
func (m *PlayerLingTongDataManager) GetBattleProperty(battlePropertyType propertytypes.BattlePropertyType) int64 {
	return m.battlePropertyGroup.Get(battlePropertyType)
}

//更新战斗属性
func (m *PlayerLingTongDataManager) UpdateBattleProperty(mask uint64) {
	//各个系统
	for _, effType := range lingtongplayerproperty.GetAllLingTongPropertyEffectors() {
		if effType.Mask()&mask != 0 {
			pef := lingtongplayerproperty.GetLingTongPropertyEffector(effType)
			if pef == nil {
				panic("never reach here")
			}
			//获取相对应属性
			p := m.battlePropertyGroup.GetPropertySegment(effType)
			//重置属性
			p.Clear()
			//属性作用
			pef(m.p, p)
		}
	}
	m.battlePropertyGroup.UpdateProperty()

	gameevent.Emit(lingtongeventtypes.EventTypeLingTongSystemPropertyChanged, m.p, nil)

}
