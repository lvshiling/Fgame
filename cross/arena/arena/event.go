package arena

import arenascene "fgame/fgame/cross/arena/scene"

//匹配事件
type MatchEventData struct {
	team1 *arenascene.ArenaTeam
	team2 *arenascene.ArenaTeam
}

func (d *MatchEventData) GetTeam1() *arenascene.ArenaTeam {
	return d.team1
}

func (d *MatchEventData) GetTeam2() *arenascene.ArenaTeam {
	return d.team2
}

func CreateMatchEventData(team1 *arenascene.ArenaTeam, team2 *arenascene.ArenaTeam) *MatchEventData {
	d := &MatchEventData{}
	d.team1 = team1
	d.team2 = team2
	return d
}

//创建
type TeamCreateEventData struct {
	teamList [][]*arenascene.TeamMemberObject
}

func (d *TeamCreateEventData) GetTeamList() [][]*arenascene.TeamMemberObject {
	return d.teamList
}

func CreateTeamCreateEventData(teamList [][]*arenascene.TeamMemberObject) *TeamCreateEventData {
	d := &TeamCreateEventData{}
	d.teamList = teamList
	return d
}
