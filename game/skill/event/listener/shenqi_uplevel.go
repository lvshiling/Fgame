package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shenqieventtypes "fgame/fgame/game/shenqi/event/types"
	shenqitemplate "fgame/fgame/game/shenqi/template"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家神器升级事件
func shenQiUpLevel(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	eventData := data.(*shenqieventtypes.PlayerShenQiUpLevelEventData)
	typ := eventData.GetShenQiType()
	oldLevel := eventData.GetShenQiOldLevel()
	newLevel := eventData.GetShenQiNewLevel()
	newSkillId := shenqitemplate.GetShenQiTemplateService().GetShenQiSkillId(typ, newLevel)
	oldSkillId := shenqitemplate.GetShenQiTemplateService().GetShenQiSkillId(typ, oldLevel)
	if newSkillId != oldSkillId {
		err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	}

	return
}

func init() {
	gameevent.AddEventListener(shenqieventtypes.EventTypeShenQiUpLevel, event.EventListenerFunc(shenQiUpLevel))
}
