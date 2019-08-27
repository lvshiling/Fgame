package team

import (
	teamtypes "fgame/fgame/game/team/types"
)

type TeamPlayerInviteDealCreateEventData struct {
	teamObj        *TeamObject
	result         teamtypes.TeamResultType
	invitePlayerId int64
}

func (d *TeamPlayerInviteDealCreateEventData) GetTeam() *TeamObject {
	return d.teamObj
}

func (d *TeamPlayerInviteDealCreateEventData) GetResult() teamtypes.TeamResultType {
	return d.result
}

func (d *TeamPlayerInviteDealCreateEventData) GetInvitePlayerId() int64 {
	return d.invitePlayerId
}

func CreateTeamPlayerInviteDealCreateEventData(teamObj *TeamObject, invitePlayerId int64, result teamtypes.TeamResultType) *TeamPlayerInviteDealCreateEventData {
	d := &TeamPlayerInviteDealCreateEventData{
		teamObj:        teamObj,
		invitePlayerId: invitePlayerId,
		result:         result,
	}
	return d
}

type TeamPlayerInviteDealJoinEventData struct {
	teamObj *TeamObject
	result  teamtypes.TeamResultType
}

func (d *TeamPlayerInviteDealJoinEventData) GetTeam() *TeamObject {
	return d.teamObj
}

func (d *TeamPlayerInviteDealJoinEventData) GetResult() teamtypes.TeamResultType {
	return d.result
}

func CreateTeamPlayerInviteDealJoinEventData(teamObj *TeamObject, result teamtypes.TeamResultType) *TeamPlayerInviteDealJoinEventData {
	d := &TeamPlayerInviteDealJoinEventData{
		teamObj: teamObj,

		result: result,
	}
	return d
}

type TeamPlayerLeaveEventData struct {
	teamObj *TeamObject
	pos     int32
}

func (d *TeamPlayerLeaveEventData) GetTeam() *TeamObject {
	return d.teamObj
}

func (d *TeamPlayerLeaveEventData) GetPos() int32 {
	return d.pos
}

func CreateTeamPlayerLeaveEventData(teamObj *TeamObject, pos int32) *TeamPlayerLeaveEventData {
	d := &TeamPlayerLeaveEventData{
		teamObj: teamObj,

		pos: pos,
	}
	return d
}
