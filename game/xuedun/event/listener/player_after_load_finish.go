package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	"fgame/fgame/game/xuedun/pbutil"
	playerxuedun "fgame/fgame/game/xuedun/player"
	xueduntemplate "fgame/fgame/game/xuedun/template"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	xueDunInfo := manager.GetXueDunInfo()
	scXueDunGet := pbutil.BuildSCXueDunGet(xueDunInfo)
	pl.SendMsg(scXueDunGet)

	xueDunTemplate := xueduntemplate.GetXueDunTemplateService().GetXueDunNumber(xueDunInfo.GetNumber(), xueDunInfo.GetStar())
	if xueDunTemplate == nil {
		return
	}
	newSkillId := xueDunTemplate.SpellId
	skilllogic.TempSkillChange(pl, 0, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
