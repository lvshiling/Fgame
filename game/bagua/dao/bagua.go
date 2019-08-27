package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	baguanentity "fgame/fgame/game/bagua/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "bagua"
)

type BaGuaDao interface {
	//查询玩家八卦秘境
	GetBaGuaEntity(playerId int64) (*baguanentity.PlayerBaGuaEntity, error)
}

type baGuaDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *baGuaDao) GetBaGuaEntity(playerId int64) (baGuaEntity *baguanentity.PlayerBaGuaEntity, err error) {
	baGuaEntity = &baguanentity.PlayerBaGuaEntity{}
	err = dao.ds.DB().First(baGuaEntity, "playerId=?", playerId).Error
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
	dao  *baGuaDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &baGuaDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetBaGuaDao() BaGuaDao {
	return dao
}
