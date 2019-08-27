package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	huiyuanentity "fgame/fgame/game/huiyuan/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "huiyuan"
)

type HuiYuanDao interface {
	//会员信息
	GetHuiYuanEntity(playerId int64) (huiyuanEntity *huiyuanentity.PlayerHuiYuanEntity, err error)
}

type huiyuanDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *huiyuanDao) GetHuiYuanEntity(playerId int64) (huiyuanEntity *huiyuanentity.PlayerHuiYuanEntity, err error) {
	huiyuanEntity = &huiyuanentity.PlayerHuiYuanEntity{}
	err = dao.ds.DB().First(huiyuanEntity, "playerId=? and deleteTime=0 ", playerId).Error
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
	dao  *huiyuanDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &huiyuanDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetHuiYuanDao() HuiYuanDao {
	return dao
}
