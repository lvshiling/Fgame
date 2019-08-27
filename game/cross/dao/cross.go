package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	crossentity "fgame/fgame/game/cross/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "cross"
)

type CrossDao interface {
	//查询玩家跨服数据
	GetCrossEntity(playerId int64) (*crossentity.PlayerCrossEntity, error)
}

type crossDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *crossDao) GetCrossEntity(playerId int64) (e *crossentity.PlayerCrossEntity, err error) {
	e = &crossentity.PlayerCrossEntity{}
	err = dao.ds.DB().First(e, "playerId=?", playerId).Error
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
	dao  *crossDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &crossDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetCrossDao() CrossDao {
	return dao
}
