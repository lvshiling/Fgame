package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	tianmoentity "fgame/fgame/game/tianmo/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "tianmo"
)

type TianMoDao interface {
	//查询玩家天魔体
	GetTianMoEntity(playerId int64) (*tianmoentity.PlayerTianMoEntity, error)
}

type anQiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *anQiDao) GetTianMoEntity(playerId int64) (tianmoEntity *tianmoentity.PlayerTianMoEntity, err error) {
	tianmoEntity = &tianmoentity.PlayerTianMoEntity{}
	err = dao.ds.DB().First(tianmoEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *anQiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &anQiDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetTianMoDao() TianMoDao {
	return dao
}
