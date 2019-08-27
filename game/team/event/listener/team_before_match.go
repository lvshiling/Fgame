package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	teameventtypes "fgame/fgame/game/team/event/types"
	"fgame/fgame/game/team/pbutil"
	playerteam "fgame/fgame/game/team/player"
	"fgame/fgame/game/team/team"
)

//队伍匹配条件不满足
func teamMatchConditionFailed(target event.EventTarget, data event.EventData) (err error) {
	teamData, ok := target.(*team.TeamObject)
	if !ok {
		return
	}
	memberIdList, ok := data.([]int64)
	if !ok {
		return
	}
	if len(memberIdList) == 0 {
		return
	}
	captainObj := teamData.GetCaptain()
	if captainObj == nil {
		return
	}
	playerId := captainObj.GetPlayerId()
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}
	scTeamMatchConditionFail := pbutil.BuildSCTeamMatchConditionFail(teamData, memberIdList)
	pl.SendMsg(scTeamMatchConditionFail)

	mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	mananger.SetMatchCondtionFailedList(memberIdList)
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamMatchNoEough, event.EventListenerFunc(teamMatchConditionFailed))
}
