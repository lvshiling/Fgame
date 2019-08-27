package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	shenmoeventtypes "fgame/fgame/game/shenmo/event/types"
	playershenmo "fgame/fgame/game/shenmo/player"
)

//玩家功勋值改变
func playerShenMoGongXunNumChanged(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	manager := p.GetPlayerDataManager(playertypes.PlayerShenMoWarDataManagerType).(*playershenmo.PlayerShenMoDataManager)
	gongXunNum := manager.GetShenMoInfo().GetGongXunNum()
	p.SetShenMoGongXunNum(gongXunNum)

	return
}

func init() {
	gameevent.AddEventListener(shenmoeventtypes.EventTypeShenMoGongXunNumChanged, event.EventListenerFunc(playerShenMoGongXunNumChanged))
}
