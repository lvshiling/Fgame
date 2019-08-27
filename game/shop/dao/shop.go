package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	shopentity "fgame/fgame/game/shop/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "shop"
)

type ShopDao interface {
	//查询玩家商铺购买信息
	GetShopList(playerId int64) ([]*shopentity.PlayerShopEntity, error)
}

type shopDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *shopDao) GetShopList(playerId int64) (shopList []*shopentity.PlayerShopEntity, err error) {
	err = dao.ds.DB().Find(&shopList, "playerId=? and deleteTime=0 ", playerId).Error
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
	dao  *shopDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &shopDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetShopDao() ShopDao {
	return dao
}
