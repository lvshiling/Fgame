package store

import (
	centerpb "fgame/fgame/center/pb"
	coredb "fgame/fgame/core/db"

	"github.com/jinzhu/gorm"
)

type PlayerEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	UserId     int64  `gorm:"column:userId"`
	ServerId   int32  `gorm:"column:serverId"`
	PlayerId   int64  `gorm:"column:playerId"`
	PlayerName string `gorm:"column:playerName"`
	Role       int32  `gorm:"column:role"`
	Sex        int32  `gorm:"column:sex"`
	Level      int32  `gorm:"column:level"`
	ZhuanShu   int32  `gorm:"column:zhuanShu"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *PlayerEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerEntity) TableName() string {
	return "t_player"
}

func (e *PlayerEntity) ConvertToGrpcFormat() (playerInfo *centerpb.PlayerServerInfo) {
	playerInfo = &centerpb.PlayerServerInfo{}
	playerInfo.UserId = e.UserId
	playerInfo.ServerId = e.ServerId
	playerInfo.PlayerId = e.PlayerId
	playerInfo.PlayerName = e.PlayerName
	playerInfo.Role = e.Role
	playerInfo.Sex = e.Sex
	playerInfo.Level = e.Level
	playerInfo.ZhuanShu = e.ZhuanShu
	return playerInfo
}

func NewPlayerEntity() *PlayerEntity {
	playerEntity := &PlayerEntity{}
	return playerEntity
}

type PlayerStore interface {
	//获取用户玩家角色服务列表
	GetPlayerServerList(userId int64) ([]*PlayerEntity, error)
	//获取玩家服务角色
	GetPlayerServerEntity(userId int64, serverId int32) (*PlayerEntity, error)
	Save(e *PlayerEntity) (err error)
}

type playerStore struct {
	db coredb.DBService
}

func (s *playerStore) GetPlayerServerList(userId int64) (eList []*PlayerEntity, err error) {
	err = s.db.DB().Find(&eList, "userId=?", userId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError("player", err)
		}
		return
	}
	return
}

func (s *playerStore) GetPlayerServerEntity(userId int64, serverId int32) (playerEntity *PlayerEntity, err error) {
	playerEntity = &PlayerEntity{}
	err = s.db.DB().First(playerEntity, "userId=? and serverId=?", userId, serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError("player", err)
		}
		return nil, nil
	}
	return
}

func (s *playerStore) Save(e *PlayerEntity) (err error) {
	err = s.db.DB().Save(e).Error
	if err != nil {
		return
	}
	return
}

func NewPlayerStore(db coredb.DBService) PlayerStore {
	s := &playerStore{
		db: db,
	}
	return s
}
