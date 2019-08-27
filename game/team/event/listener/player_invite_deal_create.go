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
func playerInvitePlayerDealCreate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*team.TeamPlayerInviteDealCreateEventData)
	if !ok {
		return
	}
	playerName := pl.GetName()
	result := eventData.GetResult()
	//拒绝
	if result == teamtypes.TeamResultTypeNo {
		//发送给邀请者
		invitePlayer := player.GetOnlinePlayerManager().GetPlayerById(eventData.GetInvitePlayerId())
		if invitePlayer == nil {
			return
		}
		scTeamInviteBroadcast := pbutil.BuildSCTeamInviteBroadcast(int32(result), playerName)
		invitePlayer.SendMsg(scTeamInviteBroadcast)
		return
	}
	//获取队伍
	teamObj := eventData.GetTeam()

	//广播加入消息
	scTeamBroadcast := pbutil.BuildSCTeamBroadcast(int32(teamtypes.TeamBroadCastTypeInviteAgree), playerName, teamObj)
	teamlogic.BroadcastMsg(teamObj, scTeamBroadcast)

	//设置队伍
	teamId := teamObj.GetTeamId()
	teamName := teamObj.GetTeamName()
	teamPurpose := teamObj.GetTeamPurpose()
	playerJoinTeam(pl, teamId, teamName, teamPurpose)
	for _, mem := range teamObj.GetMemberList() {
		//决策者
		if pl.GetId() == mem.GetPlayerId() {
			continue
		}
		memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if memPl == nil {
			continue
		}
		teamlogic.OnPlayerJoinTeam(memPl, teamId, teamName, teamPurpose)
	}

	return nil
}

func playerJoinTeam(pl player.Player, teamId int64, teamName string, teamPurpose teamtypes.TeamPurposeType) {
	pl.SyncTeam(teamId, teamName, teamPurpose)
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamPlayerInvitePlayerDealCreate, event.EventListenerFunc(playerInvitePlayerDealCreate))
}
