package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家结婚状态变化
func playerMarry(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	marryManager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	if !marryManager.IsTrueMarry() {
		return
	}

	questlogic.SetQuestData(pl, questtypes.QuestSubTypeMarryTimes, 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypePlayerMarrySpouseChange, event.EventListenerFunc(playerMarry))
}
