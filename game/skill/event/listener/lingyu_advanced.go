package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	skilllogic "fgame/fgame/game/skill/logic"

	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
	lingyutemplate "fgame/fgame/game/lingyu/template"
)

//玩家领域进阶
func lingyuAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}
	advanceId := data.(int32)
	newSkillId := lingyutemplate.GetLingyuTemplateService().GetLingyuSkill(advanceId)
	oldSkillId := lingyutemplate.GetLingyuTemplateService().GetLingyuSkill(advanceId - 1)
	if newSkillId != oldSkillId {
		err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	}

	return
}

func init() {
	gameevent.AddEventListener(lingyueventtypes.EventTypeLingyuAdvanced, event.EventListenerFunc(lingyuAdvanced))
}
