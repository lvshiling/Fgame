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
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeTokenLevelChange, event.EventListenerFunc(jieYiTokenChange))
}

// 信物等级改变
func jieYiTokenChange(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	tokenLev := jieYiManager.GetPlayerJieYiObj().GetTokenLevel()
	playerId := pl.GetId()
	jieyi.GetJieYiService().UpdateTokenLevel(playerId, tokenLev)
	jieyilogic.JieYiPropertyChange(pl)
	return
}
