package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/bagua/bagua"
	baguaeventtypes "fgame/fgame/game/bagua/event/types"
	playerbagua "fgame/fgame/game/bagua/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//天劫塔挑战结果
func realmChallengeResult(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	sucessful, ok := data.(bool)
	if !ok {
		return
	}
	if !sucessful {
		bagua.GetBaGuaService().PairChallegeFail(pl.GetId())
	} else {
		manager := pl.GetPlayerDataManager(types.PlayerBaGuaDataManagerType).(*playerbagua.PlayerBaGuaDataManager)
		level := manager.GetLevel()
		pl.SetBaGua(level)
	}
	return
}

func init() {
	gameevent.AddEventListener(baguaeventtypes.EventTypeBaGuaResult, event.EventListenerFunc(realmChallengeResult))
}
