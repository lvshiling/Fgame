package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	tianshuentity "fgame/fgame/game/tianshu/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "tianshu"
)

type TianShuDao interface {
	//获取玩家天书数据
	GetTianShuEntityList(playerId int64) (tianshuEntity []*tianshuentity.PlayerTianShuEntity, err error)
}

type tianshuDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *tianshuDao) GetTianShuEntityList(playerId int64) (entityList []*tianshuentity.PlayerTianShuEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *tianshuDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &tianshuDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetTianShuDao() TianShuDao {
	return dao
}
