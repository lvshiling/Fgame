package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	wardrobeeventtypes "fgame/fgame/game/wardrobe/event/types"
	wardrobelogic "fgame/fgame/game/wardrobe/logic"
	"fgame/fgame/game/wardrobe/pbutil"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetemplate "fgame/fgame/game/wardrobe/template"
)

//玩家衣橱套装激活
func playerWardrobeActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	wardrobeObj := data.(*playerwardrobe.PlayerWardrobeObject)
	if !wardrobeObj.GetIsActive() {
		return
	}
	scWardrobeGet := pbutil.BuildSCWardrobeActive(wardrobeObj)
	pl.SendMsg(scWardrobeGet)

	//属性变更
	wardrobelogic.WardrobePropertyChanged(pl)

	typ := wardrobeObj.GetType()
	subType := wardrobeObj.GetSubType()
	wardrobeTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuTemplate(typ, subType)
	if wardrobeTemplate == nil {
		return
	}
	if wardrobeTemplate.SkillId == 0 && wardrobeTemplate.SkillId2 == 0 {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	//技能
	peiYangLevel := manager.GetWardrobePeiYangNum(typ)
	newSkillId, newSkillId2 := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSkillId(typ, subType, peiYangLevel)
	skilllogic.TempSkillChange(pl, 0, newSkillId)
	skilllogic.TempSkillChange(pl, 0, newSkillId2)
	return
}

func init() {
	gameevent.AddEventListener(wardrobeeventtypes.EventTypeWardrobeActive, event.EventListenerFunc(playerWardrobeActivate))
}
