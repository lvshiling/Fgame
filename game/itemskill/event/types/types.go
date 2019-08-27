package types

import (
	itemskilltypes "fgame/fgame/game/itemskill/types"
)

type ItemSkillEventType string

const (
	//物品技能激活
	EventTypeItemSkillActive ItemSkillEventType = "itemSkillActive"
	//物品技能升级
	EventTypeItemSkillUpgrade ItemSkillEventType = "itemSkillUpgrade"
)

type ItemSkillUpgradeEventData struct {
	typ    itemskilltypes.ItemSkillType
	oldLev int32
}

func (d *ItemSkillUpgradeEventData) GetType() itemskilltypes.ItemSkillType {
	return d.typ
}

func (d *ItemSkillUpgradeEventData) GetOldLev() int32 {
	return d.oldLev
}

func CreateXianFuChallengeEventData(typ itemskilltypes.ItemSkillType, oldLev int32) *ItemSkillUpgradeEventData {
	return &ItemSkillUpgradeEventData{
		typ:    typ,
		oldLev: oldLev,
	}
}
