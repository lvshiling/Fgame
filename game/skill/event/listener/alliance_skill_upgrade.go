package listener

import (
	"fgame/fgame/core/event"
	alliancetemplate "fgame/fgame/game/alliance/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	skilllogic "fgame/fgame/game/skill/logic"

	allianceeventtypes "fgame/fgame/game/alliance/event/types"
)

//玩家仙盟仙术升级
func allianceSkillUpgrade(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}
	skillData := data.(*allianceeventtypes.PlayerAllianceSkillUpgradeEventData)
	level := skillData.GetSkillLevel()
	typ := skillData.GetSkillType()

	curTemp := alliancetemplate.GetAllianceTemplateService().GetAllianceSkillTemplateByType(level, typ)
	if curTemp == nil {
		return
	}
	preTemp := alliancetemplate.GetAllianceTemplateService().GetAllianceSkillTemplateByType(level-1, typ)

	newSkillId := curTemp.SkillId
	oldSkillId := int32(0)
	if preTemp != nil {
		oldSkillId = preTemp.SkillId
	}

	if newSkillId != oldSkillId {
		err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	}

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypePlayerAllianceSkillUpgrade, event.EventListenerFunc(allianceSkillUpgrade))
}
