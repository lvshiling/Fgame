package listener

import (
	"fgame/fgame/core/event"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	playerbodyshield "fgame/fgame/game/bodyshield/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/guaji/pbutil"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//坐骑进阶
func bodyshieldAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	bodyshieldDataManager := pl.GetPlayerDataManager(playertypes.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	advanceId := int32(bodyshieldDataManager.GetBodyShiedInfo().AdvanceId)
	scGuaJiAdvanceUpdateList := pbutil.BuildSCGuaJiAdvanceUpdateList(guajitypes.GuaJiAdvanceTypeBodyshield, advanceId)
	pl.SendMsg(scGuaJiAdvanceUpdateList)
	return
}

func init() {
	gameevent.AddEventListener(bodyshieldeventtypes.EventTypeBodyShieldAdvanced, event.EventListenerFunc(bodyshieldAdvanced))
}
