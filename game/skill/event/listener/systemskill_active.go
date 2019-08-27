package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	systemskilleventtypes "fgame/fgame/game/systemskill/event/types"
	playersysskill "fgame/fgame/game/systemskill/player"
	systemskill "fgame/fgame/game/systemskill/systemskill"

	systemskilltemplate "fgame/fgame/game/systemskill/template"
)

//玩家系统技能激活
func playerSystemSkillActive(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	eventData := data.(*systemskilleventtypes.SystemSkillEventData)
	typ := eventData.GetType()
	subType := eventData.GetSubType()
	manager := pl.GetPlayerDataManager(types.PlayerSystemSkillDataManagerType).(*playersysskill.PlayerSystemSkillDataManager)
	level := manager.GetSystemSkillLevelByTyp(typ, subType)
	newSkillId := systemskilltemplate.GetSystemSkillTemplateService().GetSkillId(typ, subType, level)
	err = skilllogic.TempSkillChange(pl, 0, newSkillId)
	systemskill.SystemSkillPropertyChange(pl, typ)
	return
}

func init() {
	gameevent.AddEventListener(systemskilleventtypes.EventTypeSystemSkillActive, event.EventListenerFunc(playerSystemSkillActive))
}
