package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	feishengentity "fgame/fgame/game/feisheng/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "feisheng"
)

type FeiShengDao interface {
	//获取玩家飞升数据
	GetFeiShengEntity(playerId int64) (feishengEntity *feishengentity.PlayerFeiShengEntity, err error)
	GetFeiShengReceiveEntity(playerId int64) (feishengReceiveEntity *feishengentity.PlayerFeiShengReceiveEntity, err error)
}

type feishengDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *feishengDao) GetFeiShengEntity(playerId int64) (feishengEntity *feishengentity.PlayerFeiShengEntity, err error) {
	feishengEntity = &feishengentity.PlayerFeiShengEntity{}
	err = dao.ds.DB().First(feishengEntity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *feishengDao) GetFeiShengReceiveEntity(playerId int64) (feishengReceiveEntity *feishengentity.PlayerFeiShengReceiveEntity, err error) {
	feishengReceiveEntity = &feishengentity.PlayerFeiShengReceiveEntity{}
	err = dao.ds.DB().First(feishengReceiveEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *feishengDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &feishengDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetFeiShengDao() FeiShengDao {
	return dao
}
