package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	shopdiscountentity "fgame/fgame/game/shopdiscount/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "shopdiscount"
)

type ShopDiscountDao interface {
	//查询玩家商城促销
	GetShopDiscountEntity(playerId int64) (*shopdiscountentity.PlayerShopDiscountEntity, error)
}

type shopDiscountDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *shopDiscountDao) GetShopDiscountEntity(playerId int64) (entity *shopdiscountentity.PlayerShopDiscountEntity, err error) {
	entity = &shopdiscountentity.PlayerShopDiscountEntity{}
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
	dao  *shopDiscountDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &shopDiscountDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetShopDiscountDao() ShopDiscountDao {
	return dao
}
