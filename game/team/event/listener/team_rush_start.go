package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	teameventtypes "fgame/fgame/game/team/event/types"
	"fgame/fgame/game/team/pbutil"
	playerteam "fgame/fgame/game/team/player"
)

//队伍匹配条件不满足
func teamMatchMatchRushStart(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	captainPl, ok := data.(player.Player)
	if !ok {
		return
	}

	mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	mananger.SetRushTime()

	scTeamMatchRushStart := pbutil.BuildSCTeamMatchRushStart()
	pl.SendMsg(scTeamMatchRushStart)

	//推送给队长
	scTeamMatchRushToCaptain := pbutil.BuildSCTeamMatchRushToCaptain()
	captainPl.SendMsg(scTeamMatchRushToCaptain)
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamMatchRushStart, event.EventListenerFunc(teamMatchMatchRushStart))
}
