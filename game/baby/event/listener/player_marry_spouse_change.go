package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/baby/baby"
	babylogic "fgame/fgame/game/baby/logic"
	playerbaby "fgame/fgame/game/baby/player"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//玩家婚姻状态变化
func playerMarrySpouseChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)

	spouseId := manager.GetSpouseId()
	coupleBabyList := baby.GetBabyService().GetCoupleBabyInfo(spouseId)
	babyManager.LoadAllCoupleBaby(coupleBabyList)
	babylogic.BabyPropertyChanged(pl)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypePlayerMarrySpouseChange, event.EventListenerFunc(playerMarrySpouseChanged))
}
