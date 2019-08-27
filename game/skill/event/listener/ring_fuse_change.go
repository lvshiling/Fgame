package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	ringeventtypes "fgame/fgame/game/ring/event/types"
	ringtemplate "fgame/fgame/game/ring/template"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家特戒融合等级改变事件
func ringFuseChange(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	eventData := data.(*ringeventtypes.PlayerRingFuseChangeEventData)
	lastItemId := eventData.GetLastItemId()
	curItemId := eventData.GetCurItemId()

	newSkillId := int32(0)
	newTemp := ringtemplate.GetRingTemplateService().GetRingTemplate(curItemId)
	if newTemp != nil {
		newSkillId = newTemp.SkillId
	}
	oldSkillId := int32(0)
	oldTemp := ringtemplate.GetRingTemplateService().GetRingTemplate(lastItemId)
	if oldTemp != nil {
		oldSkillId = oldTemp.SkillId
	}

	if newSkillId != oldSkillId {
		err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	}

	return
}

func init() {
	gameevent.AddEventListener(ringeventtypes.EventTypeRingFuseChange, event.EventListenerFunc(ringFuseChange))
}
