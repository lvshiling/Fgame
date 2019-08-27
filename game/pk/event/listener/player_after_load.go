package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/pk/pbutil"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	curValue := pl.GetPkValue()
	onlineTime := pl.GetPkOnlineTime()
	loginTime := pl.GetPkLoginTime()
	scMsg := pbutil.BuildSCPKValueChanged(curValue, onlineTime, loginTime)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
