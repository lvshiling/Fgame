package team

import (
	playertypes "fgame/fgame/game/player/types"
)

type TeamMemberObject struct {
	serverId   int32                //服务器id
	playerId   int64                //玩家id
	force      int64                //玩家战力
	online     bool                 //是否在线
	name       string               //玩家名字
	level      int32                //玩家等级
	role       playertypes.RoleType //玩家角色
	sex        playertypes.SexType  //性别
	fashionId  int32                //时装id
	zhuanSheng int32                //转生

}

func (obj *TeamMemberObject) GetPlayerId() int64 {
	return obj.playerId
}

func (obj *TeamMemberObject) GetForce() int64 {
	return obj.force
}

func (obj *TeamMemberObject) UpdateForce(force int64) {
	obj.force = force
}

func (obj *TeamMemberObject) GetOnline() bool {
	return obj.online
}

func (obj *TeamMemberObject) SetOnline(online bool) {
	obj.online = online
}

func (obj *TeamMemberObject) GetName() string {
	return obj.name
}

func (obj *TeamMemberObject) GetLevel() int32 {
	return obj.level
}

func (obj *TeamMemberObject) SetLevel(level int32) {
	obj.level = level
}

func (obj *TeamMemberObject) SetName(name string) {
	obj.name = name
}

func (obj *TeamMemberObject) SetSex(sex playertypes.SexType) {
	obj.sex = sex
}

func (obj *TeamMemberObject) GetRole() playertypes.RoleType {
	return obj.role
}

func (obj *TeamMemberObject) GetSex() playertypes.SexType {
	return obj.sex
}
func (obj *TeamMemberObject) GetFashionId() int32 {
	return obj.fashionId
}

func (obj *TeamMemberObject) SetFashionId(fashionId int32) {
	obj.fashionId = fashionId
}

func (obj *TeamMemberObject) GetZhuanSheng() int32 {
	return obj.zhuanSheng
}

func (obj *TeamMemberObject) SetZhuanSheng(zhuanSheng int32) {
	obj.zhuanSheng = zhuanSheng
}

func (obj *TeamMemberObject) GetServerId() int32 {
	return obj.serverId
}

func NewTeamMemberObject(serverId int32, playerId int64, force int64, online bool, name string, level int32, role playertypes.RoleType, sex playertypes.SexType, fashionId int32, zhuanSheng int32) *TeamMemberObject {
	data := &TeamMemberObject{
		serverId:   serverId,
		playerId:   playerId,
		force:      force,
		online:     online,
		name:       name,
		level:      level,
		role:       role,
		sex:        sex,
		fashionId:  fashionId,
		zhuanSheng: zhuanSheng,
	}
	return data
}
