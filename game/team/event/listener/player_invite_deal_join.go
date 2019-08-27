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

//被邀请邀请玩家创建
func playerInvitePlayerDealJoin(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*team.TeamPlayerInviteDealJoinEventData)
	if !ok {
		return
	}
	playerName := pl.GetName()
	//获取队伍
	teamObj := eventData.GetTeam()
	result := eventData.GetResult()
	//拒绝
	if result == teamtypes.TeamResultTypeNo {
		scTeamInviteBroadcast := pbutil.BuildSCTeamInviteBroadcast(int32(result), playerName)
		teamlogic.BroadcastMsg(teamObj, scTeamInviteBroadcast)
		return
	}

	//广播加入消息
	scTeamBroadcast := pbutil.BuildSCTeamBroadcast(int32(teamtypes.TeamBroadCastTypeInviteAgree), playerName, teamObj)
	teamlogic.BroadcastMsg(teamObj, scTeamBroadcast)

	teamId := teamObj.GetTeamId()
	teamName := teamObj.GetTeamName()
	teamPurpose := teamObj.GetTeamPurpose()
	playerJoinTeam(pl, teamId, teamName, teamPurpose)
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamPlayerInvitePlayerDealJoin, event.EventListenerFunc(playerInvitePlayerDealJoin))
}
