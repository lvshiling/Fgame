package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	cacheentity "fgame/fgame/game/cache/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "cache"
)

type CacheDao interface {
	//查询玩家根据id
	GetPlayerCacheByPlayerId(playerId int64) (*cacheentity.PlayerCacheEntity, error)
	//查询玩家根据名字
	GetPlayerCacheByName(name string, serverId int32) (*cacheentity.PlayerCacheEntity, error)
}

type cacheDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *cacheDao) GetPlayerCacheByPlayerId(playerId int64) (cacheEntity *cacheentity.PlayerCacheEntity, err error) {
	cacheEntity = &cacheentity.PlayerCacheEntity{}
	err = dao.ds.DB().First(cacheEntity, "playerId=? and deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *cacheDao) GetPlayerCacheByName(name string, serverId int32) (cacheEntity *cacheentity.PlayerCacheEntity, err error) {
	cacheEntity = &cacheentity.PlayerCacheEntity{}
	err = dao.ds.DB().First(cacheEntity, "name=? and serverId=? and deleteTime=0", name, serverId).Error
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
	dao  *cacheDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &cacheDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetCacheDao() CacheDao {
	return dao
}
