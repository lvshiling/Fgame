package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	huntentity "fgame/fgame/game/hunt/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "hunt"
)

type HuntDao interface {
	//获取寻宝列表
	GetHuntEntityList(playerId int64) ([]*huntentity.PlayerHuntEntity, error)
}

type tuLongEquipDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *tuLongEquipDao) GetHuntEntityList(playerId int64) (entityList []*huntentity.PlayerHuntEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=?", playerId).Error
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
	dao  *tuLongEquipDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &tuLongEquipDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetHuntDao() HuntDao {
	return dao
}
