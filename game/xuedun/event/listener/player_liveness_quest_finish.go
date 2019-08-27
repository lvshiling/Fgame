package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	livenesseventtypes "fgame/fgame/game/liveness/event/types"
	playerliveness "fgame/fgame/game/liveness/player"
	livenesstemplate "fgame/fgame/game/liveness/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerxuedun "fgame/fgame/game/xuedun/player"
)

//血炼值改变
func playerLivenessChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	livenessQuest := data.(*playerliveness.PlayerLivenessQuestObject)
	questId := livenessQuest.GetQuestId()
	livenessLevelTempalte, flag := livenesstemplate.GetHuoYueTempalteService().GetHuoYueLevelTemplate(questId, p.GetLevel())
	if !flag {
		return
	}
	addBlodd := int64(livenessLevelTempalte.XueLian)
	manager := p.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	manager.XueDunBloodChanged(addBlodd)
	return
}

func init() {
	gameevent.AddEventListener(livenesseventtypes.EventTypeLivenessChanged, event.EventListenerFunc(playerLivenessChanged))
}
