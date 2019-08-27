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
	teamtypes "fgame/fgame/game/team/types"
)

//邀请玩家
func teamPlayerInvitePlayer(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	inviteId := data.(int64)
	invitePlayer := player.GetOnlinePlayerManager().GetPlayerById(inviteId)
	if invitePlayer == nil {
		return
	}
	ctx := scene.WithPlayer(context.Background(), invitePlayer)
	invitePlayer.Post(message.NewScheduleMessage(onPlayerInvitePlayer, ctx, pl, nil))
	return nil
}

func onPlayerInvitePlayer(ctx context.Context, result interface{}, err error) error {

	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)

	invitePlayer := result.(player.Player)
	invitePlayerId := invitePlayer.GetId()
	invitePlayerName := invitePlayer.GetName()
	inviteTyp := teamtypes.TeamInviteTypeCreate

	mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	mananger.TeamInvite(inviteTyp, invitePlayerId)

	//推送
	scTeamInviteToInvited := pbutil.BuildSCTeamInviteToInvited(int32(inviteTyp), invitePlayerId, invitePlayerName, invitePlayerName)
	pl.SendMsg(scTeamInviteToInvited)
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamPlayerInvitePlayer, event.EventListenerFunc(teamPlayerInvitePlayer))
}
