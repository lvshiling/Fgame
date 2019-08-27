package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	majorentity "fgame/fgame/game/major/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "major_num"
)

type MajorDao interface {
	//查询玩家双休数
	GetMajorNumEntityList(playerId int64) ([]*majorentity.PlayerMajorNumEntity, error)
}

type majorDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *majorDao) GetMajorNumEntityList(playerId int64) (majorNumEntityList []*majorentity.PlayerMajorNumEntity, err error) {
	err = dao.ds.DB().Find(&majorNumEntityList, "playerId=?", playerId).Error
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
	dao  *majorDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &majorDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetMajorDao() MajorDao {
	return dao
}
