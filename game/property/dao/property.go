package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	propertyentity "fgame/fgame/game/property/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "Property"
)

type PropertyDao interface {
	//查询玩家属性
	GetPropertyEntity(playerId int64) (*propertyentity.PlayerPropertyEntity, error)
	//查询玩家每日消费记录
	GetPlayerCycleCostRecordEntity(playerId int64) (entity *propertyentity.PlayerCycleCostRecordEntity, err error)
	//查询玩家战力记录
	GetPowerRecordEntity(playerId int64) (*propertyentity.PlayerPowerRecordEntity, error)
	//查询玩家魅力改变日志
	GetCharmAddLogList(playerId int64) (logEntityList []*propertyentity.PlayerCharmAddLogEntity, err error)
}

type propertyDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *propertyDao) GetPropertyEntity(playerId int64) (propertyEntity *propertyentity.PlayerPropertyEntity, err error) {
	propertyEntity = &propertyentity.PlayerPropertyEntity{}
	err = dao.ds.DB().First(propertyEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *propertyDao) GetPlayerCycleCostRecordEntity(playerId int64) (entity *propertyentity.PlayerCycleCostRecordEntity, err error) {
	entity = &propertyentity.PlayerCycleCostRecordEntity{}
	err = dao.ds.DB().First(entity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *propertyDao) GetPowerRecordEntity(playerId int64) (entity *propertyentity.PlayerPowerRecordEntity, err error) {
	entity = &propertyentity.PlayerPowerRecordEntity{}
	err = dao.ds.DB().First(entity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//查询玩家魅力改变日志
func (dao *propertyDao) GetCharmAddLogList(playerId int64) (logEntityList []*propertyentity.PlayerCharmAddLogEntity, err error) {
	err = dao.ds.DB().Find(&logEntityList, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *propertyDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &propertyDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetPropertyDao() PropertyDao {
	return dao
}
