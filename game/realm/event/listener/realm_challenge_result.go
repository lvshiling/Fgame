package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	playerrealm "fgame/fgame/game/realm/player"
	"fgame/fgame/game/realm/realm"
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
		realm.GetRealmRankService().PairChallegeFail(pl.GetId())
	} else {
		manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
		level := manager.GetTianJieTaLevel()
		pl.SetRealm(level)
	}
	return
}

func init() {
	gameevent.AddEventListener(realmeventtypes.EventTypeRealmResult, event.EventListenerFunc(realmChallengeResult))
}
