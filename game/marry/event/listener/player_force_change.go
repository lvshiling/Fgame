package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
)

//玩家战力变化
func playerForceChanged(target event.EventTarget, data event.EventData) (err error) {
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
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	force := propertyManager.GetForce()
	spl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)

	scMarryForceChange := pbuitl.BuildSCMarryForceChange(pl.GetId(), force)
	pl.SendMsg(scMarryForceChange)
	if spl == nil {
		return
	}
	spl.SendMsg(scMarryForceChange)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerForceChanged, event.EventListenerFunc(playerForceChanged))
}
