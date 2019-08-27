package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	funcopenentity "fgame/fgame/game/funcopen/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "funcOpen"
)

type FuncOpenDao interface {
	//查询玩家功能开启列表
	GetFuncOpenEntity(playerId int64) (*funcopenentity.PlayerFuncOpenEntity, error)
}

type funcOpenDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *funcOpenDao) GetFuncOpenEntity(playerId int64) (funcOpenEntity *funcopenentity.PlayerFuncOpenEntity, err error) {
	funcOpenEntity = &funcopenentity.PlayerFuncOpenEntity{}
	err = dao.ds.DB().First(funcOpenEntity, "playerId=?", playerId).Error
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
	dao  *funcOpenDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &funcOpenDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetFuncOpenDao() FuncOpenDao {
	return dao
}
