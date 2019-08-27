package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	shenyuentity "fgame/fgame/game/shenyu/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "shenyu"
)

type ShenYuDao interface {
	//获取玩家神域数据
	GetShenYuEntity(playerId int64) (shenyuEntity *shenyuentity.PlayerShenYuEntity, err error)
}

type shenyuDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *shenyuDao) GetShenYuEntity(playerId int64) (shenyuEntity *shenyuentity.PlayerShenYuEntity, err error) {
	shenyuEntity = &shenyuentity.PlayerShenYuEntity{}
	err = dao.ds.DB().First(shenyuEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *shenyuDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &shenyuDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetShenYuDao() ShenYuDao {
	return dao
}
