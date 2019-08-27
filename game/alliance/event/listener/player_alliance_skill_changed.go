package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancetemplate "fgame/fgame/game/alliance/template"
	alliancetypes "fgame/fgame/game/alliance/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	"fmt"
)

//玩家仙盟仙术改变
func playerAllianceSkillChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	oldSkillMap, ok := data.(map[alliancetypes.AllianceSkillType]*playeralliance.PlayerAllianceSkillObject)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	newSkillMap := manager.GetEffectiveAllianceSkillList()
	for skillType, skillObj := range newSkillMap {
		_, ok := oldSkillMap[skillType]
		if ok {
			continue
		}

		tem := alliancetemplate.GetAllianceTemplateService().GetAllianceSkillTemplateByType(skillObj.GetLevel(), skillType)
		if tem == nil {
			return fmt.Errorf("allianceSkill:仙术技能模板不存在,level:%d,type:%d", skillObj.GetLevel(), skillType)
		}
		skilllogic.TempSkillChange(pl, 0, tem.SkillId)
	}

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypePlayerAllianceSkillChanged, event.EventListenerFunc(playerAllianceSkillChanged))
}
