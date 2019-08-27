package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	pkentity "fgame/fgame/game/pk/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "pk"
)

type PKDao interface {
	//查询玩家pk
	GetPKEntity(playerId int64) (*pkentity.PlayerPkEntity, error)
}

type pKDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *pKDao) GetPKEntity(playerId int64) (pkEntity *pkentity.PlayerPkEntity, err error) {
	pkEntity = &pkentity.PlayerPkEntity{}
	err = dao.ds.DB().First(pkEntity, "playerId=?", playerId).Error
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
	dao  *pKDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &pKDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetPKDao() PKDao {
	return dao
}
