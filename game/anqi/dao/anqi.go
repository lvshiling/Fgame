package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	anqientity "fgame/fgame/game/anqi/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "anqi"
)

type AnQiDao interface {
	//查询玩家暗器
	GetAnQiEntity(playerId int64) (*anqientity.PlayerAnQiEntity, error)
}

type anQiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *anQiDao) GetAnQiEntity(playerId int64) (anqiEntity *anqientity.PlayerAnQiEntity, err error) {
	anqiEntity = &anqientity.PlayerAnQiEntity{}
	err = dao.ds.DB().First(anqiEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *anQiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &anQiDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetAnQiDao() AnQiDao {
	return dao
}
