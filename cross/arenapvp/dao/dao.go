package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	arenapvpentity "fgame/fgame/cross/arenapvp/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "arenapvp_rank"
)

type ArenapvpDao interface {
	//查询历届霸主
	GetArenapvpBaZhuList(platform int32, serverId int32) ([]*arenapvpentity.ArenapvpBaZhuEntity, error)
}

type arenapvpDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *arenapvpDao) GetArenapvpBaZhuList(platform int32, serverId int32) (arenapvpRankList []*arenapvpentity.ArenapvpBaZhuEntity, err error) {
	err = dao.ds.DB().Order("raceNumber ASC").Find(&arenapvpRankList, "deleteTime=0 and platform=? and serverId=?", platform, serverId).Error
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
	dao  *arenapvpDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &arenapvpDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetArenapvpDao() ArenapvpDao {
	return dao
}
