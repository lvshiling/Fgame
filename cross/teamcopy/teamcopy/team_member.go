package teamcopy

import (
	teamcopyscene "fgame/fgame/cross/teamcopy/scene"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	skillcommon "fgame/fgame/game/skill/common"
	teamtypes "fgame/fgame/game/team/types"
	teamcopytemplate "fgame/fgame/game/teamcopy/template"
	"fmt"
)

type BattleTeamMember struct {
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
	damage          int64                     //伤害
	purpose         teamtypes.TeamPurposeType //队伍标识
}

func (o *BattleTeamMember) String() string {
	return fmt.Sprintf("id:%d,名字:%s", o.playerId, o.name)
}

func (o *BattleTeamMember) GetPlayerId() int64 {
	return o.playerId
}

func (o *BattleTeamMember) GetForce() int64 {
	return o.force
}

func (o *BattleTeamMember) GetName() string {
	return o.name
}

func (o *BattleTeamMember) GetLevel() int32 {
	return o.level
}

func (o *BattleTeamMember) GetRole() playertypes.RoleType {
	return o.role
}

func (o *BattleTeamMember) GetSex() playertypes.SexType {
	return o.sex
}

func (o *BattleTeamMember) GetFashionId() int32 {
	return o.fashionId
}

func (o *BattleTeamMember) GetReliveTime() int32 {
	return o.reliveTime
}

func (o *BattleTeamMember) GetRemainReliveTime() int32 {
	return teamcopytemplate.GetTeamCopyTemplateService().GetTeamCopyTempalte(o.purpose).ResurrectionNumber - o.reliveTime
}

func (o *BattleTeamMember) GetServerId() int32 {
	return o.serverId
}

func (o *BattleTeamMember) GetBattleProperties() map[int32]int64 {
	return o.battleProperies
}

func (o *BattleTeamMember) GetSkillList() []skillcommon.SkillObject {
	return o.skillList
}

func (o *BattleTeamMember) GetRobot() bool {
	return o.robot
}

func (o *BattleTeamMember) GetDamage() int64 {
	return o.damage
}

func (o *BattleTeamMember) GetTeamPurpose() teamtypes.TeamPurposeType {
	return o.purpose
}

func CreateBattleTeamMemberObject(
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
	damage int64,
	purpose teamtypes.TeamPurposeType) *BattleTeamMember {

	to := &BattleTeamMember{
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
		damage:          damage,
		purpose:         purpose,
	}
	return to
}

func convertToTeamMemberObject(m *BattleTeamMember) *teamcopyscene.TeamMemberObject {
	o := teamcopyscene.CreateTeamMemberObject(
		m.GetServerId(),
		m.GetPlayerId(),
		m.GetForce(),
		m.GetName(),
		m.GetLevel(),
		m.GetRole(),
		m.GetSex(),
		m.GetFashionId(),
		m.GetRobot(),
		m.GetDamage(),
		m.GetTeamPurpose(),
	)
	o.SetStatus(teamcopyscene.MemberStatusOnline)
	return o
}

func convertToTeamMemberObjectList(mList []*BattleTeamMember) []*teamcopyscene.TeamMemberObject {
	oList := make([]*teamcopyscene.TeamMemberObject, 0, len(mList))

	for _, m := range mList {
		oList = append(oList, convertToTeamMemberObject(m))
	}
	return oList
}

func convertMemberFromRobotPlayer(pl scene.RobotPlayer, purpoes teamtypes.TeamPurposeType) *BattleTeamMember {
	serverId := pl.GetServerId()
	playerId := pl.GetId()
	force := pl.GetForce()
	name := pl.GetOriginName()
	level := pl.GetLevel()
	role := pl.GetRole()
	sex := pl.GetSex()
	fashionId := pl.GetFashionId()
	damage := int64(0)
	skillList := make([]skillcommon.SkillObject, 0, len(pl.GetAllSkills()))
	for _, skill := range pl.GetAllSkills() {
		skillList = append(skillList, skill)
	}
	memObj := CreateBattleTeamMemberObject(
		serverId,
		playerId,
		force,
		name,
		level,
		role,
		sex,
		fashionId,
		pl.GetAllSystemBattleProperties(),
		skillList,
		true,
		damage,
		purpoes,
	)
	return memObj
}
