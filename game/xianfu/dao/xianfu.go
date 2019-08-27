package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/xianfu/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "xianfu"
)

type XianfuDao interface {
	//查找玩家仙府信息
	GetXianfuInfo(playerId int64) ([]*entity.PlayerXianFuEntity, error)
}

type xianfuDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *xianfuDao) GetXianfuInfo(playerId int64) (xfEntityArr []*entity.PlayerXianFuEntity, err error) {
	err = dao.ds.DB().Find(&xfEntityArr, "playerId = ? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

var (
	once  sync.Once
	dao *xianfuDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) error {
	once.Do(func() {
		dao = &xianfuDao{
			ds: ds,
			rs: rs,
		}
	})

	return nil
}

func GetXianfuDao() XianfuDao {
	return dao
}
