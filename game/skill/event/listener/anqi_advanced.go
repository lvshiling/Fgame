package listener

import (
	"fgame/fgame/core/event"
	anqieventtypes "fgame/fgame/game/anqi/event/types"
	anqitemplate "fgame/fgame/game/anqi/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家暗器进阶
func anqiAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	advanceId := data.(int32)
	newSkillId := anqitemplate.GetAnqiTemplateService().GetAnqiSkill(advanceId)
	oldSkillId := anqitemplate.GetAnqiTemplateService().GetAnqiSkill(advanceId - 1)
	if newSkillId != oldSkillId {
		err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	}

	return
}

func init() {
	gameevent.AddEventListener(anqieventtypes.EventTypeAnqiAdvanced, event.EventListenerFunc(anqiAdvanced))
}
