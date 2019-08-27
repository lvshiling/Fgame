package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/lingtongdev/pbutil"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	for classType := types.LingTongDevSysTypeMin; classType <= types.LingTongDevSysTypeMax; classType++ {
		lingTongInfo := manager.GetLingTongDevInfo(classType)
		if lingTongInfo == nil {
			continue
		}
		container := manager.GetLingTongDevOtherMap(classType)
		classType := int32(lingTongInfo.GetClassType())
		scLingTongDevGet := pbutil.BuildSCLingTongDevGet(classType, lingTongInfo, container)
		pl.SendMsg(scLingTongDevGet)
	}
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
