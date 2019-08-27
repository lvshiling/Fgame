package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	tulongentity "fgame/fgame/cross/tulong/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "tulong_rank"
)

type TuLongDao interface {
	//查询玩家
	GetTuLongRankList(platform int32, areaId int32) ([]*tulongentity.TuLongRankEntity, error)
}

type tuLongDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *tuLongDao) GetTuLongRankList(platform int32, areaId int32) (tuLongRankList []*tulongentity.TuLongRankEntity, err error) {
	err = dao.ds.DB().Find(&tuLongRankList, "platform=? and areaId=? and deleteTime=0", platform, areaId).Error
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
	dao  *tuLongDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &tuLongDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetTuLongDao() TuLongDao {
	return dao
}
