package scene

import (
	playertypes "fgame/fgame/game/player/types"
	teamtypes "fgame/fgame/game/team/types"
	teamcopytempalte "fgame/fgame/game/teamcopy/template"
	"fmt"
)

type MemberStatus int32

const (
	MemberStatusOffline MemberStatus = iota //离线
	MemberStatusOnline                      //在线
	MemberStatusFailed                      //失败
	MemberStatusGoAway                      //退出
)

type TeamMemberObject struct {
	serverId   int32                     //服务器id
	playerId   int64                     //玩家id
	force      int64                     //玩家战力
	status     MemberStatus              //是否在线
	name       string                    //玩家名字
	level      int32                     //玩家等级
	role       playertypes.RoleType      //玩家角色
	sex        playertypes.SexType       //性别
	fashionId  int32                     //时装id
	reliveTime int32                     //复活次数
	robot      bool                      //机器人
	damage     int64                     //伤害
	purpose    teamtypes.TeamPurposeType //队伍标识
}

func (o *TeamMemberObject) String() string {
	return fmt.Sprintf("id:%d,名字:%s", o.playerId, o.name)
}

func (o *TeamMemberObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *TeamMemberObject) GetForce() int64 {
	return o.force
}

func (o *TeamMemberObject) SetStatus(status MemberStatus) {
	o.status = status
}

func (o *TeamMemberObject) GetStatus() MemberStatus {
	return o.status
}

func (o *TeamMemberObject) GetName() string {
	return o.name
}

func (o *TeamMemberObject) GetLevel() int32 {
	return o.level
}

func (o *TeamMemberObject) GetRole() playertypes.RoleType {
	return o.role
}

func (o *TeamMemberObject) GetSex() playertypes.SexType {
	return o.sex
}

func (o *TeamMemberObject) GetFashionId() int32 {
	return o.fashionId
}

func (o *TeamMemberObject) AddReliveTime() {
	reliveTime := o.reliveTime
	purpose := o.purpose
	if reliveTime >= teamcopytempalte.GetTeamCopyTemplateService().GetTeamCopyTempalte(purpose).ResurrectionNumber {
		return
	}
	o.reliveTime += 1
}

func (o *TeamMemberObject) AddDamage(damage int64) {
	if damage <= 0 {
		return
	}
	o.damage += damage
}

func (o *TeamMemberObject) GetReliveTime() int32 {
	return o.reliveTime
}

func (o *TeamMemberObject) GetRemainReliveTime() int32 {
	return teamcopytempalte.GetTeamCopyTemplateService().GetTeamCopyTempalte(o.purpose).ResurrectionNumber - o.reliveTime
}

func (o *TeamMemberObject) GetServerId() int32 {
	return o.serverId
}

func (o *TeamMemberObject) IsRobot() bool {
	return o.robot
}

func (o *TeamMemberObject) GetDamage() int64 {
	return o.damage
}

func (o *TeamMemberObject) GetTeamPurpose() teamtypes.TeamPurposeType {
	return o.purpose
}

func CreateTeamMemberObject(
	serverId int32,
	playerId int64,
	force int64,
	name string,
	level int32,
	role playertypes.RoleType,
	sex playertypes.SexType,
	fashionId int32,
	robot bool,
	damage int64,
	purpose teamtypes.TeamPurposeType) *TeamMemberObject {
	to := &TeamMemberObject{
		serverId:  serverId,
		playerId:  playerId,
		force:     force,
		status:    MemberStatusOffline,
		name:      name,
		level:     level,
		role:      role,
		sex:       sex,
		fashionId: fashionId,
		robot:     robot,
		damage:    damage,
		purpose:   purpose,
	}
	return to
}

type TeamObject struct {
	teamId     int64
	purpose    teamtypes.TeamPurposeType
	force      int64
	memberList []*TeamMemberObject
}

func (o *TeamObject) String() string {
	return fmt.Sprintf("队伍:%d,战力:%d,成员:%s", o.teamId, o.force, o.memberList)
}

func (o *TeamObject) GetTeamId() int64 {
	return o.teamId
}

func (o *TeamObject) GetTeamName() string {
	captain := o.GetCaptain()
	if captain == nil {
		return ""
	}
	return captain.GetName()
}

func (o *TeamObject) GetTeamPurpose() teamtypes.TeamPurposeType {
	return o.purpose
}

func (o *TeamObject) GetMemberList() []*TeamMemberObject {
	return o.memberList
}

func (o *TeamObject) GetMember(playerId int64) (mem *TeamMemberObject, pos int32) {
	pos = -1
	for index, member := range o.memberList {
		if member.GetPlayerId() == playerId {
			pos = int32(index)
			mem = member
			break
		}
	}
	return
}

func (o *TeamObject) RemoveMember(memberId int64) {
	mem, pos := o.GetMember(memberId)
	if mem == nil {
		return
	}
	o.memberList = append(o.memberList[:pos], o.memberList[pos+1:]...)
}

func (o *TeamObject) GetCaptain() *TeamMemberObject {
	if len(o.memberList) == 0 {
		return nil
	}
	for _, mem := range o.memberList {
		if mem.status == MemberStatusOffline || mem.status == MemberStatusOnline {
			return mem
		}
	}
	return nil
}

func (o *TeamObject) GetForce() int64 {
	return o.force
}

func (o *TeamObject) IfAllLeave() bool {
	for _, mem := range o.memberList {
		if mem.status == MemberStatusOnline {
			return false
		}
	}
	return true
}

func (o *TeamObject) IfAllRobot() bool {
	for _, mem := range o.memberList {
		if !mem.robot {
			return false
		}
	}
	return true
}

func (o *TeamObject) GetTotalReliveTime() int32 {
	total := int32(0)
	for _, mem := range o.memberList {
		total += mem.GetRemainReliveTime()
	}
	return total
}

func CreateTeamObject(teamId int64, purpose teamtypes.TeamPurposeType, force int64, memberList []*TeamMemberObject) *TeamObject {
	to := &TeamObject{
		teamId:     teamId,
		purpose:    purpose,
		force:      force,
		memberList: memberList,
	}
	return to
}

func CreateTeamWithMemberLists(teamId int64, purpose teamtypes.TeamPurposeType, memListOfList ...[]*TeamMemberObject) *TeamObject {
	teamMemList := make([]*TeamMemberObject, 0, 3)
	for _, memList := range memListOfList {
		teamMemList = append(teamMemList, memList...)
	}
	t := CreateTeamObject(teamId, purpose, 0, teamMemList)
	return t
}
