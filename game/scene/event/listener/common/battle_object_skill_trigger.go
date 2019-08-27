package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
)

//玩家复活
func battleObjectSkillTrigger(target event.EventTarget, data event.EventData) (err error) {

	skillId := data.(int32)
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)

	switch bo := target.(type) {
	case scene.LingTong:
		if skillTemplate.GetSkillSecondType() == skilltypes.SkillSecondTypePositive {
			scenelogic.CalculateLingTongAttack(bo, bo.GetPosition(), bo.GetAngle(), skillTemplate, true)
		} else {
			scenelogic.CalculateLingTongAttack(bo, bo.GetPosition(), bo.GetAngle(), skillTemplate, false)
		}
		break
	case scene.BattleObject:
		if skillTemplate.GetSkillSecondType() == skilltypes.SkillSecondTypePositive {
			scenelogic.CalculateAttack(bo, bo.GetPosition(), bo.GetAngle(), skillTemplate, true)
		} else {
			scenelogic.CalculateAttack(bo, bo.GetPosition(), bo.GetAngle(), skillTemplate, false)
		}
		break
	}

	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattleObjectSkillTrigger, event.EventListenerFunc(battleObjectSkillTrigger))
}
