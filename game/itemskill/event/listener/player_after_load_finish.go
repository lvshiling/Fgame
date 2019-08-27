package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/itemskill/pbutil"
	playeritemskill "fgame/fgame/game/itemskill/player"
	itemskilltemplate "fgame/fgame/game/itemskill/template"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerItemSkillDataManagerType).(*playeritemskill.PlayerItemSkillDataManager)
	skillMap := manager.GetItemSkillAllMap()
	scItemSkillAllGet := pbutil.BuildSCItemSkillAllGet(skillMap)
	pl.SendMsg(scItemSkillAllGet)

	for _, obj := range skillMap {
		typ := obj.Typ
		level := obj.Level
		newSkillId := itemskilltemplate.GetItemSkillTemplateService().GetSkillId(typ, level)
		skilllogic.TempSkillChange(pl, 0, newSkillId)
	}
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
