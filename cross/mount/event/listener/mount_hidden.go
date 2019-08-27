package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/mount/pbutil"
	"fgame/fgame/cross/player/player"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
)

//坐骑隐藏或穿戴
func mountHidden(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(*player.Player)
	if !ok {
		return
	}

	hiddenMount := pl.IsMountHidden()
	isPlayerMountSync := pbutil.BuildISPlayerMountSync(hiddenMount)
	pl.SendMsg(isPlayerMountSync)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowMountHidden, event.EventListenerFunc(mountHidden))
}
