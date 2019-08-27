package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
	godsiegelogic "fgame/fgame/game/godsiege/logic"
)

//玩家完成排队
func godsiegeFinishLineUp(target event.EventTarget, data event.EventData) (err error) {
	lineList := target.([]int64)
	eventData := data.(*godsiegeeventtypes.GodSiegeLineUpFinishEventData)
	// playerId := eventData.GetPlayerId()
	godType := eventData.GetGodType()

	// TODO:xzk 本服排队成功未处理
	// pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	// if pl != nil {
	// 	pl.GodSiegeCancleLineUp()
	// 	scGodSiegeLineUpSuccess := pbutil.BuildSCGodSiegeLineUpSuccess(godType)
	// 	pl.SendMsg(scGodSiegeLineUpSuccess)
	// 	//进入跨服神兽攻城
	// 	crosslogic.CrossPlayerDataLogin(pl)
	// }

	godsiegelogic.BroadGodSiegeLineUpChanged(int32(godType), -1, lineList)
	return
}

func init() {
	gameevent.AddEventListener(godsiegeeventtypes.EventTypeGodSiegePlayerLineUpFinish, event.EventListenerFunc(godsiegeFinishLineUp))
}
