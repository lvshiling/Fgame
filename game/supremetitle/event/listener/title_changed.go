package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	supremetitleeventtypes "fgame/fgame/game/supremetitle/event/types"
	playersupremetitle "fgame/fgame/game/supremetitle/player"
	"fgame/fgame/game/title/pbutil"
	playertitle "fgame/fgame/game/title/player"
)

//至尊称号改变
func supremeTitleChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	m := pl.GetPlayerDataManager(playertypes.PlayerSupremeTitleDataManagerType).(*playersupremetitle.PlayerSupremeTitleDataManager)
	titleId := m.GetTitleId()
	//卸下称号
	titleManager := pl.GetPlayerDataManager(playertypes.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	curTitleId := titleManager.GetTitleId()
	if curTitleId != 0 && titleId != 0 {
		titleManager.TitleNoWear()
		scTitleUnload := pbutil.BuildSCTitleUnload(curTitleId)
		pl.SendMsg(scTitleUnload)
	}
	return
}

func init() {
	gameevent.AddEventListener(supremetitleeventtypes.EventTypeSupremeTitleChanged, event.EventListenerFunc(supremeTitleChanged))
}
