package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/relive/pbutil"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//加载完成后
func playerRelive(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(scene.Player)
	if !ok {
		return
	}
	isPlayerReliveSync := pbutil.BuildISPlayerReliveSync(p)
	//发送同步
	p.SendMsg(isPlayerReliveSync)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerRelive, event.EventListenerFunc(playerRelive))
}
