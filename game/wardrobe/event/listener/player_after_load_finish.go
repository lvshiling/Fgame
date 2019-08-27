package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	"fgame/fgame/game/wardrobe/pbutil"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetemplate "fgame/fgame/game/wardrobe/template"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.RefreshAfterLoad()
	wardrobeMap := manager.GetWardrobeMap()

	for typ, suitMap := range wardrobeMap {
		peiYangNum := manager.GetWardrobePeiYangNum(typ)
		for subType, wardrobeObj := range suitMap {
			if !wardrobeObj.GetIsActive() {
				continue
			}
			wardrobeTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuTemplate(typ, subType)
			if wardrobeTemplate == nil {
				continue
			}
			if wardrobeTemplate.SkillId == 0 && wardrobeTemplate.SkillId2 == 0 {
				continue
			}
			newSkillId, newSkillId2 := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSkillId(typ, subType, peiYangNum)
			skilllogic.TempSkillChange(pl, 0, newSkillId)
			skilllogic.TempSkillChange(pl, 0, newSkillId2)
		}
	}

	scWardrobeGet := pbutil.BuildSCWardrobeGet(pl)
	pl.SendMsg(scWardrobeGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
