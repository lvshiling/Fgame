package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	xueduneventtypes "fgame/fgame/game/xuedun/event/types"
	playerxuedun "fgame/fgame/game/xuedun/player"
	xueduntemplate "fgame/fgame/game/xuedun/template"
)

//玩家血盾升阶
func playerXueDunUpgrade(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	oldSkillId, ok := data.(int32)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	xueDunInfo := manager.GetXueDunInfo()
	number := xueDunInfo.GetNumber()
	star := xueDunInfo.GetStar()
	bloodShieldTemplate := xueduntemplate.GetXueDunTemplateService().GetXueDunNumber(number, star)
	if bloodShieldTemplate == nil {
		return
	}
	newSkillId := bloodShieldTemplate.SpellId
	err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(xueduneventtypes.EventTypeXueDunUpgrade, event.EventListenerFunc(playerXueDunUpgrade))
}
