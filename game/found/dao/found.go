package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	foundentity "fgame/fgame/game/found/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "found"
)

type FoundResDao interface {
	//查询可以找回资源的列表
	GetFoundResList(playerId int64) ([]*foundentity.PlayerFoundEntity, error)
	//查询上一天资源找回列表
	GetFoundBackList(playerId int64) ([]*foundentity.PlayerFoundBackEntity, error)
}

type foundResDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *foundResDao) GetFoundResList(playerId int64) (foundEntityList []*foundentity.PlayerFoundEntity, err error) {
	err = dao.ds.DB().Find(&foundEntityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *foundResDao) GetFoundBackList(playerId int64) (foundBackEntityList []*foundentity.PlayerFoundBackEntity, err error) {
	err = dao.ds.DB().Find(&foundBackEntityList, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *foundResDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &foundResDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetFoundResDao() FoundResDao {
	return dao
}
