package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	xuedunentity "fgame/fgame/game/xuedun/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "xuedun"
)

type XueDunDao interface {
	//查询玩家血盾信息
	GetXueDunEntity(playerId int64) (*xuedunentity.PlayerXueDunEntity, error)
}

type xueDunDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *xueDunDao) GetXueDunEntity(playerId int64) (xueDunEntity *xuedunentity.PlayerXueDunEntity, err error) {
	xueDunEntity = &xuedunentity.PlayerXueDunEntity{}
	err = dao.ds.DB().First(xueDunEntity, "playerId=?", playerId).Error
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
	dao  *xueDunDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &xueDunDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetXueDunDao() XueDunDao {
	return dao
}
