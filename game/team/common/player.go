package common

import "fgame/fgame/game/team/types"

type PlayerTeamObject interface {
	GetTeamId() int64
	GetTeamPurpose() types.TeamPurposeType
	GetTeamName() string
}

type playerTeamObject struct {
	teamId      int64
	teamPurpose types.TeamPurposeType
	teamName    string
}

func (o *playerTeamObject) GetTeamId() int64 {
	return o.teamId
}

func (o *playerTeamObject) GetTeamName() string {
	return o.teamName
}

func (o *playerTeamObject) GetTeamPurpose() types.TeamPurposeType {
	return o.teamPurpose
}

func CreatePlayerTeamObject(teamId int64, teamName string, purpose types.TeamPurposeType) PlayerTeamObject {
	obj := &playerTeamObject{}
	obj.teamId = teamId
	obj.teamName = teamName
	obj.teamPurpose = purpose
	return obj
}
