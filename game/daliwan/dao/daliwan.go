package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	daliwanentity "fgame/fgame/game/daliwan/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "daliwan"
)

type DaLiWanDao interface {
	GetDailiWanList(playerId int64) ([]*daliwanentity.DaLiWanEntity, error)
}

type daLiWanDao struct {
	ds coredb.DBService
}

func (t *daLiWanDao) GetDailiWanList(playerId int64) ([]*daliwanentity.DaLiWanEntity, error) {
	rst := make([]*daliwanentity.DaLiWanEntity, 0, 8)
	exdb := t.ds.DB().Where("playerid = ? and deleteTime =0 ", playerId).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, coredb.NewDBError(dbName, exdb.Error)
	}
	return rst, nil
}

var (
	once sync.Once
	dao  *daLiWanDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &daLiWanDao{
			ds: ds,
		}
	})
	return nil
}

func GetDaLiWanDao() DaLiWanDao {
	return dao
}
