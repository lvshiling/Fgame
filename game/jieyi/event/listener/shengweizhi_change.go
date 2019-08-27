package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	jieyieventtypes "fgame/fgame/game/jieyi/event/types"
	"fgame/fgame/game/jieyi/jieyi"
	jieyilogic "fgame/fgame/game/jieyi/logic"
	playerjieyi "fgame/fgame/game/jieyi/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeShengWeiZhiChange, event.EventListenerFunc(jieYiShengWeiZhiChange))
}

// 声威值改变
func jieYiShengWeiZhiChange(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	nameLev := jieYiManager.GetNameLevel()
	shengWei := jieYiManager.GetShengWeiZhi()
	playerId := pl.GetId()
	jieyi.GetJieYiService().UpdateShengWeiZhi(playerId, nameLev, shengWei)
	jieyilogic.JieYiPropertyChange(pl)
	return
}
