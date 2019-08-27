package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	teameventtypes "fgame/fgame/game/team/event/types"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
	teamtypes "fgame/fgame/game/team/types"
)

//玩家离队
func teamPlayerBeLeaved(target event.EventTarget, data event.EventData) (err error) {
	teamData, ok := target.(*team.TeamObject)
	if !ok {
		return
	}
	beLeavedMember, ok := data.(*team.TeamMemberObject)
	if !ok {
		return
	}
	//记录玩家被踢时间
	teamData.SetKickTime(beLeavedMember.GetPlayerId())

	name := beLeavedMember.GetName()
	teamName := teamData.GetTeamName()
	//广播
	scTeamBroadcast := pbutil.BuildSCTeamBroadcast(int32(teamtypes.TeamBroadcastTypeleaved), name, teamData)
	teamlogic.BroadcastMsg(teamData, scTeamBroadcast)

	beLeavedPlayerId := beLeavedMember.GetPlayerId()
	beLeavedPlayer := player.GetOnlinePlayerManager().GetPlayerById(beLeavedPlayerId)
	if beLeavedPlayer == nil {
		return
	}
	ctx := scene.WithPlayer(context.Background(), beLeavedPlayer)
	beLeavedPlayer.Post(message.NewScheduleMessage(onTeamPlayerBeLeaved, ctx, teamName, nil))
	return nil
}

func onTeamPlayerBeLeaved(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	teamName := result.(string)
	scTeamByLeavedToLeave := pbutil.BuildSCTeamByLeavedToLeave(teamName)
	pl.SendMsg(scTeamByLeavedToLeave)
	pl.SyncTeam(0, "", 0)
	// mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	// mananger.ResetTeam()
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamPlayerBeLeaved, event.EventListenerFunc(teamPlayerBeLeaved))
}
