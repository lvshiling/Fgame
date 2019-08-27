package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	babyentity "fgame/fgame/game/baby/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "baby"
)

type BabyDao interface {
	//获取玩家怀孕数据
	GetPregnantEntity(playerId int64) (entity *babyentity.PlayerPregnantEntity, err error)
	//获取玩家宝宝数据
	GetBabyEntityList(playerId int64) (entityList []*babyentity.PlayerBabyEntity, err error)
	//获取玩家宝宝玩具数据
	GetBabyToySlotEntityList(playerId int64) (entityList []*babyentity.PlayerBabyToySlotEntity, err error)
	//获取配偶宝宝数据
	GetCoupleBabyEntityList(serverId int32) (entityList []*babyentity.CoupleBabyEntity, err error)
	//获取玩家宝宝战力数据
	GetBabyPowerEntity(playerId int64) (entity *babyentity.PlayerBabyPowerEntity, err error)
}

type babyDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *babyDao) GetPregnantEntity(playerId int64) (entity *babyentity.PlayerPregnantEntity, err error) {
	entity = &babyentity.PlayerPregnantEntity{}
	err = dao.ds.DB().First(entity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *babyDao) GetBabyEntityList(playerId int64) (entityList []*babyentity.PlayerBabyEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *babyDao) GetBabyToySlotEntityList(playerId int64) (entityList []*babyentity.PlayerBabyToySlotEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *babyDao) GetCoupleBabyEntityList(serverId int32) (entityList []*babyentity.CoupleBabyEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "serverId=? AND deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *babyDao) GetBabyPowerEntity(playerId int64) (entity *babyentity.PlayerBabyPowerEntity, err error) {
	entity = &babyentity.PlayerBabyPowerEntity{}
	err = dao.ds.DB().First(entity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *babyDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &babyDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetBabyDao() BabyDao {
	return dao
}
