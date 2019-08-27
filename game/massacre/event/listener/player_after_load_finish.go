package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/massacre/pbutil"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}

	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeMassacre) {
		return
	}

	//暗器信息
	scMassacreGet := pbutil.BuildSCMassacreInfo(p.GetMassacreInfo())
	p.SendMsg(scMassacreGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
