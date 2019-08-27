package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家噬魂幡进阶
func shiHunFanAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	advanceId := data.(int32)
	newSkillId := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanSkill(advanceId)
	oldSkillId := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanSkill(advanceId - 1)
	if newSkillId != oldSkillId {
		err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	}

	return
}

func init() {
	gameevent.AddEventListener(shihunfaneventtypes.EventTypeShiHunFanAdvanced, event.EventListenerFunc(shiHunFanAdvanced))
}
