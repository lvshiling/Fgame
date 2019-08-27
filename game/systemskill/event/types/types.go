package types

import (
	sysskilltypes "fgame/fgame/game/systemskill/types"
)

type SystemSkillEventType string

const (
	//系统技能激活
	EventTypeSystemSkillActive SystemSkillEventType = "systemSkillActive"
	//系统技能升级
	EventTypeSystemSkillUpgrade SystemSkillEventType = "systemSkillUpgrade"
)

type SystemSkillEventData struct {
	typ     sysskilltypes.SystemSkillType
	subType sysskilltypes.SystemSkillSubType
}

func (d *SystemSkillEventData) GetType() sysskilltypes.SystemSkillType {
	return d.typ
}

func (d *SystemSkillEventData) GetSubType() sysskilltypes.SystemSkillSubType {
	return d.subType
}

func CreateSystemSkillEventData(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType) *SystemSkillEventData {
	tted := &SystemSkillEventData{
		typ:     typ,
		subType: subType,
	}
	return tted
}
