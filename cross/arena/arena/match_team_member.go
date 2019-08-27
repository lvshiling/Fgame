package arena

import (
	arenascene "fgame/fgame/cross/arena/scene"
	arenatemplate "fgame/fgame/game/arena/template"
	playertypes "fgame/fgame/game/player/types"
	skillcommon "fgame/fgame/game/skill/common"
	"fmt"
)

type MatchTeamMember struct {
	serverId        int32                     //服务器id
	playerId        int64                     //玩家id
	force           int64                     //玩家战力
	name            string                    //玩家名字
	level           int32                     //玩家等级
	role            playertypes.RoleType      //玩家角色
	sex             playertypes.SexType       //性别
	fashionId       int32                     //时装id
	reliveTime      int32                     //复活次数
	robot           bool                      //机器人
	battleProperies map[int32]int64           //系统属性
	skillList       []skillcommon.SkillObject //技能列表
	winCount        int32                     //连胜次数
}

func (o *MatchTeamMember) String() string {
	return fmt.Sprintf("id:%d,名字:%s", o.playerId, o.name)
}

func (o *MatchTeamMember) GetPlayerId() int64 {
	return o.playerId
}

func (o *MatchTeamMember) GetForce() int64 {
	return o.force
}

func (o *MatchTeamMember) GetName() string {
	return o.name
}

func (o *MatchTeamMember) GetLevel() int32 {
	return o.level
}

func (o *MatchTeamMember) GetRole() playertypes.RoleType {
	return o.role
}

func (o *MatchTeamMember) GetSex() playertypes.SexType {
	return o.sex
}

func (o *MatchTeamMember) GetFashionId() int32 {
	return o.fashionId
}

func (o *MatchTeamMember) GetReliveTime() int32 {
	return o.reliveTime
}

func (o *MatchTeamMember) GetRemainReliveTime() int32 {
	return arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RebornAmountMax - o.reliveTime
}

func (o *MatchTeamMember) GetServerId() int32 {
	return o.serverId
}

func (o *MatchTeamMember) GetBattleProperties() map[int32]int64 {
	return o.battleProperies
}

func (o *MatchTeamMember) GetSkillList() []skillcommon.SkillObject {
	return o.skillList
}

func (o *MatchTeamMember) GetRobot() bool {
	return o.robot
}

func (o *MatchTeamMember) GetWinCount() int32 {
	return o.winCount
}

func CreateMatchTeamMemberObject(
	serverId int32,
	playerId int64,
	force int64,
	name string,
	level int32,
	role playertypes.RoleType,
	sex playertypes.SexType,
	fashionId int32,
	battleProperties map[int32]int64,
	skillList []skillcommon.SkillObject,
	robot bool,
	winCount int32) *MatchTeamMember {

	to := &MatchTeamMember{
		serverId:        serverId,
		playerId:        playerId,
		force:           force,
		name:            name,
		level:           level,
		role:            role,
		sex:             sex,
		fashionId:       fashionId,
		battleProperies: battleProperties,
		skillList:       skillList,
		robot:           robot,
		winCount:        winCount,
	}
	return to
}

func convertToTeamMemberObject(m *MatchTeamMember) *arenascene.TeamMemberObject {
	o := arenascene.CreateTeamMemberObject(
		m.GetServerId(),
		m.GetPlayerId(),
		m.GetForce(),
		m.GetName(),
		m.GetLevel(),
		m.GetRole(),
		m.GetSex(),
		m.GetFashionId(),
		m.GetRobot(),
		m.GetWinCount(),
	)
	o.SetStatus(arenascene.MemberStatusOnline)
	return o
}

func convertToTeamMemberObjectList(mList []*MatchTeamMember) []*arenascene.TeamMemberObject {
	oList := make([]*arenascene.TeamMemberObject, 0, len(mList))

	for _, m := range mList {
		oList = append(oList, convertToTeamMemberObject(m))
	}
	return oList
}
