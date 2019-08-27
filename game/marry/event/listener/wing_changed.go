package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	wingeventtypes "fgame/fgame/game/wing/event/types"
	playerwing "fgame/fgame/game/wing/player"
)

//战翼改变
func wingChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	if marryInfo.Status == marrytypes.MarryStatusTypeUnmarried ||
		marryInfo.Status == marrytypes.MarryStatusTypeDivorce {
		return
	}

	wingManager := pl.GetPlayerDataManager(playertypes.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingId := wingManager.GetWingId()
	scMsg := pbuitl.BuildSCMarryWingChange(pl.GetId(), wingId)
	pl.SendMsg(scMsg)

	spl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)
	if spl == nil {
		return
	}

	spl.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeWingChanged, event.EventListenerFunc(wingChanged))
}
