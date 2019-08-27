package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrylogic "fgame/fgame/game/marry/logic"
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
	spouseName := manager.GetSpouseName()
	pl.SetSpouse(spouseName)
	spouseId := manager.GetSpouseId()
	pl.SetSpouseId(spouseId)
	marrylogic.MarryPropertyChanged(pl)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypePlayerMarrySpouseChange, event.EventListenerFunc(playerMarrySpouseChanged))
}
