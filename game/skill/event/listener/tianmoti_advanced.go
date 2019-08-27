package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	skilllogic "fgame/fgame/game/skill/logic"
	tianmoeventtypes "fgame/fgame/game/tianmo/event/types"
	tianmotemplate "fgame/fgame/game/tianmo/template"
)

//玩家天魔体进阶
func tianMoTiAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	advanceId := data.(int32)
	newSkillId := tianmotemplate.GetTianMoTemplateService().GetTianMoSkill(advanceId)
	oldSkillId := tianmotemplate.GetTianMoTemplateService().GetTianMoSkill(advanceId - 1)
	if newSkillId != oldSkillId {
		err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	}

	return
}

func init() {
	gameevent.AddEventListener(tianmoeventtypes.EventTypeTianMoAdvanced, event.EventListenerFunc(tianMoTiAdvanced))
}
