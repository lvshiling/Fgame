package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	mergeentity "fgame/fgame/game/merge/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "merge"
)

type MergeDao interface {
	//获取合服数据
	GetMergeEntity(serverId int32) (*mergeentity.MergeEntity, error)
}

type mergeDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *mergeDao) GetMergeEntity(serverId int32) (mergeEntity *mergeentity.MergeEntity, err error) {
	mergeEntity = &mergeentity.MergeEntity{}
	err = dao.ds.DB().First(mergeEntity, "serverId=?", serverId).Error
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
	dao  *mergeDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &mergeDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetMergeDao() MergeDao {
	return dao
}
