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
func teamPlayerLeave(target event.EventTarget, data event.EventData) (err error) {
	leavePlayer, ok := target.(player.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*team.TeamPlayerLeaveEventData)
	if !ok {
		return
	}

	teamData := eventData.GetTeam()
	pos := eventData.GetPos()
	name := leavePlayer.GetName()

	//修改队伍
	ctx := scene.WithPlayer(context.Background(), leavePlayer)
	leavePlayer.Post(message.NewScheduleMessage(onTeamPlayerLeave, ctx, nil, nil))

	//解散了
	if len(teamData.GetMemberList()) == 0 {
		scTeamLeave := pbutil.BuildSCTeamLeave(name)
		leavePlayer.SendMsg(scTeamLeave)
		return
	}

	teamName := teamData.GetTeamName()
	scTeamLeave := pbutil.BuildSCTeamLeave(teamName)
	leavePlayer.SendMsg(scTeamLeave)

	//广播数据
	scTeamBroadcast := pbutil.BuildSCTeamBroadcast(int32(teamtypes.TeamBroadcastTypeleave), name, teamData)
	teamlogic.BroadcastMsg(teamData, scTeamBroadcast)

	//队长转让
	if pos == 0 {
		for _, mem := range teamData.GetMemberList() {
			pl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
			if pl == nil {
				continue
			}
			ctx := scene.WithPlayer(context.Background(), pl)
			pl.Post(message.NewScheduleMessage(onTeamCaptainTransfer, ctx, teamName, nil))
		}
	}
	return nil
}

func onTeamPlayerLeave(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	pl.SyncTeam(0, "", 0)

	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamPlayerLeave, event.EventListenerFunc(teamPlayerLeave))
}
