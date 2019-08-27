package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	skilltemplate "fgame/fgame/game/skill/template"
)

type SkillActionManager struct {
	bo       scene.BattleObject
	skillMap map[int32]int64
}

func (m *SkillActionManager) AddSkillAction(skillId int32) {
	now := global.GetGame().GetTimeService().Now()
	m.skillMap[skillId] = now
}

func (m *SkillActionManager) ClearAllSkillAction() {
	for skillId, _ := range m.skillMap {
		delete(m.skillMap, skillId)
	}
}

func (m *SkillActionManager) Heartbeat() {

	now := global.GetGame().GetTimeService().Now()
	for skillId, startTime := range m.skillMap {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
		elapse := now - startTime
		if elapse >= int64(skillTemplate.GetDelayTime()) {
			//发送事件
			gameevent.Emit(battleeventtypes.EventTypeBattleObjectSkillTrigger, m.bo, skillId)
			delete(m.skillMap, skillId)
		}
	}
}

func CreateSkillActionManager(bo scene.BattleObject) *SkillActionManager {
	m := &SkillActionManager{}
	m.bo = bo
	m.skillMap = make(map[int32]int64)
	return m
}
