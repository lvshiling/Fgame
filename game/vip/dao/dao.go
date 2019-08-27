package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	vipentity "fgame/fgame/game/vip/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "vip"
)

type VipDao interface {
	//获取玩家VIP数据
	GetVipEntity(playerId int64) (vipEntity *vipentity.PlayerVipEntity, err error)
}

type vipDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *vipDao) GetVipEntity(playerId int64) (vipEntity *vipentity.PlayerVipEntity, err error) {
	vipEntity = &vipentity.PlayerVipEntity{}
	err = dao.ds.DB().First(vipEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *vipDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &vipDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetVipDao() VipDao {
	return dao
}
