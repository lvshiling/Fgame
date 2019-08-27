package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	dianxingentity "fgame/fgame/game/dianxing/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "dianxing"
)

type DianXingDao interface {
	//查询玩家点星
	GetDianXingEntity(playerId int64) (*dianxingentity.PlayerDianXingEntity, error)
}

type dianXingDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *dianXingDao) GetDianXingEntity(playerId int64) (dianXingEntity *dianxingentity.PlayerDianXingEntity, err error) {
	dianXingEntity = &dianxingentity.PlayerDianXingEntity{}
	err = dao.ds.DB().First(dianXingEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *dianXingDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &dianXingDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetDianXingDao() DianXingDao {
	return dao
}
