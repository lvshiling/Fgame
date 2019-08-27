package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	qixueentity "fgame/fgame/game/qixue/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "qixue"
)

type QiXueDao interface {
	//查询玩家泣血枪
	GetQiXueEntity(playerId int64) (*qixueentity.PlayerQiXueEntity, error)
}

type qixueDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *qixueDao) GetQiXueEntity(playerId int64) (qixueEntity *qixueentity.PlayerQiXueEntity, err error) {
	qixueEntity = &qixueentity.PlayerQiXueEntity{}
	err = dao.ds.DB().First(qixueEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *qixueDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &qixueDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetQiXueDao() QiXueDao {
	return dao
}
