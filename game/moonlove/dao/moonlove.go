package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	moonloveentity "fgame/fgame/game/moonlove/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "moonlove"
)

type MoonloveDao interface {
	//获取玩家月下情缘信息
	GetMoonloveInfo(playerId int64) (*moonloveentity.PlayerMoonloveEntity, error)
}

type moonloveDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *moonloveDao) GetMoonloveInfo(playerId int64) (moonlove *moonloveentity.PlayerMoonloveEntity, err error) {
	moonlove = &moonloveentity.PlayerMoonloveEntity{}
	err = dao.ds.DB().First(moonlove, "playerId = ? AND deleteTime = 0 ", playerId).Error
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
	dao  *moonloveDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) error {
	once.Do(func() {
		dao = &moonloveDao{
			ds: ds,
			rs: rs,
		}
	})

	return nil
}

func GetMoonloveDao() MoonloveDao {
	return dao
}
