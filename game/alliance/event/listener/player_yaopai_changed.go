package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancelogic "fgame/fgame/game/alliance/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//玩家腰牌变化
func playerYaoPaiChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	//同步腰牌
	alliancelogic.SnapYaoPaiChanged(p)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypePlayerYaoPaiChanged, event.EventListenerFunc(playerYaoPaiChanged))
}
