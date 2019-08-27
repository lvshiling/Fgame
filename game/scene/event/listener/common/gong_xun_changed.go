package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/shenmo/pbutil"
)

//玩家功勋值改变
func playerShenMoGongXunNumChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(scene.Player)
	if !ok {
		return
	}
	if p.GetScene() == nil {
		return
	}
	gongXunNum := p.GetShenMoGongXunNum()

	scShenMoGongXunNumChanged := pbutil.BuildSCShenMoGongXunNumChanged(gongXunNum)
	p.SendMsg(scShenMoGongXunNumChanged)

	scPlayerGongXunChanged := pbutil.BuildSCPlayerGongXunChanged(gongXunNum)
	p.SendMsg(scPlayerGongXunChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShenMoGongXunNumChanged, event.EventListenerFunc(playerShenMoGongXunNumChanged))
}
