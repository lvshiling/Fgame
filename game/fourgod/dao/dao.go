package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	fourgodentity "fgame/fgame/game/fourgod/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "four_god"
)

type FourGodDao interface {
	//查询玩家四神遗迹数据
	GetFourGodEntity(playerId int64) (*fourgodentity.PlayerFourGodEntity, error)
}

type fourGodDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *fourGodDao) GetFourGodEntity(playerId int64) (fourGodEntity *fourgodentity.PlayerFourGodEntity, err error) {
	fourGodEntity = &fourgodentity.PlayerFourGodEntity{}
	err = dao.ds.DB().First(fourGodEntity, "playerId=?", playerId).Error
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
	dao  *fourGodDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &fourGodDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetFourGodDao() FourGodDao {
	return dao
}
