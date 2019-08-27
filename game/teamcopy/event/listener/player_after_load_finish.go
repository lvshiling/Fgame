package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/teamcopy/pbutil"
	playerteamcopy "fgame/fgame/game/teamcopy/player"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	manager := p.GetPlayerDataManager(types.PlayerTeamCopyDataManagerType).(*playerteamcopy.PlayerTeamCopyDataManager)

	teamCopyMap := manager.GetTeamCopyMap()
	scTeamCopyAllGet := pbutil.BuildSCTeamCopyAllGet(teamCopyMap)
	p.SendMsg(scTeamCopyAllGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
