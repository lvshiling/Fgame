package center

import (
	"fgame/fgame/center/store"
	"fgame/fgame/core/storage"
)

type PlayerInfo struct {
	id         int64
	userId     int64
	serverId   int32
	playerId   int64
	playerName string
	role       int32
	sex        int32
	level      int32
	zhuanShu   int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func (info *PlayerInfo) GetId() int64 {
	return info.id
}

func (info *PlayerInfo) GetUserId() int64 {
	return info.userId
}

func (info *PlayerInfo) GetServerId() int32 {
	return info.serverId
}

func (info *PlayerInfo) GetPlayerId() int64 {
	return info.playerId
}

func (info *PlayerInfo) GetPlayerName() string {
	return info.playerName
}

func (info *PlayerInfo) GetRole() int32 {
	return info.role
}

func (info *PlayerInfo) GetSex() int32 {
	return info.sex
}

func (info *PlayerInfo) GetLevel() int32 {
	return info.level
}

func (info *PlayerInfo) GetZhuanShu() int32 {
	return info.zhuanShu
}

func (info *PlayerInfo) FromEntity(e *store.PlayerEntity) {
	info.id = e.Id
	info.userId = e.UserId
	info.serverId = e.ServerId
	info.playerId = e.PlayerId
	info.playerName = e.PlayerName
	info.role = e.Role
	info.sex = e.Sex
	info.level = e.Level
	info.zhuanShu = e.ZhuanShu
	info.updateTime = e.UpdateTime
	info.deleteTime = e.DeleteTime
	info.createTime = e.CreateTime
}

func (po *PlayerInfo) ToEntity() (e storage.Entity, err error) {
	pe := &store.PlayerEntity{}
	pe.Id = po.id
	pe.ServerId = po.serverId
	pe.PlayerId = po.playerId
	pe.PlayerName = po.playerName
	pe.Role = po.role
	pe.Sex = po.sex
	pe.Level = po.level
	pe.ZhuanShu = po.zhuanShu
	pe.UpdateTime = po.updateTime
	pe.CreateTime = po.createTime
	pe.DeleteTime = po.deleteTime
	e = pe
	return
}

func newPlayerInfo() *PlayerInfo {
	info := &PlayerInfo{}
	return info
}
