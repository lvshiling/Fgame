package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	xuechientity "fgame/fgame/game/xuechi/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "xuechi"
)

type XueChiDao interface {
	//查询玩家血池
	GetXueChiEntity(playerId int64) (*xuechientity.PlayerXueChiEntity, error)
}

type xueChiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *xueChiDao) GetXueChiEntity(playerId int64) (xueChiEntity *xuechientity.PlayerXueChiEntity, err error) {
	xueChiEntity = &xuechientity.PlayerXueChiEntity{}
	err = dao.ds.DB().First(xueChiEntity, "playerId=?", playerId).Error
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
	dao  *xueChiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &xueChiDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetXueChiDao() XueChiDao {
	return dao
}
