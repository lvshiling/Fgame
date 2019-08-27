package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	houseentity "fgame/fgame/game/house/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "house"
)

type HouseDao interface {
	//查询玩家房子
	GetHouseEntityList(playerId int64) ([]*houseentity.PlayerHouseEntity, error)
}

type houseDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *houseDao) GetHouseEntityList(playerId int64) (houseEntityList []*houseentity.PlayerHouseEntity, err error) {
	err = dao.ds.DB().Find(&houseEntityList, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *houseDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &houseDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetHouseDao() HouseDao {
	return dao
}
