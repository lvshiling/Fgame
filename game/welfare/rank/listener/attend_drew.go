package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	ranklogic "fgame/fgame/game/welfare/rank/logic"
)

//玩家活动抽奖
func playerAttendDrew(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*welfareeventtypes.PlayerAttendDrewEventData)
	if !ok {
		return
	}

	//更新次数排行榜
	ranklogic.UpdateAddCountRankData(pl, eventData.GetGroupId(), eventData.GetAttendNum())
	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeAttendDrew, event.EventListenerFunc(playerAttendDrew))
}
