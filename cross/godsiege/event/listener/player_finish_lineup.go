package listener

import (
	"fgame/fgame/core/event"
	godsiegelogic "fgame/fgame/cross/godsiege/logic"
	"fgame/fgame/cross/godsiege/pbutil"
	"fgame/fgame/cross/player/player"
	gameevent "fgame/fgame/game/event"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
)

//玩家完成排队
func godsiegeFinishLineUp(target event.EventTarget, data event.EventData) (err error) {
	lineList := target.([]int64)
	eventData := data.(*godsiegeeventtypes.GodSiegeLineUpFinishEventData)
	playerId := eventData.GetPlayerId()
	godType := eventData.GetGodType()

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl != nil {
		isGodSiegeLineUpSuccess := pbutil.BuildISGodSiegeLineUpSuccess(int32(godType))
		pl.SendMsg(isGodSiegeLineUpSuccess)
	}

	godsiegelogic.BroadGodSiegeLineUpChanged(int32(godType), -1, lineList)
	return
}

func init() {
	gameevent.AddEventListener(godsiegeeventtypes.EventTypeGodSiegePlayerLineUpFinish, event.EventListenerFunc(godsiegeFinishLineUp))
}
