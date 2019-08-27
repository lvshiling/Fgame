package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	unrealbossentity "fgame/fgame/game/unrealboss/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "unrealboss"
)

type UnrealBossDao interface {
	//获取玩家幻境boss
	GetUnrealBossEntity(playerId int64) (unrealbossEntity *unrealbossentity.PlayerUnrealBossEntity, err error)
}

type unrealbossDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *unrealbossDao) GetUnrealBossEntity(playerId int64) (unrealbossEntity *unrealbossentity.PlayerUnrealBossEntity, err error) {
	unrealbossEntity = &unrealbossentity.PlayerUnrealBossEntity{}
	err = dao.ds.DB().First(unrealbossEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *unrealbossDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &unrealbossDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetUnrealBossDao() UnrealBossDao {
	return dao
}
