package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerfound "fgame/fgame/game/found/player"
	foundtypes "fgame/fgame/game/found/types"
	majoreventtypes "fgame/fgame/game/major/event/types"
	majortemplate "fgame/fgame/game/major/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//玩家进入双修场景
func majorTasksRecord(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	fubenTemplate, ok := data.(majortemplate.MajorTemplate)
	if !ok {
		return
	}

	majorType := fubenTemplate.GetMajorType()
	resType, ok := foundtypes.MajorTypeToFoundResType(majorType)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)
	manager.IncreFoundResJoinTimes(resType)
	return
}

func init() {
	gameevent.AddEventListener(majoreventtypes.EventTypePlayerEnterMajorScene, event.EventListenerFunc(majorTasksRecord))
}
