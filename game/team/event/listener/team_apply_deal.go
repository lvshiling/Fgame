package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	teameventtypes "fgame/fgame/game/team/event/types"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
	teamtypes "fgame/fgame/game/team/types"
)

//队长对申请决策
func teamApplyDeal(target event.EventTarget, data event.EventData) (err error) {
	teamData := target.(*team.TeamObject)
	teamName := teamData.GetTeamName()
	teamPurpose := teamData.GetTeamPurpose()
	eventData := data.(*teameventtypes.TeamApplyDealEventData)
	applyPlayerId := eventData.GetApplyId()
	result := eventData.GetResult()
	applyPlayer := player.GetOnlinePlayerManager().GetPlayerById(applyPlayerId)
	//拒绝
	if result == teamtypes.TeamResultTypeNo {
		if applyPlayer != nil {
			//推送给申请者
			scTeamNearJoinResultToApply := pbutil.BuildSCTeamNearJoinResultToApply(teamName)
			applyPlayer.SendMsg(scTeamNearJoinResultToApply)
		}
		return
	}

	applyMem, _ := teamData.GetMember(applyPlayerId)
	name := applyMem.GetName()
	//广播所有
	scTeamBroadcast := pbutil.BuildSCTeamBroadcast(int32(teamtypes.TeamBroadcastTypeApplyJoin), name, teamData)
	teamlogic.BroadcastMsg(teamData, scTeamBroadcast)

	if applyPlayer == nil {
		return
	}
	//异步加入队伍
	teamId := teamData.GetTeamId()
	teamlogic.OnPlayerJoinTeam(applyPlayer, teamId, teamName, teamPurpose)
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamApplyDeal, event.EventListenerFunc(teamApplyDeal))
}
