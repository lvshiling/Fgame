package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerfound "fgame/fgame/game/found/player"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	xianfueventtypes "fgame/fgame/game/xianfu/event/types"
	xianfutypes "fgame/fgame/game/xianfu/types"
)

//记录仙府挑战事件
func xianfuTasksRecord(target event.EventTarget, data event.EventData) (err error) {

	pl := target.(player.Player)
	eventData := data.(*xianfueventtypes.XianFuChallengeEventData)
	typ := eventData.GetType()
	num := eventData.GetNum()
	if num <= 0 {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)

	var resType foundtypes.FoundResourceType
	switch typ {
	case xianfutypes.XianfuTypeExp:
		resType = foundtypes.FoundResourceTypeExp
	case xianfutypes.XianfuTypeSilver:
		resType = foundtypes.FoundResourceTypeSilver
	default:
		return
	}

	manager.IncreFoundResJoinTimesBatch(resType, num)
	return
}

func init() {
	gameevent.AddEventListener(xianfueventtypes.EventTypeXianFuChallenge, event.EventListenerFunc(xianfuTasksRecord))

}
