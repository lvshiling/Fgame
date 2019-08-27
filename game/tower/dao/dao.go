package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	towerentity "fgame/fgame/game/tower/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "tower"
)

type TowerDao interface {
	//获取玩家打宝塔数据
	GetTowerEntity(playerId int64) (towerEntity *towerentity.PlayerTowerEntity, err error)
}

type towerDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *towerDao) GetTowerEntity(playerId int64) (towerEntity *towerentity.PlayerTowerEntity, err error) {
	towerEntity = &towerentity.PlayerTowerEntity{}
	err = dao.ds.DB().First(towerEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *towerDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &towerDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetTowerDao() TowerDao {
	return dao
}
