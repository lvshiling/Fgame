package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	teamcommon "fgame/fgame/game/team/common"
	teamtypes "fgame/fgame/game/team/types"
)

type PlayerTeamManager struct {
	p           scene.Player
	teamId      int64
	teamName    string
	teamPurpose teamtypes.TeamPurposeType
}

func (m *PlayerTeamManager) GetTeamId() int64 {
	return m.teamId
}

func (m *PlayerTeamManager) GetTeamName() string {
	return m.teamName
}

func (m *PlayerTeamManager) GetTeamPurpose() teamtypes.TeamPurposeType {
	return m.teamPurpose
}

func (m *PlayerTeamManager) SyncTeam(teamId int64, teamName string, teamPurpose teamtypes.TeamPurposeType) {
	m.teamId = teamId
	m.teamName = teamName
	m.teamPurpose = teamPurpose
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerTeamChanged, m.p, nil)
}

func CreatePlayerTeamManagerWithObject(p scene.Player, playerTeamObj teamcommon.PlayerTeamObject) *PlayerTeamManager {
	m := &PlayerTeamManager{
		p:           p,
		teamId:      playerTeamObj.GetTeamId(),
		teamName:    playerTeamObj.GetTeamName(),
		teamPurpose: playerTeamObj.GetTeamPurpose(),
	}
	return m
}
