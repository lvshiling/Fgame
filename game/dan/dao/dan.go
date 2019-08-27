package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	danentity "fgame/fgame/game/dan/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "dan"
)

type DanDao interface {
	//查询玩家食丹
	GetDanEntity(playerId int64) (*danentity.PlayerDanEntity, error)
	//查询玩家炼丹列表
	GetAlchemyList(playerId int64) ([]*danentity.PlayerAlchemyEntity, error)
}

type danDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *danDao) GetDanEntity(playerId int64) (danEntity *danentity.PlayerDanEntity, err error) {
	danEntity = &danentity.PlayerDanEntity{}
	err = dao.ds.DB().First(danEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return

}

func (dao *danDao) GetAlchemyList(playerId int64) (achemyList []*danentity.PlayerAlchemyEntity, err error) {
	err = dao.ds.DB().Find(&achemyList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

var (
	once sync.Once
	dao  *danDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &danDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetDanDao() DanDao {
	return dao
}
