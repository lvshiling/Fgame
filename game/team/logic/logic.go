package logic

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/team/team"

	teamtypes "fgame/fgame/game/team/types"

	"github.com/golang/protobuf/proto"
)

//广播消息
func BroadcastMsg(t *team.TeamObject, msg proto.Message) {
	for _, member := range t.GetMemberList() {
		pl := player.GetOnlinePlayerManager().GetPlayerById(member.GetPlayerId())
		if pl != nil {
			pl.SendMsg(msg)
		}
	}
}

func BroadcastPlayerMsg(t *team.TeamObject, excludePlayerId int64, msg proto.Message) (err error) {
	for _, member := range t.GetMemberList() {
		playerId := member.GetPlayerId()
		if member.GetPlayerId() == excludePlayerId {
			continue
		}
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		pl.SendMsg(msg)
	}
	return
}

type teamResult struct {
	teamId      int64
	teamName    string
	teamPurpose teamtypes.TeamPurposeType
}

//玩家加入队伍
func OnPlayerJoinTeam(pl player.Player, teamId int64, teamName string, teamPurpose teamtypes.TeamPurposeType) {
	result := &teamResult{
		teamId:      teamId,
		teamName:    teamName,
		teamPurpose: teamPurpose,
	}
	ctx := scene.WithPlayer(context.Background(), pl)
	pl.Post(message.NewScheduleMessage(onPlayerJoinTeam, ctx, result, nil))
}

func onPlayerJoinTeam(ctx context.Context, result interface{}, err error) error {
	tr := result.(*teamResult)
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)

	newTeamId := tr.teamId
	newTeamName := tr.teamName
	newTeamPurpose := tr.teamPurpose
	pl.SyncTeam(newTeamId, newTeamName, newTeamPurpose)
	return nil
}

//获取附近玩家
func GetNearPlayers(pl player.Player) (nearPlayerList []player.Player) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	nearPlayers := s.GetAllPlayers()
	for nearPlayerId, _ := range nearPlayers {
		nearPlayer := player.GetOnlinePlayerManager().GetPlayerById(nearPlayerId)
		if nearPlayer == nil {
			return
		}
		if nearPlayerId == pl.GetId() {
			continue
		}
		teamId := nearPlayer.GetTeamId()
		if teamId != 0 {
			continue
		}
		nearPlayerList = append(nearPlayerList, nearPlayer)
	}
	return
}
