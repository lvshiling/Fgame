package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	chuangshientity "fgame/fgame/game/chuangshi/entity"
	"fgame/fgame/game/global"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "chuangshi"
)

type ChuangShiDao interface {
	// 查询玩家创世之战报名信息
	GetPlayerChuangShiYuGaoEntity(playerId int64) (*chuangshientity.PlayerChuangShiYuGaoEntity, error)
	// 查询创世之战报名人数
	GetChuangShiYuGaoEntity() (*chuangshientity.ChuangShiYuGaoEntity, error)
	// 查询玩家创世之战信息
	GetPlayerChuangShiEntity(playerId int64) (*chuangshientity.PlayerChuangShiEntity, error)
	// 查询玩家神王报名记录
	GetPlayerChuangShiSignEntity(playerId int64) (*chuangshientity.PlayerChuangShiSignEntity, error)
	// 查询玩家神王投票记录
	GetPlayerChuangShiVoetEntity(playerId int64) (*chuangshientity.PlayerChuangShiVoteEntity, error)

	// 查询创世神王报名记录
	GetChuangShiShenWangSignUpEntityList(serverId int32) ([]*chuangshientity.ChuangShiShenWangSignUpEntity, error)
	// 查询创世神王投票记录
	GetChuangShiShenWangVoteEntityList(serverId int32) ([]*chuangshientity.ChuangShiShenWangVoteEntity, error)
	// 查询创世城池建设记录
	GetChuangShiChengFangJianSheEntityList(serverId int32) ([]*chuangshientity.ChuangShiChengFangJianSheEntity, error)
	// 查询创世之战玩家官职信息
	GetPlayerChuangShiGuanZhiEntity(playerId int64) (*chuangshientity.PlayerChuangShiGuanZhiEntity, error)
}

type chuangShiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *chuangShiDao) GetPlayerChuangShiYuGaoEntity(playerId int64) (e *chuangshientity.PlayerChuangShiYuGaoEntity, err error) {
	e = &chuangshientity.PlayerChuangShiYuGaoEntity{}
	err = dao.ds.DB().First(e, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chuangShiDao) GetPlayerChuangShiSignEntity(playerId int64) (e *chuangshientity.PlayerChuangShiSignEntity, err error) {
	e = &chuangshientity.PlayerChuangShiSignEntity{}
	err = dao.ds.DB().First(e, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chuangShiDao) GetPlayerChuangShiVoetEntity(playerId int64) (e *chuangshientity.PlayerChuangShiVoteEntity, err error) {
	e = &chuangshientity.PlayerChuangShiVoteEntity{}
	err = dao.ds.DB().First(e, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chuangShiDao) GetPlayerChuangShiEntity(playerId int64) (e *chuangshientity.PlayerChuangShiEntity, err error) {
	e = &chuangshientity.PlayerChuangShiEntity{}
	err = dao.ds.DB().First(e, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chuangShiDao) GetPlayerChuangShiGuanZhiEntity(playerId int64) (e *chuangshientity.PlayerChuangShiGuanZhiEntity, err error) {
	e = &chuangshientity.PlayerChuangShiGuanZhiEntity{}
	err = dao.ds.DB().First(e, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chuangShiDao) GetChuangShiYuGaoEntity() (e *chuangshientity.ChuangShiYuGaoEntity, err error) {
	e = &chuangshientity.ChuangShiYuGaoEntity{}
	err = dao.ds.DB().First(e, "serverId=? AND deleteTime=0", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chuangShiDao) GetChuangShiShenWangSignUpEntityList(serverId int32) (eList []*chuangshientity.ChuangShiShenWangSignUpEntity, err error) {
	err = dao.ds.DB().Find(&eList, "serverId=? AND deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chuangShiDao) GetChuangShiShenWangVoteEntityList(serverId int32) (eList []*chuangshientity.ChuangShiShenWangVoteEntity, err error) {
	err = dao.ds.DB().Find(&eList, "serverId=? AND deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chuangShiDao) GetChuangShiChengFangJianSheEntityList(serverId int32) (eList []*chuangshientity.ChuangShiChengFangJianSheEntity, err error) {
	err = dao.ds.DB().Find(&eList, "serverId=? AND deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

var (
	once sync.Once
	dao  *chuangShiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) error {
	once.Do(func() {
		dao = &chuangShiDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetChuangShiDao() ChuangShiDao {
	return dao
}
