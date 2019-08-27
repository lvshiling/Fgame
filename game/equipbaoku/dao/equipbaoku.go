package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	equipbaokuentity "fgame/fgame/game/equipbaoku/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "equipbaoku"
)

type EquipBaoKuDao interface {
	//获取玩家装备宝库信息
	GetPlayerEquipBaoKuEntity(playerId int64) (baokuEntity []*equipbaokuentity.PlayerEquipBaoKuEntity, err error)
	//查询玩家商铺购买信息
	GetEquipBaoKuShopList(playerId int64) ([]*equipbaokuentity.PlayerEquipBaoKuShopEntity, error)
}

type equipBaoKuDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *equipBaoKuDao) GetPlayerEquipBaoKuEntity(playerId int64) (baokuEntity []*equipbaokuentity.PlayerEquipBaoKuEntity, err error) {
	err = dao.ds.DB().Find(&baokuEntity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *equipBaoKuDao) GetEquipBaoKuShopList(playerId int64) (shopList []*equipbaokuentity.PlayerEquipBaoKuShopEntity, err error) {
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
	dao  *equipBaoKuDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &equipBaoKuDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetEquipBaoKuDao() EquipBaoKuDao {
	return dao
}
