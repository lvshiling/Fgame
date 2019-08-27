package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	skillcommon "fgame/fgame/game/skill/common"
	skilleventtypes "fgame/fgame/game/skill/event/types"
	skilltemplate "fgame/fgame/game/skill/template"
)

//玩家技能添加下发
func playerSkillAdd(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	skillObj, ok := data.(skillcommon.SkillObject)
	if !ok {
		return
	}
	skillTypeId := skillObj.GetSkillId()
	level := skillObj.GetLevel()

	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(skillTypeId, level)
	if skillTemplate == nil {
		return
	}
	skillId := int32(skillTemplate.TemplateId())
	err = questlogic.FillQuestData(pl, questtypes.QuestSubTypeActiveSkill, skillId)
	return
}

func init() {
	gameevent.AddEventListener(skilleventtypes.EventTypeSkillAdd, event.EventListenerFunc(playerSkillAdd))
}
