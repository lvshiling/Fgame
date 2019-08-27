package listener

import (
	"fgame/fgame/core/event"
	baguaeventtypes "fgame/fgame/game/bagua/event/types"
	"fgame/fgame/game/bagua/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//夫妻助战邀请对方无应答
func pairInviteNoAnswer(target event.EventTarget, data event.EventData) (err error) {
	playerId, ok := target.(int64)
	if !ok {
		return
	}
	spouseName, ok := data.(string)
	if !ok {
		return
	}

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}
	scBaGuaSpouseRefused := pbutil.BuildSCBaGuaSpouseRefused(spouseName)
	pl.SendMsg(scBaGuaSpouseRefused)
	return
}

func init() {
	gameevent.AddEventListener(baguaeventtypes.EventTypeBaGuaPairNoAnswer, event.EventListenerFunc(pairInviteNoAnswer))
}
