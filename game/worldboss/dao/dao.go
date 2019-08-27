package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	worldbossentity "fgame/fgame/game/worldboss/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "worldboss"
)

type WorldbossDao interface {
	//周卡信息
	GetPlayerBossReliveList(playerId int64) (bossReliveList []*worldbossentity.PlayerBossReliveEntity, err error)
}

type worldbossDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *worldbossDao) GetPlayerBossReliveList(playerId int64) (bossReliveEntityList []*worldbossentity.PlayerBossReliveEntity, err error) {
	bossReliveEntityList = make([]*worldbossentity.PlayerBossReliveEntity, 0, 8)
	err = dao.ds.DB().Find(&bossReliveEntityList, "playerId=? and deleteTime=0 ", playerId).Error
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
	dao  *worldbossDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &worldbossDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetWorldbossDao() WorldbossDao {
	return dao
}
