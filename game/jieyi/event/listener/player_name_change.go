package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//玩家名字变化
func playerNameChanged(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)

	//同步成员信息
	jieyi.GetJieYiService().UpdatePlayerName(p.GetId(), p.GetName())

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerNameChanged, event.EventListenerFunc(playerNameChanged))
}
