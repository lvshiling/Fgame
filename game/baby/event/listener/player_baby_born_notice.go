package listener

import (
	"fgame/fgame/core/event"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	"fgame/fgame/game/baby/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

// 玩家宝宝出生提醒
func playerBabyBornNotice(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*babyeventtypes.PlayerBabyBornMsgNoticeEventData)
	if !ok {
		return
	}

	scMsg := pbutil.BuildSCBabyBornMessageNotice(eventData.GetBornTime(), eventData.GetNoticeTime())
	pl.SendMsg(scMsg)

	spl := player.GetOnlinePlayerManager().GetPlayerById(pl.GetSpouseId())
	if spl != nil {
		spl.SendMsg(scMsg)
	}
	return
}

func init() {
	gameevent.AddEventListener(babyeventtypes.EventTypeBabyBornNotice, event.EventListenerFunc(playerBabyBornNotice))
}
