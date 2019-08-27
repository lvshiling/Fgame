package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"

	arenabossentity "fgame/fgame/cross/arenaboss/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "arena_boss"
)

type ArenaBossDao interface {
	//查询玩家点星
	GetArenaBossEntityList(platform int32, serverId int32) (arenaBossList []*arenabossentity.ArenaBossEntity, err error)
}

type arenaBossDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *arenaBossDao) GetArenaBossEntityList(platform int32, serverId int32) (arenaBossList []*arenabossentity.ArenaBossEntity, err error) {
	arenaBossList = make([]*arenabossentity.ArenaBossEntity, 0, 8)
	err = dao.ds.DB().Find(&arenaBossList, "platform=? AND serverId=? AND deleteTime=0", platform, serverId).Error
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
	dao  *arenaBossDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &arenaBossDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetArenaBossDao() ArenaBossDao {
	return dao
}
