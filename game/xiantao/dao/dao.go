package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	xiantaoentity "fgame/fgame/game/xiantao/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "xian_tao"
)

type XianTaoDao interface {
	//查询玩家仙桃大会数据
	GetXianTaoEntity(playerId int64) (*xiantaoentity.PlayerXianTaoEntity, error)
}

type xianTaoDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *xianTaoDao) GetXianTaoEntity(playerId int64) (xianTaoEntity *xiantaoentity.PlayerXianTaoEntity, err error) {
	xianTaoEntity = &xiantaoentity.PlayerXianTaoEntity{}
	err = dao.ds.DB().First(xianTaoEntity, "playerId=?", playerId).Error
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
	dao  *xianTaoDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &xianTaoDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetXianTaoDao() XianTaoDao {
	return dao
}
