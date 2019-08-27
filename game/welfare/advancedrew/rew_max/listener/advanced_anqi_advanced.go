package listener

import (
	"fgame/fgame/core/event"
	anqiventtypes "fgame/fgame/game/anqi/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewrewlogic "fgame/fgame/game/welfare/advancedrew/rew_max/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家暗器进阶
func playerAnqiAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}

	// 进阶奖励(每个类型一个活动id，时间、顺序自由定义)
	advancedType := welfaretypes.AdvancedTypeAnqi
	advancedrewrewlogic.UpdateAdvancedRewData(pl, advanceId, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(anqiventtypes.EventTypeAnqiAdvanced, event.EventListenerFunc(playerAnqiAdavanced))
}
