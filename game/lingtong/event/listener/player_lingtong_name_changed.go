package listener

import (
	"fgame/fgame/core/event"
	crosspbutil "fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	playerlingtong "fgame/fgame/game/lingtong/player"
	"fgame/fgame/game/player"
)

//灵童重命名
func playerLingTongRename(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	lingTong := p.GetLingTong()
	if lingTong == nil {
		return
	}
	lingTongInfoObject := data.(*playerlingtong.PlayerLingTongInfoObject)
	if lingTongInfoObject.GetLingTongId() != lingTong.GetLingTongId() {
		return
	}

	lingTong.UpdateName(lingTongInfoObject.GetLingTongName())

	if p.IsCross() {
		lingTongNameChanged := crosspbutil.BuildLingTongNameChanged(lingTongInfoObject.GetLingTongName())
		p.SendCrossMsg(lingTongNameChanged)
		return
	}

	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeLingTongRename, event.EventListenerFunc(playerLingTongRename))
}
