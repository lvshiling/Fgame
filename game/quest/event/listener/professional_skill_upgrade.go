package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	skilleventtypes "fgame/fgame/game/skill/event/types"
	playerskill "fgame/fgame/game/skill/player"
)

//职业技能升级
func professionalSkillUpgrade(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	skillId, ok := data.(int32)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	totalLevel := manager.GetTotaolLevel()
	skillObj := manager.GetSkill(skillId)
	if skillObj == nil {
		return
	}
	level := skillObj.Level

	err = professionalSkillLevel(pl, skillId, level)
	if err != nil {
		return
	}

	err = professionalSkillTotalLevel(pl, totalLevel)
	if err != nil {
		return
	}
	return
}

//指定职业技能达到X级
func professionalSkillLevel(pl player.Player, skillId int32, level int32) (err error) {
	return questlogic.SetQuestData(pl, questtypes.QuestSubTypeProfessionalSkillLevel, skillId, level)
}

//职业技能总等级为X级
func professionalSkillTotalLevel(pl player.Player, totalLevel int32) (err error) {
	return questlogic.SetQuestData(pl, questtypes.QuestSubTypeProfessionalSkillTotalLevel, 0, totalLevel)
}

func init() {
	gameevent.AddEventListener(skilleventtypes.EventTypeSkillUpgrade, event.EventListenerFunc(professionalSkillUpgrade))
}
