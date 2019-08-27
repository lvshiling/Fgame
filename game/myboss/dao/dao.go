package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	mybossentity "fgame/fgame/game/myboss/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "myboss"
)

type MyBossDao interface {
	//获取玩家个人boss
	GetMyBossEntity(playerId int64) (mybossEntity *mybossentity.PlayerMyBossEntity, err error)
}

type mybossDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *mybossDao) GetMyBossEntity(playerId int64) (mybossEntity *mybossentity.PlayerMyBossEntity, err error) {
	mybossEntity = &mybossentity.PlayerMyBossEntity{}
	err = dao.ds.DB().First(mybossEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *mybossDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &mybossDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetMyBossDao() MyBossDao {
	return dao
}
