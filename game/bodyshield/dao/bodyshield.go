package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	bshieldentity "fgame/fgame/game/bodyshield/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "bodyshield"
)

type BodyShieldDao interface {
	//查询玩家护体盾
	GetBodyShieldEntity(playerId int64) (*bshieldentity.PlayerBodyShieldEntity, error)
}

type bodyShieldDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *bodyShieldDao) GetBodyShieldEntity(playerId int64) (bShieldEntity *bshieldentity.PlayerBodyShieldEntity, err error) {
	bShieldEntity = &bshieldentity.PlayerBodyShieldEntity{}
	err = dao.ds.DB().First(bShieldEntity, "playerId=?", playerId).Error
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
	dao  *bodyShieldDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &bodyShieldDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetBodyShieldDao() BodyShieldDao {
	return dao
}
