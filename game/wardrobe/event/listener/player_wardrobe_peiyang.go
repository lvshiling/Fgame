package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	wardrobeeventtypes "fgame/fgame/game/wardrobe/event/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetemplate "fgame/fgame/game/wardrobe/template"
)

//玩家衣橱套装培养
func playerWardrobePeiYang(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*wardrobeeventtypes.WardrobePeiYangEventData)
	typ := eventData.GetType()
	oldLevel := eventData.GetOldLevel()
	newLevel := eventData.GetNewLevel()

	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	suitMap := manager.GetWardrobeMapByType(typ)
	if suitMap == nil {
		return
	}

	for subType, wardrobeObj := range suitMap {
		if !wardrobeObj.GetIsActive() {
			continue
		}
		wardrobeTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuTemplate(typ, subType)
		if wardrobeTemplate == nil {
			continue
		}
		if wardrobeTemplate.SkillId == 0 {
			continue
		}
		//技能
		oldSkillId, oldSkillId2 := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSkillId(typ, subType, oldLevel)
		newSkillId, newSkillId2 := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSkillId(typ, subType, newLevel)
		skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
		skilllogic.TempSkillChange(pl, oldSkillId2, newSkillId2)
	}

	return
}

func init() {
	gameevent.AddEventListener(wardrobeeventtypes.EventTypeWardrobePeiYangUpgrade, event.EventListenerFunc(playerWardrobePeiYang))
}
