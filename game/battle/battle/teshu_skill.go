package battle

import (
	battleventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

type TeShuSkillManager struct {
	bo        scene.BattleObject
	skillList []*scene.TeshuSkillObject
}

func (m *TeShuSkillManager) ResetTeShuSkills(oList []*scene.TeshuSkillObject) {
	m.skillList = oList
	gameevent.Emit(battleventtypes.EventTypeBattleObjectTeshuSkillReset, m.bo, nil)
}

func (m *TeShuSkillManager) GetTeShuSkills() []*scene.TeshuSkillObject {
	return m.skillList
}

func (m *TeShuSkillManager) GetTeShuSkill(skillId int32) *scene.TeshuSkillObject {
	for _, skillObj := range m.skillList {
		if skillObj.GetSkillId() == skillId {
			return skillObj
		}
	}
	return nil
}

func CreateTeShuSkillManager(bo scene.BattleObject, oList []*scene.TeshuSkillObject) *TeShuSkillManager {
	m := &TeShuSkillManager{
		bo:        bo,
		skillList: oList,
	}
	return m
}
