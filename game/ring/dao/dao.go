package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/ring/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "ring"
)

type RingDao interface {
	// 查询特戒信息
	GetPlayerRingEntity(playerId int64) ([]*entity.PlayerRingEntity, error)
	// 查询特戒宝库信息
	GetPlayerRingBaoKuEntity(playerId int64) ([]*entity.PlayerRingBaoKuEntity, error)
}

type ringDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *ringDao) GetPlayerRingEntity(playerId int64) (entityList []*entity.PlayerRingEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *ringDao) GetPlayerRingBaoKuEntity(playerId int64) (ringEntity []*entity.PlayerRingBaoKuEntity, err error) {
	err = dao.ds.DB().Find(&ringEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *ringDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &ringDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetRingDao() RingDao {
	return dao
}
