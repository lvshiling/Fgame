package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	shihunfanentity "fgame/fgame/game/shihunfan/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "shihunfan"
)

type ShiHunFanDao interface {
	//查询玩家噬魂幡
	GetShiHunFanEntity(playerId int64) (*shihunfanentity.PlayerShiHunFanEntity, error)
}

type shiHunFanDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *shiHunFanDao) GetShiHunFanEntity(playerId int64) (shihunfanEntity *shihunfanentity.PlayerShiHunFanEntity, err error) {
	shihunfanEntity = &shihunfanentity.PlayerShiHunFanEntity{}
	err = dao.ds.DB().First(shihunfanEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *shiHunFanDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &shiHunFanDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetShiHunFanDao() ShiHunFanDao {
	return dao
}
