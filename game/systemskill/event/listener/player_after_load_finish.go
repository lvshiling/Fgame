package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	"fgame/fgame/game/systemskill/pbutil"
	playersysskill "fgame/fgame/game/systemskill/player"
	systemskilltemplate "fgame/fgame/game/systemskill/template"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerSystemSkillDataManagerType).(*playersysskill.PlayerSystemSkillDataManager)
	sysSkillMap := manager.GetSystemSkillAllMap()
	scSystemSkillAllGet := pbutil.BuildSCSystemSkillAllGet(sysSkillMap)
	pl.SendMsg(scSystemSkillAllGet)

	for _, sysTypeSkill := range sysSkillMap {
		for _, obj := range sysTypeSkill.GetSysSkillMap() {
			typ := obj.Type
			subType := obj.SubType
			level := obj.Level
			newSkillId := systemskilltemplate.GetSystemSkillTemplateService().GetSkillId(typ, subType, level)
			skilllogic.TempSkillChange(pl, 0, newSkillId)
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
