package listener

import (
	"fgame/fgame/core/event"
	crosseventtypes "fgame/fgame/game/cross/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertitle "fgame/fgame/game/title/player"
	"fgame/fgame/game/title/title"
	titletypes "fgame/fgame/game/title/types"
)

//玩家跨服断开
func playerCrossExit(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	titleId := manager.GetTitleId()
	if titleId == 0 {
		return
	}
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	if titleTemplate == nil {
		return
	}
	if titleTemplate.GetTitleType() != titletypes.TitleTypeShenMo {
		return
	}
	manager.TitleNoWear()
	return
}

func init() {
	gameevent.AddEventListener(crosseventtypes.EventTypePlayerCrossExit, event.EventListenerFunc(playerCrossExit))
}
