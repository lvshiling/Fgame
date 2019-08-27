package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	systemskill "fgame/fgame/game/systemskill/systemskill"

	systemskilleventtypes "fgame/fgame/game/systemskill/event/types"
	playersysskill "fgame/fgame/game/systemskill/player"
	systemskilltemplate "fgame/fgame/game/systemskill/template"
)

//玩家系统技能升级
func playerSystemSkillUpgrade(target event.EventTarget, data event.EventData) (err error) {
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
	oldSkillId := systemskilltemplate.GetSystemSkillTemplateService().GetSkillId(typ, subType, level-1)
	err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	systemskill.SystemSkillPropertyChange(pl, typ)
	return
}

func init() {
	gameevent.AddEventListener(systemskilleventtypes.EventTypeSystemSkillUpgrade, event.EventListenerFunc(playerSystemSkillUpgrade))
}
