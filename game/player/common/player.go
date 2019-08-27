package common

import playertypes "fgame/fgame/game/player/types"

type PlayerCommonObject interface {
	GetId() int64
	GetUserId() int64
	GetServerId() int32
	GetPlatform() int32
	GetName() string
	GetRole() playertypes.RoleType
	GetSex() playertypes.SexType
	IsGuaJi() bool
}

type playerCommonObject struct {
	id       int64
	userId   int64
	serverId int32
	platform int32
	name     string
	role     playertypes.RoleType
	sex      playertypes.SexType
	guaJi    bool
}

func (o *playerCommonObject) GetId() int64 {
	return o.id
}

func (o *playerCommonObject) GetUserId() int64 {
	return o.userId
}

func (o *playerCommonObject) GetServerId() int32 {
	return o.serverId
}

func (o *playerCommonObject) GetPlatform() int32 {
	return o.platform
}

func (o *playerCommonObject) GetName() string {
	return o.name
}

func (o *playerCommonObject) GetRole() playertypes.RoleType {
	return o.role
}

func (o *playerCommonObject) GetSex() playertypes.SexType {
	return o.sex
}

func (o *playerCommonObject) IsGuaJi() bool {
	return o.guaJi
}

func NewBasicPlayerCommonObject(id int64, userId int64, serverId int32, name string, role playertypes.RoleType, sex playertypes.SexType, guaJi bool, platform int32) PlayerCommonObject {
	o := &playerCommonObject{
		id:       id,
		userId:   userId,
		serverId: serverId,
		name:     name,
		role:     role,
		sex:      sex,
		guaJi:    guaJi,
		platform: platform,
	}
	return o
}

func NewPlayerCommonObject(id int64, userId int64, serverId int32, name string, role playertypes.RoleType, sex playertypes.SexType, guaJi bool) PlayerCommonObject {
	return NewBasicPlayerCommonObject(id, userId, serverId, name, role, sex, guaJi, 0)
}
