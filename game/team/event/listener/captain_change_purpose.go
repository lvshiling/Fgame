package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	teameventtypes "fgame/fgame/game/team/event/types"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
)

//玩家队长修改队伍标识
func teamCaptainChangePurpose(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	teamData, ok := data.(*team.TeamObject)
	if !ok {
		return
	}
	pl.SyncTeam(teamData.GetTeamId(), teamData.GetTeamName(), teamData.GetTeamPurpose())
	scTeamCreateHouse := pbutil.BuildSCTeamCreateHouse(int32(teamData.GetTeamPurpose()), teamData)
	pl.SendMsg(scTeamCreateHouse)

	for _, member := range teamData.GetMemberList() {
		if member.GetPlayerId() == pl.GetId() {
			continue
		}
		mpl := player.GetOnlinePlayerManager().GetPlayerById(member.GetPlayerId())
		if mpl == nil {
			continue
		}

		ctx := scene.WithPlayer(context.Background(), mpl)
		mpl.Post(message.NewScheduleMessage(onTeamCaptainChangePurpose, ctx, teamData, nil))
	}

	return nil
}

func onTeamCaptainChangePurpose(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	teamObject := result.(*team.TeamObject)

	teamId := teamObject.GetTeamId()
	teamName := teamObject.GetTeamName()
	teamPurpose := teamObject.GetTeamPurpose()
	pl.SyncTeam(teamId, teamName, teamPurpose)
	//推送
	scTeamPurposeChange := pbutil.BuildSCTeamPurposeChange(pl.GetId(), int32(teamPurpose))
	pl.SendMsg(scTeamPurposeChange)
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamCaptainChangePurpose, event.EventListenerFunc(teamCaptainChangePurpose))
}
