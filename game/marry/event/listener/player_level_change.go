package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
)

//玩家等级变化
func playerLevelChanged(target event.EventTarget, data event.EventData) (err error) {
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

	playerId := pl.GetId()
	level := data.(int32)
	spl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)
	scMarryLevelChange := pbuitl.BuildSCMarryLevelChange(playerId, level)
	pl.SendMsg(scMarryLevelChange)
	if spl == nil {
		return
	}
	spl.SendMsg(scMarryLevelChange)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerLevelChanged, event.EventListenerFunc(playerLevelChanged))
}
