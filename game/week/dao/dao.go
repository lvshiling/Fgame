package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	weekentity "fgame/fgame/game/week/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "week"
)

type WeekDao interface {
	//周卡信息
	GetWeekEntity(playerId int64) (weekEntity *weekentity.PlayerWeekEntity, err error)
}

type weekDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *weekDao) GetWeekEntity(playerId int64) (weekEntity *weekentity.PlayerWeekEntity, err error) {
	weekEntity = &weekentity.PlayerWeekEntity{}
	err = dao.ds.DB().First(weekEntity, "playerId=? and deleteTime=0 ", playerId).Error
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
	dao  *weekDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &weekDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetWeekDao() WeekDao {
	return dao
}
