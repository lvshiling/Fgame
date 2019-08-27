package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	dragonentity "fgame/fgame/game/dragon/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "dragon"
)

type DragonDao interface {
	//查询玩家神龙
	GetDragonEntity(playerId int64) (*dragonentity.PlayerDragonEntity, error)
}

type dragonDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *dragonDao) GetDragonEntity(playerId int64) (dragonEntity *dragonentity.PlayerDragonEntity, err error) {
	dragonEntity = &dragonentity.PlayerDragonEntity{}
	err = dao.ds.DB().First(dragonEntity, "playerId=?", playerId).Error
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
	dao  *dragonDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &dragonDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetDragonDao() DragonDao {
	return dao
}
