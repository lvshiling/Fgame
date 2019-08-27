package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/mingge/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "mingge"
)

type MingGeDao interface {
	//查询玩家命盘
	GetMingPanList(playerId int64) ([]*entity.PlayerMingGePanEntity, error)
	//查询玩家命理
	GetMingLiList(playerId int64) ([]*entity.PlayerMingGeMingLiEntity, error)
	//查询玩家命盘祭炼
	GetMingGeRefinedList(playerId int64) ([]*entity.PlayerMingGeRefinedEntity, error)
	//获取补偿
	GetMingGeBuchang(playerId int64) (*entity.PlayerMingGeBuchangEntity, error)
	// 查询玩家命格数据
	GetMingGeEntity(playerId int64) (*entity.PlayerMingGeEntity, error)
}

type mingGeDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

//查询玩家命盘
func (dao *mingGeDao) GetMingPanList(playerId int64) (mingGeList []*entity.PlayerMingGePanEntity, err error) {
	err = dao.ds.DB().Find(&mingGeList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//查询玩家命理
func (dao *mingGeDao) GetMingLiList(playerId int64) (mingGeList []*entity.PlayerMingGeMingLiEntity, err error) {
	err = dao.ds.DB().Find(&mingGeList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//查询玩家命盘祭炼
func (dao *mingGeDao) GetMingGeRefinedList(playerId int64) (mingGeList []*entity.PlayerMingGeRefinedEntity, err error) {
	err = dao.ds.DB().Find(&mingGeList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//查询玩家命盘祭炼
func (dao *mingGeDao) GetMingGeBuchang(playerId int64) (e *entity.PlayerMingGeBuchangEntity, err error) {
	e = &entity.PlayerMingGeBuchangEntity{}
	err = dao.ds.DB().First(e, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *mingGeDao) GetMingGeEntity(playerId int64) (mingGeEntity *entity.PlayerMingGeEntity, err error) {
	mingGeEntity = &entity.PlayerMingGeEntity{}
	err = dao.ds.DB().First(mingGeEntity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

var (
	once sync.Once
	dao  *mingGeDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &mingGeDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetMingGeDao() MingGeDao {
	return dao
}
