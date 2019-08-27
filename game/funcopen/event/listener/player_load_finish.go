package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopenlogic "fgame/fgame/game/funcopen/logic"
	"fgame/fgame/game/funcopen/pbutil"
	playerfuncopen "fgame/fgame/game/funcopen/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	_, err = funcopenlogic.CheckFuncOpen(p)
	if err != nil {
		return
	}
	funcOpenManager := p.GetPlayerDataManager(playertypes.PlayerFuncOpenDataManagerType).(*playerfuncopen.PlayerFuncOpenDataManager)
	scFuncOpenList := pbutil.BuildSCFuncOpenList(funcOpenManager.GetOpenFuncList())
	p.SendMsg(scFuncOpenList)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
