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

//玩家衣橱套装失效
func playerWardrobeRemove(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	wardrobeObj := data.(*playerwardrobe.PlayerWardrobeObject)
	if wardrobeObj.GetIsActive() {
		return
	}
	typ := wardrobeObj.GetType()
	subType := wardrobeObj.GetSubType()
	scWardrobeRemove := pbutil.BuildSCWardrobeRemove(typ, subType)
	pl.SendMsg(scWardrobeRemove)

	//属性变更
	wardrobelogic.WardrobePropertyChanged(pl)

	//技能
	wardrobeTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuTemplate(typ, subType)
	if wardrobeTemplate == nil {
		return
	}
	if wardrobeTemplate.SkillId == 0 && wardrobeTemplate.SkillId2 == 0 {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	peiYangLevel := manager.GetWardrobePeiYangNum(typ)
	oldSkillId, oldSkillId2 := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSkillId(typ, subType, peiYangLevel)
	skilllogic.TempSkillChange(pl, oldSkillId, 0)
	skilllogic.TempSkillChange(pl, oldSkillId2, 0)
	return
}

func init() {
	gameevent.AddEventListener(wardrobeeventtypes.EventTypeWardrobeRemove, event.EventListenerFunc(playerWardrobeRemove))
}
