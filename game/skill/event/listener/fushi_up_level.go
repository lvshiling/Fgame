package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fushieventtypes "fgame/fgame/game/fushi/event/types"
	fushitemplate "fgame/fgame/game/fushi/template"
	"fgame/fgame/game/player"
	skilllogic "fgame/fgame/game/skill/logic"
)

func init() {
	gameevent.AddEventListener(fushieventtypes.FushiEventTypeLevelChanged, event.EventListenerFunc(fushiUpLevel))
}

func fushiUpLevel(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	fushiInfo, ok := data.(*fushieventtypes.FuShiLevelChangedData)
	if !ok {
		return
	}

	typ := fushiInfo.GetType()
	level := fushiInfo.GetFushiLevel()

	fushiLevelTemp := fushitemplate.GetFuShiTemplateService().GetFuShiLevelByFuShiTypeAndLevel(typ, level)
	if fushiLevelTemp == nil {
		return
	}

	newSkillId := fushiLevelTemp.SkillId

	// 符石激活
	if level == 1 {
		err = skilllogic.TempSkillChange(pl, 0, newSkillId)
	} else {
		temp := fushitemplate.GetFuShiTemplateService().GetFuShiLevelByFuShiTypeAndLevel(typ, level-1)
		if temp == nil {
			return
		}
		oldSkillId := temp.SkillId
		err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	}

	return
}
