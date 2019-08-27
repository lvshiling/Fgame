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

//队长转让
func teamCaptainTransfer(target event.EventTarget, data event.EventData) (err error) {
	oldCaptainPlayer, ok := target.(player.Player)
	if !ok {
		return
	}
	teamData, ok := data.(*team.TeamObject)
	if !ok {
		return
	}
	teamName := teamData.GetTeamName()
	oldCaptainName := oldCaptainPlayer.GetName()

	//广播转让
	scTeamBroadcast := pbutil.BuildSCTeamBroadcast(int32(teamtypes.TeamBroadcastTypeTransfer), oldCaptainName, teamData)
	teamlogic.BroadcastMsg(teamData, scTeamBroadcast)

	for _, mem := range teamData.GetMemberList() {
		pl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if pl == nil {
			continue
		}
		ctx := scene.WithPlayer(context.Background(), pl)
		pl.Post(message.NewScheduleMessage(onTeamCaptainTransfer, ctx, teamName, nil))
	}
	return nil
}

func onTeamCaptainTransfer(ctx context.Context, result interface{}, err error) error {

	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	teamName := result.(string)
	pl.SyncTeam(pl.GetTeamId(), teamName, pl.GetTeamPurpose())
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamCaptainTransfer, event.EventListenerFunc(teamCaptainTransfer))
}
