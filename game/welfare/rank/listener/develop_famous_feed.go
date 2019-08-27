package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	ranklogic "fgame/fgame/game/welfare/rank/logic"
)

//玩家名人培养
func playerDevelopFamousFeed(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*welfareeventtypes.PlayerFamousFeedEventData)
	if !ok {
		return
	}

	//更新数量排行榜
	ranklogic.UpdateSetCountRankData(pl, eventData.GetGroupId(), eventData.GetTotalFavorbaleNum())
	ranklogic.UpdateSetDayCountRankData(pl, eventData.GetGroupId(), eventData.GetDayFavorbaleNum())

	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeDevelopFamousFeed, event.EventListenerFunc(playerDevelopFamousFeed))
}
