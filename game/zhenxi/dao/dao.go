package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/zhenxi/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "zhenxi"
)

type ZhenXiDao interface {
	//查询玩家阵法
	GetPlayerZhenXiBoss(playerId int64) (*entity.PlayerZhenXiBossEntity, error)
}

type zhenXiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

//查询玩家珍稀boss信息
func (dao *zhenXiDao) GetPlayerZhenXiBoss(playerId int64) (e *entity.PlayerZhenXiBossEntity, err error) {
	e = &entity.PlayerZhenXiBossEntity{}
	err = dao.ds.DB().First(e, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

var (
	once sync.Once
	dao  *zhenXiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &zhenXiDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetZhenXiDao() ZhenXiDao {
	return dao
}
