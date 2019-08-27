package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	livenesseventtypes "fgame/fgame/game/liveness/event/types"
	"fgame/fgame/game/liveness/pbutil"
	playerliveness "fgame/fgame/game/liveness/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//活跃值改变
func playerLivenessChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	livenessQuest := data.(*playerliveness.PlayerLivenessQuestObject)
	manager := p.GetPlayerDataManager(types.PlayerLivenessDataManagerType).(*playerliveness.PlayerLivenessDataManager)
	livenessObj := manager.GetLiveness()
	scLivenessNumChanged := pbutil.BuildSCLivenessNumChanged(livenessQuest, livenessObj.GetLiveness())
	p.SendMsg(scLivenessNumChanged)
	return
}

func init() {
	gameevent.AddEventListener(livenesseventtypes.EventTypeLivenessChanged, event.EventListenerFunc(playerLivenessChanged))
}
