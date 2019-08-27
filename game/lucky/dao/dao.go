package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	luckyentity "fgame/fgame/game/lucky/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "lucky"
)

type LuckyDao interface {
	//获取玩家幸运符数据
	GetLuckyEntityList(playerId int64) (entityList []*luckyentity.PlayerLuckyEntity, err error)
}

type luckyDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *luckyDao) GetLuckyEntityList(playerId int64) (entityList []*luckyentity.PlayerLuckyEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=?", playerId).Error
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
	dao  *luckyDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &luckyDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetLuckyDao() LuckyDao {
	return dao
}
