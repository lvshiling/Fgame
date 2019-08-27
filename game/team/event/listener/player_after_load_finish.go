package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	propertytypes "fgame/fgame/game/property/types"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	playerId := pl.GetId()
	teamData := team.GetTeamService().GetTeamByPlayerId(playerId)
	//没有队伍 清空队伍
	if teamData == nil {
		pl.SyncTeam(0, "", 0)
		// manager.SetTeam(0, "")
	} else {
		teamId := teamData.GetTeamId()
		teamName := teamData.GetTeamName()
		teamPurpose := teamData.GetTeamPurpose()
		//设置队伍名称
		pl.SyncTeam(teamId, teamName, teamPurpose)

		member, _ := teamData.GetMember(playerId)
		member.SetOnline(true)

		hp := pl.GetHP()
		maxHp := pl.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
		scTeamMember := pbutil.BuildSCTeamMemberLogin(pl.GetId(), hp, maxHp)
		//广播队内成员
		teamlogic.BroadcastPlayerMsg(teamData, playerId, scTeamMember)
		scTeamGet := pbutil.BuildSCTeamGet(teamData, true, playerId)
		pl.SendMsg(scTeamGet)

		capMember := teamData.GetCaptain()
		//队长
		if capMember.GetPlayerId() == pl.GetId() {
			applyList := teamData.GetAllApplyList()
			scTeamApplyGet := pbutil.BuildSCTeamApplyGet(applyList)
			pl.SendMsg(scTeamApplyGet)
		}
		scTeamMatchCondtionPrepareBroadcast := pbutil.BuildSCTeamMatchCondtionPrepareBroadcast(playerId)
		teamlogic.BroadcastMsg(teamData, scTeamMatchCondtionPrepareBroadcast)
	}
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
