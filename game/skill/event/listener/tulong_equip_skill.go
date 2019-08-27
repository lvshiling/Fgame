package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	skilllogic "fgame/fgame/game/skill/logic"
	tulongequipeventtypes "fgame/fgame/game/tulongequip/event/types"
	playertulongequip "fgame/fgame/game/tulongequip/player"
	tulongequiptemplate "fgame/fgame/game/tulongequip/template"
)

//屠龙装备技能升级
func playerTuLongEquipSkillUpgrade(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	skillObj, ok := data.(*playertulongequip.PlayerTuLongSuitSkillObject)
	if !ok {
		return
	}
	curLevel := skillObj.GetLevel()
	oldSkillId := int32(0)

	preSkillTemp := tulongequiptemplate.GetTuLongEquipTemplateService().GetTuLongEquipTemplateSkill(skillObj.GetSuitType(), curLevel-1)
	if preSkillTemp != nil {
		oldSkillId = preSkillTemp.SkillId
	}

	skillTemp := tulongequiptemplate.GetTuLongEquipTemplateService().GetTuLongEquipTemplateSkill(skillObj.GetSuitType(), curLevel)
	if skillTemp == nil {
		return
	}

	skilllogic.TempSkillChange(pl, oldSkillId, skillTemp.SkillId)
	return
}

// 套装技能（
func init() {
	gameevent.AddEventListener(tulongequipeventtypes.EventTypeTuLongEquipSkillUpgrade, event.EventListenerFunc(playerTuLongEquipSkillUpgrade))
}
