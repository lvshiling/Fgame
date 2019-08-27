package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/shenmo/pbutil"
)

//玩家功勋值改变
func playerShenMoGongXunNumChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	if !p.IsCross() {
		return
	}
	gongXunNum := p.GetShenMoGongXunNum()
	siPlayerGongXunNumChanged := pbutil.BuildSIPlayerGongXunNumChanged(gongXunNum)
	p.SendCrossMsg(siPlayerGongXunNumChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShenMoGongXunNumChanged, event.EventListenerFunc(playerShenMoGongXunNumChanged))
}
