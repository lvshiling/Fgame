package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shenqieventtypes "fgame/fgame/game/shenqi/event/types"
	"fgame/fgame/game/shenqi/pbutil"
	playershenqi "fgame/fgame/game/shenqi/player"
)

//灵气值变化
func playerShenQiLingQiNumChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	shenQiObj, ok := data.(*playershenqi.PlayerShenQiObject)
	if !ok {
		return
	}

	scMsg := pbutil.BuildSCShenQiLingQiNumChanged(shenQiObj.LingQiNum)
	pl.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(shenqieventtypes.EventTypeShenQiLingQiNumChanged, event.EventListenerFunc(playerShenQiLingQiNumChanged))
}
