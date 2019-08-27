package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerfound "fgame/fgame/game/found/player"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	teamtypes "fgame/fgame/game/team/types"
	teamcopyeventtypes "fgame/fgame/game/teamcopy/event/types"
)

//记录组队副本事件
func teamCopyTasksRecord(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	typ := data.(teamtypes.TeamPurposeType)

	resType, ok := foundtypes.MaterialTeamTypeToFoundResType(typ)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)
	manager.IncreFoundResJoinTimes(resType)
	return
}

func init() {
	gameevent.AddEventListener(teamcopyeventtypes.EventTypeTeamCopyFinishSucess, event.EventListenerFunc(teamCopyTasksRecord))

}
