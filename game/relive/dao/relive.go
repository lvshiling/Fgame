package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	reliveentity "fgame/fgame/game/relive/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "relive"
)

type ReliveDao interface {
	//查询玩家坐骑信息
	GetPlayerReliveEntity(playerId int64) (*reliveentity.PlayerReliveEntity, error)
}

type reliveDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *reliveDao) GetPlayerReliveEntity(playerId int64) (reliveEntity *reliveentity.PlayerReliveEntity, err error) {
	reliveEntity = &reliveentity.PlayerReliveEntity{}
	err = dao.ds.DB().First(reliveEntity, "playerId=?", playerId).Error
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
	dao  *reliveDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &reliveDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetReliveDao() ReliveDao {
	return dao
}
