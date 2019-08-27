package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/major/pbutil"
	playermajor "fgame/fgame/game/major/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := p.GetPlayerDataManager(playertypes.PlayerMajorDataManagerType).(*playermajor.PlayerMajorDataManager)
	majorNumObjMap := manager.GetAllMajorNumObj()
	scMsg := pbutil.BuildSCMajorNumNotice(majorNumObjMap)
	p.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
