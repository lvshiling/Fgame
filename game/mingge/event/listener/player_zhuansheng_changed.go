package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/mingge/pbutil"
	playermingge "fgame/fgame/game/mingge/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
)

//玩家转生变化
func playerZhuanShengChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	zhuanShu := data.(int32)
	level := pl.GetLevel()
	manager := pl.GetPlayerDataManager(types.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	mingGongTypeMap := manager.CheckMingGongActivate(level, zhuanShu)
	if len(mingGongTypeMap) == 0 {
		return
	}
	mingLiMap := manager.GetMingLiMap()
	scMingGeMingGongActivate := pbutil.BuildSCMingGeMingGongActivate(mingLiMap, mingGongTypeMap)
	pl.SendMsg(scMingGeMingGongActivate)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerZhuanShengChanged, event.EventListenerFunc(playerZhuanShengChanged))
}
