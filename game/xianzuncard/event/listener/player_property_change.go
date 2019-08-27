package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	xianzuncardeventtypes "fgame/fgame/game/xianzuncard/event/types"
	xianzuncardlogic "fgame/fgame/game/xianzuncard/logic"
	"fgame/fgame/game/xianzuncard/pbutil"
	xianzuncardtypes "fgame/fgame/game/xianzuncard/types"
)

func init() {
	gameevent.AddEventListener(xianzuncardeventtypes.EventTypeXianZunCardExpire, event.EventListenerFunc(playerPropertyChange))
}

func playerPropertyChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	typ, ok := data.(xianzuncardtypes.XianZunCardType)
	if !ok {
		return
	}

	xianzuncardlogic.PropertyChanged(pl)

	scXianZunCardNotice := pbutil.BuildSCXianZunCardNotice(int32(typ))
	pl.SendMsg(scXianZunCardNotice)
	return
}
