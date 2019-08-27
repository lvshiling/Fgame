package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	fushientity "fgame/fgame/game/fushi/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "fushi"
)

type FushiDao interface {
	// 查询玩家符石数据
	GetFuShiEntity(playerId int64) ([]*fushientity.PlayerFuShiEntity, error)
}

type fushiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *fushiDao) GetFuShiEntity(playerId int64) (fushiEntity []*fushientity.PlayerFuShiEntity, err error) {
	err = dao.ds.DB().Find(&fushiEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *fushiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &fushiDao{
			ds: ds,
			rs: rs,
		}
	})

	return err
}

func GetFushiDao() FushiDao {
	return dao
}
