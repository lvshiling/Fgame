package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	teameventtypes "fgame/fgame/game/team/event/types"
	"fgame/fgame/game/team/pbutil"
	playerteam "fgame/fgame/game/team/player"
	"fgame/fgame/game/team/team"
	teamtypes "fgame/fgame/game/team/types"
)

//邀请玩家
func teamTeamInvitePlayer(target event.EventTarget, data event.EventData) (err error) {
	teamObject, ok := target.(*team.TeamObject)
	if !ok {
		return
	}
	inviteId := data.(int64)
	invitePlayer := player.GetOnlinePlayerManager().GetPlayerById(inviteId)
	if invitePlayer == nil {
		return
	}
	ctx := scene.WithPlayer(context.Background(), invitePlayer)
	invitePlayer.Post(message.NewScheduleMessage(onTeamInvitePlayer, ctx, teamObject, nil))
	return nil
}

func onTeamInvitePlayer(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	teamObject := result.(*team.TeamObject)

	teamId := teamObject.GetTeamId()
	teamName := teamObject.GetTeamName()
	inviteTyp := teamtypes.TeamInviteTypeJoin

	mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	mananger.TeamInvite(inviteTyp, teamId)

	//推送
	scTeamInviteToInvited := pbutil.BuildSCTeamInviteToInvited(int32(inviteTyp), teamId, pl.GetName(), teamName)
	pl.SendMsg(scTeamInviteToInvited)
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamTeamInvitePlayer, event.EventListenerFunc(teamTeamInvitePlayer))
}
