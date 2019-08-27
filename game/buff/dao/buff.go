package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	buffentity "fgame/fgame/game/buff/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "buff"
)

type BuffDao interface {
	GetBuff(playerId int64) (e *buffentity.PlayerBuffEntity, err error)
}

type buffDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (d *buffDao) GetBuff(playerId int64) (e *buffentity.PlayerBuffEntity, err error) {
	e = &buffentity.PlayerBuffEntity{}
	err = dao.ds.DB().First(e, "playerId=?", playerId).Error
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
	dao  *buffDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &buffDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetBuffDao() BuffDao {
	return dao
}
