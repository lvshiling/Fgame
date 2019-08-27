package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	densewatentity "fgame/fgame/game/densewat/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "dense_wat"
)

type DenseWatDao interface {
	//查询玩家金银密窟
	GetDenseWatEntity(playerId int64) (*densewatentity.PlayerDenseWatEntity, error)
}

type denseWatDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *denseWatDao) GetDenseWatEntity(playerId int64) (denseWatEntity *densewatentity.PlayerDenseWatEntity, err error) {
	denseWatEntity = &densewatentity.PlayerDenseWatEntity{}
	err = dao.ds.DB().First(denseWatEntity, "playerId=?", playerId).Error
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
	dao  *denseWatDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &denseWatDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetDenseWatDao() DenseWatDao {
	return dao
}
