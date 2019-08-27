package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	skilllogic "fgame/fgame/game/skill/logic"

	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
	shenfatemplate "fgame/fgame/game/shenfa/template"
)

//玩家身法进阶
func shenfaAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	advanceId := data.(int32)
	newSkillId := shenfatemplate.GetShenfaTemplateService().GetShenfaSkill(advanceId)
	oldSkillId := shenfatemplate.GetShenfaTemplateService().GetShenfaSkill(advanceId - 1)
	if newSkillId != oldSkillId {
		err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	}

	return
}

func init() {
	gameevent.AddEventListener(shenfaeventtypes.EventTypeShenfaAdvanced, event.EventListenerFunc(shenfaAdvanced))
}
