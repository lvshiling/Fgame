package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	shangguzhilingentity "fgame/fgame/game/shangguzhiling/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "shangguzhiling"
)

type ShangguzhilingDao interface {
	//上古灵兽信息
	GetShangguzhilingEntity(playerId int64) ([]*shangguzhilingentity.PlayerShangguzhilingEntity, error)
}

type shangguzhilingDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

//上古灵兽信息
func (scd *shangguzhilingDao) GetShangguzhilingEntity(playerId int64) (lingEntityList []*shangguzhilingentity.PlayerShangguzhilingEntity, err error) {
	err = dao.ds.DB().Order("`lingShouType` ASC").Find(&lingEntityList, "playerId=?", playerId).Error
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
	dao  *shangguzhilingDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &shangguzhilingDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetShangguzhilingDao() ShangguzhilingDao {
	return dao
}
