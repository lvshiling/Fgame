package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"

	dingshientity "fgame/fgame/game/dingshi/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "dingshi"
)

type DingShiDao interface {
	//查询玩家点星
	GetDingShiBossEntityList(serverId int32) (dingShiBossList []*dingshientity.DingShiBossEntity, err error)
}

type dingShiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *dingShiDao) GetDingShiBossEntityList(serverId int32) (dingShiBossList []*dingshientity.DingShiBossEntity, err error) {
	dingShiBossList = make([]*dingshientity.DingShiBossEntity, 0, 8)
	err = dao.ds.DB().Find(&dingShiBossList, "serverId=? AND deleteTime=0", serverId).Error
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
	dao  *dingShiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &dingShiDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetDingShiDao() DingShiDao {
	return dao
}
